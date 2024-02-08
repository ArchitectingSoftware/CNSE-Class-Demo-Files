package db

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/redis/go-redis/v9"
)

// ToDoItem is the struct that represents a single ToDo item
type ToDoItem struct {
	Id     int    `json:"id"`
	Title  string `json:"title"`
	IsDone bool   `json:"done"`
}

const (
	RedisNilError        = "redis: nil"
	RedisDefaultLocation = "0.0.0.0:6379"
	RedisKeyPrefix       = "todo:"
)

type cache struct {
	client  *redis.Client
	context context.Context
}

// ToDo is the struct that represents the main object of our
// todo app.  It contains a reference to a cache object
type ToDo struct {
	//more things would be included in a real implementation

	//Redis cache connections
	cache
}

// New is a constructor function that returns a pointer to a new
// ToDo struct.  If this is called it uses the default Redis URL
// with the companion constructor NewWithCacheInstance.
func New() (*ToDo, error) {
	//We will use an override if the REDIS_URL is provided as an environment
	//variable, which is the preferred way to wire up a docker container
	redisUrl := os.Getenv("REDIS_URL")
	//This handles the default condition
	if redisUrl == "" {
		redisUrl = RedisDefaultLocation
	}
	return NewWithCacheInstance(redisUrl)
}

// NewWithCacheInstance is a constructor function that returns a pointer to a new
// ToDo struct.  It accepts a string that represents the location of the redis
// cache.
func NewWithCacheInstance(location string) (*ToDo, error) {

	//Connect to redis.  Other options can be provided, but the
	//defaults are OK
	client := redis.NewClient(&redis.Options{
		Addr: location,
	})

	//We use this context to coordinate betwen our go code and
	//the redis operaitons
	ctx := context.TODO()

	//This is the reccomended way to ensure that our redis connection
	//is working
	err := client.Ping(ctx).Err()
	if err != nil {
		log.Println("Error connecting to redis" + err.Error())
		return nil, err
	}

	//Return a pointer to a new ToDo struct
	return &ToDo{
		cache: cache{
			client:  client,
			context: ctx,
		},
	}, nil
}

//------------------------------------------------------------
// REDIS HELPERS
//------------------------------------------------------------

// We will use this later, you can ignore for now
func isRedisNilError(err error) bool {
	return errors.Is(err, redis.Nil) || err.Error() == RedisNilError
}

// In redis, our keys will be strings, they will look like
// todo:<number>.  This function will take an integer and
// return a string that can be used as a key in redis
func redisKeyFromId(id int) string {
	return fmt.Sprintf("%s%d", RedisKeyPrefix, id)
}

// getAllKeys will return all keys in the database that match the prefix
// used in this application - RedisKeyPrefix.  It will return a string slice
// of all keys.  Used by GetAll and DeleteAll
func (t *ToDo) getAllKeys() ([]string, error) {
	key := fmt.Sprintf("%s*", RedisKeyPrefix)
	return t.client.Keys(t.context, key).Result()
}

func fromJsonString(s string, item *ToDoItem) error {
	err := json.Unmarshal([]byte(s), &item)
	if err != nil {
		return err
	}
	return nil
}

// upsertToDo will be used by insert and update, Redis only supports upserts
// so we will check if an item exists before update, and if it does not exist
// before insert
func (t *ToDo) upsertToDo(item *ToDoItem) error {
	log.Println("Adding new Id:", redisKeyFromId(item.Id))
	return t.client.JSONSet(t.context, redisKeyFromId(item.Id), ".", item).Err()
}

// Helper to return a ToDoItem from redis provided a key
func (t *ToDo) getItemFromRedis(key string, item *ToDoItem) error {

	//Lets query redis for the item, note we can return parts of the
	//json structure, the second parameter "." means return the entire
	//json structure
	itemJson, err := t.client.JSONGet(t.context, key, ".").Result()
	if err != nil {
		return err
	}

	return fromJsonString(itemJson, item)
}

func (t *ToDo) doesKeyExist(id int) bool {
	kc, _ := t.client.Exists(t.context, redisKeyFromId(id)).Result()
	return kc > 0
}

//------------------------------------------------------------
// THESE ARE THE PUBLIC FUNCTIONS THAT SUPPORT OUR TODO APP
//------------------------------------------------------------

// AddItem accepts a ToDoItem and adds it to the DB.
// Preconditions:   (1) The database file must exist and be a valid
//
//					(2) The item must not already exist in the DB
//	    				because we use the item.Id as the key, this
//						function must check if the item already
//	    				exists in the DB, if so, return an error
//
// Postconditions:
//
//	 (1) The item will be added to the DB
//		(2) The DB file will be saved with the item added
//		(3) If there is an error, it will be returned
func (t *ToDo) AddItem(item *ToDoItem) error {

	if t.doesKeyExist(item.Id) {
		return fmt.Errorf("ToDo item with id %d already exists", item.Id)
	}
	return t.upsertToDo(item)
}

// DeleteItem accepts an item id and removes it from the DB.
// Preconditions:   (1) The database file must exist and be a valid
//
//					(2) The item must exist in the DB
//	    				because we use the item.Id as the key, this
//						function must check if the item already
//	    				exists in the DB, if not, return an error
//
// Postconditions:
//
//	 (1) The item will be removed from the DB
//		(2) The DB file will be saved with the item removed
//		(3) If there is an error, it will be returned
func (t *ToDo) DeleteItem(id int) error {
	if !t.doesKeyExist(id) {
		return fmt.Errorf("ToDo item with id %d does not exist", id)
	}
	return t.client.Del(t.context, redisKeyFromId(id)).Err()
}

// DeleteAll removes all items from the DB.
// It will be exposed via a DELETE /todo endpoint
func (t *ToDo) DeleteAll() (int, error) {
	keyList, err := t.getAllKeys()
	if err != nil {
		return 0, err
	}

	//Notice how we can deconstruct the slice into a variadic argument
	//for the Del function by using the ... operator
	numDeleted, err := t.client.Del(t.context, keyList...).Result()
	return int(numDeleted), err
}

// UpdateItem accepts a ToDoItem and updates it in the DB.
// Preconditions:   (1) The database file must exist and be a valid
//
//					(2) The item must exist in the DB
//	    				because we use the item.Id as the key, this
//						function must check if the item already
//	    				exists in the DB, if not, return an error
//
// Postconditions:
//
//	 (1) The item will be updated in the DB
//		(2) The DB file will be saved with the item updated
//		(3) If there is an error, it will be returned
func (t *ToDo) UpdateItem(item *ToDoItem) error {
	if !t.doesKeyExist(item.Id) {
		return fmt.Errorf("ToDo item with id %d does not exist", item.Id)
	}
	return t.upsertToDo(item)
}

// GetItem accepts an item id and returns the item from the DB.
// Preconditions:   (1) The database file must exist and be a valid
//
//					(2) The item must exist in the DB
//	    				because we use the item.Id as the key, this
//						function must check if the item already
//	    				exists in the DB, if not, return an error
//
// Postconditions:
//
//	 (1) The item will be returned, if it exists
//		(2) If there is an error, it will be returned
//			along with an empty ToDoItem
//		(3) The database file will not be modified
func (t *ToDo) GetItem(id int) (*ToDoItem, error) {
	newToDo := &ToDoItem{}
	err := t.getItemFromRedis(redisKeyFromId(id), newToDo)
	if err != nil {
		return nil, err
	}
	return newToDo, nil
}

// ChangeItemDoneStatus accepts an item id and a boolean status.
// It returns an error if the status could not be updated for any
// reason.  For example, the item itself does not exist, or an
// IO error trying to save the updated status.

// Preconditions:   (1) The database file must exist and be a valid
//
//					(2) The item must exist in the DB
//	    				because we use the item.Id as the key, this
//						function must check if the item already
//	    				exists in the DB, if not, return an error
//
// Postconditions:
//
//	 (1) The items status in the database will be updated
//		(2) If there is an error, it will be returned.
//		(3) This function MUST use existing functionality for most of its
//			work.  For example, it should call GetItem() to get the item
//			from the DB, then it should call UpdateItem() to update the
//			item in the DB (after the status is changed).
func (t *ToDo) ChangeItemDoneStatus(id int, value bool) error {

	//update was successful
	return errors.New("not implemented")
}

// GetAllItems returns all items from the DB.  If successful it
// returns a slice of all of the items to the caller
// Preconditions:   (1) The database file must exist and be a valid
//
// Postconditions:
//
//	 (1) All items will be returned, if any exist
//		(2) If there is an error, it will be returned
//			along with an empty slice
//		(3) The database file will not be modified
func (t *ToDo) GetAllItems() ([]ToDoItem, error) {
	keyList, err := t.getAllKeys()
	if err != nil {
		return nil, err
	}

	//preallocate the slice, will make things faster
	resList := make([]ToDoItem, len(keyList))

	for idx, k := range keyList {
		err := t.getItemFromRedis(k, &resList[idx])
		if err != nil {
			return nil, err
		}
	}

	return resList, nil
}

// PrintItem accepts a ToDoItem and prints it to the console
// in a JSON pretty format. As some help, look at the
// json.MarshalIndent() function from our in class go tutorial.
func (t *ToDo) PrintItem(item ToDoItem) {
	jsonBytes, _ := json.MarshalIndent(item, "", "  ")
	fmt.Println(string(jsonBytes))
}

// PrintAllItems accepts a slice of ToDoItems and prints them to the console
// in a JSON pretty format.  It should call PrintItem() to print each item
// versus repeating the code.
func (t *ToDo) PrintAllItems(itemList []ToDoItem) {
	for _, item := range itemList {
		t.PrintItem(item)
	}
}

// JsonToItem accepts a json string and returns a ToDoItem
// This is helpful because the CLI accepts todo items for insertion
// and updates in JSON format.  We need to convert it to a ToDoItem
// struct to perform any operations on it.
func (t *ToDo) JsonToItem(jsonString string) (ToDoItem, error) {
	var item ToDoItem
	err := json.Unmarshal([]byte(jsonString), &item)
	if err != nil {
		return ToDoItem{}, err
	}

	return item, nil
}
