package api

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"architectingsoftware.com/reading-list-api/schema"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/nitishm/go-rejson/v4"

	"github.com/go-resty/resty/v2"
)

type cache struct {
	client  *redis.Client
	helper  *rejson.Handler
	context context.Context
}

type ReadingListAPI struct {
	cache
	pubAPIURL string
	apiClient *resty.Client
}

func NewReadingListAPI(location string, pubAPIurl string) (*ReadingListAPI, error) {

	apiClient := resty.New()
	//Connect to redis.  Other options can be provided, but the
	//defaults are OK
	client := redis.NewClient(&redis.Options{
		Addr: location,
	})

	//We use this context to coordinate betwen our go code and
	//the redis operaitons
	ctx := context.Background()

	//This is the reccomended way to ensure that our redis connection
	//is working
	err := client.Ping(ctx).Err()
	if err != nil {
		log.Println("Error connecting to redis" + err.Error())
		return nil, err
	}

	//By default, redis manages keys and values, where the values
	//are either strings, sets, maps, etc.  Redis has an extension
	//module called ReJSON that allows us to store JSON objects
	//however, we need a companion library in order to work with it
	//Below we create an instance of the JSON helper and associate
	//it with our redis connnection
	jsonHelper := rejson.NewReJSONHandler()
	jsonHelper.SetGoRedisClientWithContext(ctx, client)

	//Return a pointer to a new ToDo struct
	return &ReadingListAPI{
		cache: cache{
			client:  client,
			helper:  jsonHelper,
			context: ctx,
		},
		pubAPIURL: pubAPIurl,
		apiClient: apiClient,
	}, nil
}

func (r *ReadingListAPI) GetReadingList(c *gin.Context) {

	rlId := c.Param("id")
	if rlId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No publication ID provided"})
		return
	}

	cacheKey := "publist:" + rlId
	rlBytes, err := r.helper.JSONGet(cacheKey, ".")
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Could not find reading list in cache with id=" + cacheKey})
		return
	}

	var rl schema.ReadingList
	err = json.Unmarshal(rlBytes.([]byte), &rl)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "cached data seems to be wrong type"})
		return
	}

	c.JSON(http.StatusOK, rl)
}

func (r *ReadingListAPI) GetPubFromReadingList(c *gin.Context) {
	rlId := c.Param("id")
	if rlId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No publication ID provided"})
		return
	}

	rlIdxKey := c.Param("idx")
	if rlIdxKey == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No publication incdex provided"})
		return
	}

	cacheKey := "publist:" + rlId
	var rl schema.ReadingList
	err := r.getItemFromRedis(cacheKey, &rl)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Could not find reading list in cache with id=" + cacheKey})
		return
	}

	pubItemLocation, ok := rl.Items[rlIdxKey]
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Could not find publication in reading list with id=" + rlIdxKey})
		return
	}

	pubURL := r.pubAPIURL + pubItemLocation
	var pub schema.Publication

	_, err = r.apiClient.R().SetResult(&pub).Get(pubURL)
	if err != nil {
		emsg := "Could not get publication from API: (" + pubURL + ")" + err.Error()
		c.JSON(http.StatusNotFound, gin.H{"error": emsg})
		return
	}

	c.JSON(http.StatusOK, pub)
}

func (r *ReadingListAPI) RedirectWithPublication(c *gin.Context) {
	rlId := c.Param("id")
	if rlId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No publication ID provided"})
		return
	}

	rlIdxKey := c.Param("idx")
	if rlIdxKey == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No publication incdex provided"})
		return
	}

	cacheKey := "publist:" + rlId
	var rl schema.ReadingList
	err := r.getItemFromRedis(cacheKey, &rl)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Could not find reading list in cache with id=" + cacheKey})
		return
	}

	pubItemLocation, ok := rl.Items[rlIdxKey]
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Could not find publication in reading list with id=" + rlIdxKey})
		return
	}

	pubURL := r.pubAPIURL + pubItemLocation
	var pub schema.Publication

	_, err = r.apiClient.R().SetResult(&pub).Get(pubURL)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Could not get publication from API"})
		return
	}

	if pub.Link == "" {
		c.JSON(http.StatusNotFound, gin.H{"error": "Publication does not have a link"})
		return
	}

	c.Redirect(http.StatusMovedPermanently, pub.Link)
}

func (r *ReadingListAPI) GetReadingLists(c *gin.Context) {

	var readList []schema.ReadingList
	var readItem schema.ReadingList

	//Lets query redis for all of the items
	pattern := "publist:*"
	ks, _ := r.client.Keys(r.context, pattern).Result()
	for _, key := range ks {
		err := r.getItemFromRedis(key, &readItem)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not find reading list in cache with id=" + key})
			return
		}
		readList = append(readList, readItem)
	}

	c.JSON(http.StatusOK, readList)
}

// Helper to return a ToDoItem from redis provided a key
func (r *ReadingListAPI) getItemFromRedis(key string, rl *schema.ReadingList) error {

	//Lets query redis for the item, note we can return parts of the
	//json structure, the second parameter "." means return the entire
	//json structure
	itemObject, err := r.helper.JSONGet(key, ".")
	if err != nil {
		return err
	}

	//JSONGet returns an "any" object, or empty interface,
	//we need to convert it to a byte array, which is the
	//underlying type of the object, then we can unmarshal
	//it into our ToDoItem struct
	err = json.Unmarshal(itemObject.([]byte), rl)
	if err != nil {
		return err
	}

	return nil
}
