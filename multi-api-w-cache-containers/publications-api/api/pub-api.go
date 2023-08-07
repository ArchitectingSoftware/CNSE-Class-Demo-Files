package api

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"architectingsoftware.com/pub-api/schema"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/nitishm/go-rejson/v4"
)

type cache struct {
	client  *redis.Client
	helper  *rejson.Handler
	context context.Context
}

type PubAPI struct {
	cache
}

func NewPubAPI(location string) (*PubAPI, error) {

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
	return &PubAPI{
		cache: cache{
			client:  client,
			helper:  jsonHelper,
			context: ctx,
		},
	}, nil
}

func (p *PubAPI) GetPublication(c *gin.Context) {

	pubid := c.Param("id")
	if pubid == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No publication ID provided"})
		return
	}

	cacheKey := "pubs:" + pubid
	pubBytes, err := p.helper.JSONGet(cacheKey, ".")
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Could not find publication in cache with id=" + cacheKey})
		return
	}

	var pub schema.Publication
	err = json.Unmarshal(pubBytes.([]byte), &pub)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "cached data seems to be wrong type"})
		return
	}

	c.JSON(http.StatusOK, pub)
}

func (p *PubAPI) GetPublications(c *gin.Context) {

	var pubList []schema.Publication
	var pubItem schema.Publication

	//Lets query redis for all of the items
	pattern := "pubs:*"
	ks, _ := p.client.Keys(p.context, pattern).Result()
	for _, key := range ks {
		err := p.getItemFromRedis(key, &pubItem)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not find publication in cache with id=" + key})
			return
		}
		pubList = append(pubList, pubItem)
	}

	c.JSON(http.StatusOK, pubList)
}

// Helper to return a ToDoItem from redis provided a key
func (p *PubAPI) getItemFromRedis(key string, pub *schema.Publication) error {

	//Lets query redis for the item, note we can return parts of the
	//json structure, the second parameter "." means return the entire
	//json structure
	itemObject, err := p.helper.JSONGet(key, ".")
	if err != nil {
		return err
	}

	//JSONGet returns an "any" object, or empty interface,
	//we need to convert it to a byte array, which is the
	//underlying type of the object, then we can unmarshal
	//it into our ToDoItem struct
	err = json.Unmarshal(itemObject.([]byte), pub)
	if err != nil {
		return err
	}

	return nil
}
