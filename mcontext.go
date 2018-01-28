/*
Context store to pass data between HTTP middlewares in Golang.
Its a thread safe context store.
*/

// mcontext
package mcontext

import (
	"log"

	"sync"
)

//context store as a map which stores keys/values
var contextStore map[interface{}]interface{}

//make it thread safe
var mutex *sync.Mutex

func init() {
	contextStore = make(map[interface{}]interface{})
	mutex = &sync.Mutex{}
}

//set a value to the context for a key value passed to this fucntion
func Set(key interface{}, value interface{}) {

	mutex.Lock()

	contextStore[key] = value

	mutex.Unlock()
}

//get the corresponidng value of a key passed to this fucntion
func Get(key interface{}) (interface{}, bool) {
	mutex.Lock()
	if value, ok := contextStore[key]; ok {
		log.Println("Found store....")

		mutex.Unlock()
		return value, true
	} else {
		mutex.Unlock()
		return nil, false
	}
}

//delete a aprticualr key,value from the context
func Erase(key interface{}) {
	delete(contextStore, key)
}

//Delete all the entries fro mthe context
func celar() {
	for key := range contextStore {
		delete(contextStore, key)
	}
}
