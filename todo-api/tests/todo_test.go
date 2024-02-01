package tests

import (
	"log"
	"os"
	"testing"

	"drexel.edu/todo/db"
	fake "github.com/brianvoe/gofakeit/v6" //aliasing package name
	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
)

var (
	BASE_API = "http://localhost:1080"

	cli = resty.New()
)

func TestMain(m *testing.M) {

	//SETUP GOES FIRST
	rsp, err := cli.R().Delete(BASE_API + "/todo")

	if rsp.StatusCode() != 200 {
		log.Printf("error clearing database, %v", err)
		os.Exit(1)
	}

	code := m.Run()

	//CLEANUP

	//Now Exit
	os.Exit(code)
}

func newRandToDoItem(id int) db.ToDoItem {
	return db.ToDoItem{
		Id:     id,
		Title:  fake.Sentence(5),
		IsDone: fake.Bool(),
	}
}

func Test_LoadDB(t *testing.T) {
	numLoad := 3
	for i := 0; i < numLoad; i++ {
		item := newRandToDoItem(i)
		rsp, err := cli.R().
			SetBody(item).
			Post(BASE_API + "/todo")

		assert.Nil(t, err)
		assert.Equal(t, 200, rsp.StatusCode())
	}
}

func Test_GetAllItems(t *testing.T) {
	var items []db.ToDoItem

	rsp, err := cli.R().SetResult(&items).Get(BASE_API + "/todo")

	assert.Nil(t, err)
	assert.Equal(t, 200, rsp.StatusCode())

	assert.Equal(t, 3, len(items))
}

func Test_DeleteToDo(t *testing.T) {
	var item db.ToDoItem

	rsp, err := cli.R().SetResult(&item).Get(BASE_API + "/todo/2")
	assert.Nil(t, err)
	assert.Equal(t, 200, rsp.StatusCode(), "todo #2 expected")

	rsp, err = cli.R().Delete(BASE_API + "/todo/2")
	assert.Nil(t, err)
	assert.Equal(t, 200, rsp.StatusCode(), "todo not deleted expected")

	rsp, err = cli.R().SetResult(item).Get(BASE_API + "/todo/2")
	assert.Nil(t, err)
	assert.Equal(t, 404, rsp.StatusCode(), "expected not found error code")
}
