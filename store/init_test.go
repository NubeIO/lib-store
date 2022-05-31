package store

import (
	"fmt"
	"testing"
)

var memStore = initStore()

func initStore() *Handler {
	return Init()
}

type Foo struct {
	Message string `json:"message"`
}

func SetStore(key string, data interface{}) {
	memStore.SetNoExpire(key, data)
}

func TestInit(t *testing.T) {

	key := "key"

	data, ok := memStore.Get(key)
	if ok {
		parse := data.(*Foo)
		fmt.Println("parse", parse)
	}
	resp := &Foo{
		Message: "not found",
	}

	fmt.Println("empty", resp)
	resp.Message = "added val"
	SetStore(key, resp)

	data, ok = memStore.Get(key)
	if ok {
		parse := data.(*Foo)
		fmt.Println("not-empty", parse)
	}

}
