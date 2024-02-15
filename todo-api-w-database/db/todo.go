package db

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ToDoItem is the struct that represents a single ToDo item
type ToDoItem struct {
	Id     int    `json:"id" bson:"id"`
	Title  string `json:"title" bson:"title"`
	IsDone bool   `json:"done" bson:"done"`
}

const (
	MongoDefaultLocation = "mongodb://0.0.0.0:27017"
)

type cache struct {
	cli *mongo.Client
	ctx context.Context
}

// ToDo is the struct that represents the main object of our
// todo app.  It contains a reference to a cache object
type ToDo struct {
	//more things would be included in a real implementation

	//Mongo connections
	cache
}

// New is a constructor function that returns a pointer to a new
// ToDo struct.  If this is called it uses the default Redis URL
// with the companion constructor NewWithCacheInstance.
func New() (*ToDo, error) {
	//We will use an override if the REDIS_URL is provided as an environment
	//variable, which is the preferred way to wire up a docker container
	dbUrl := os.Getenv("MONGO_URL")
	//This handles the default condition
	if dbUrl == "" {
		dbUrl = MongoDefaultLocation
	}
	return NewWithDbInstance(dbUrl)
}

// NewWithCacheInstance is a constructor function that returns a pointer to a new
// ToDo struct.  It accepts a string that represents the location of the redis
// cache.
func NewWithDbInstance(dbURI string) (*ToDo, error) {
	ctx := context.TODO()

	//Connect to the database, 10 second timeout
	c, err := mongo.Connect(ctx, options.Client().
		ApplyURI(dbURI).
		SetServerSelectionTimeout(10*time.Second))
	if err != nil {
		return nil, err
	}

	//Mongo is normally lazy about connecting, so we will use ping
	//to ensure we are connected before moving on
	err = c.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	//Do some database setup, specifically set indexes
	err = dbSetup(ctx, c)
	if err != nil {
		return nil, err
	}

	//Return a pointer to a new ToDo struct
	return &ToDo{
		cache: cache{
			cli: c,
			ctx: ctx,
		},
	}, nil
}

// ------------------------------------------------------------
// Database HELPERS
// ------------------------------------------------------------
func dbSetup(ctx context.Context, c *mongo.Client) error {
	err := makeIndex(ctx, c, "todos", "id")
	if err != nil {
		return errors.New("Error creating index on polls collection: " + err.Error())
	}

	return nil
}

func makeIndex(ctx context.Context, c *mongo.Client, collection, field string, addFields ...string) error {

	//Create indexes where needed, note mongo will ignore indexes that already exist so
	//this is safe to run on every startup

	pollColl := c.Database("todoDB").Collection(collection)

	keys := bson.D{{Key: field, Value: 1}}
	for _, af := range addFields {
		keys = append(keys, bson.E{Key: af, Value: 1})
	}

	model := mongo.IndexModel{
		Keys:    keys,
		Options: options.Index().SetUnique(true),
	}

	_, err := pollColl.Indexes().CreateOne(ctx, model)
	if err != nil {
		return err
	}

	return nil
}

func (t *ToDo) DoesToDoExist(id int) bool {
	todoColl := t.cli.Database("todoDB").Collection("todos")

	query := bson.D{{Key: "id", Value: id}}
	opts := options.FindOne().SetProjection(bson.D{{Key: "id", Value: 1}})

	var vid any

	if err := todoColl.FindOne(context.TODO(), query, opts).Decode(&vid); err != nil {
		return false
	}

	return true
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

	if t.DoesToDoExist(item.Id) {
		return fmt.Errorf("ToDo item with id %d already exists", item.Id)
	}

	todoColl := t.cli.Database("todoDB").Collection("todos")

	res, err := todoColl.InsertOne(t.ctx, item)

	if err != nil {
		return err
	}

	log.Printf("Inserted voter document with _id: %v\n", res.InsertedID)

	return nil
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
	if t.DoesToDoExist(id) {
		return fmt.Errorf("ToDo item with id %d already exists", id)
	}

	todoColl := t.cli.Database("todoDB").Collection("todos")
	query := bson.D{{Key: "id", Value: id}}

	res, err := todoColl.DeleteOne(t.ctx, query)

	if err != nil {
		return err
	}

	log.Printf("Deleted count: %v\n", res.DeletedCount)

	return nil
}

// DeleteAll removes all items from the DB.
// It will be exposed via a DELETE /todo endpoint
func (t *ToDo) DeleteAll() (int, error) {

	todoColl := t.cli.Database("todoDB").Collection("todos")
	query := bson.D{}

	res, err := todoColl.DeleteMany(t.ctx, query)

	if err != nil {
		return 0, err
	}

	log.Printf("Deleted count: %v\n", res.DeletedCount)

	return int(res.DeletedCount), nil
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
	if !t.DoesToDoExist(item.Id) {
		return fmt.Errorf("ToDo item with id %d does not exist", item.Id)
	}

	todoColl := t.cli.Database("todoDB").Collection("todos")

	query := bson.D{{Key: "id", Value: item.Id}}
	res, err := todoColl.ReplaceOne(t.ctx, query, item)

	if err != nil {
		return err
	}

	if res.UpsertedCount > 0 {
		log.Printf("replaced voter document with _id: %v\n", res.UpsertedID)
	}
	return nil
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
	todoColl := t.cli.Database("todoDB").Collection("todos")

	query := bson.D{{Key: "id", Value: id}}
	err := todoColl.FindOne(t.ctx, query).Decode(newToDo)

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
	todoColl := t.cli.Database("todoDB").Collection("todos")

	cursor, err := todoColl.Find(t.ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(t.ctx)

	var items []ToDoItem
	if err = cursor.All(t.ctx, &items); err != nil {
		return nil, err
	}

	return items, nil
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
