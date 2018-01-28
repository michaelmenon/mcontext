/*
Context store to pass data between HTTP middlewares in Golang.
Its a thread safe context store.
*/

// mcontext
package mcontext

import (
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

	if key == nil || value == nil {
		return
	}
	mutex.Lock()

	contextStore[key] = value

	mutex.Unlock()
}

//get the corresponidng value of a key passed to this fucntion
func Get(key interface{}) (interface{}, bool) {

	if key == nil {
		return nil, false
	}
	mutex.Lock()
	defer mutex.Unlock()
	if value, ok := contextStore[key]; ok {

		return value, true
	} else {

		return nil, false
	}
}

//delete a aprticualr key,value from the context
func Delete(key interface{}) {
	if key == nil {
		return
	}
	mutex.Lock()
	delete(contextStore, key)
	mutex.Unlock()
}

//Delete all the entries fro mthe context
func Clear() {
	mutex.Lock()
	for key := range contextStore {
		delete(contextStore, key)
	}
	mutex.Unlock()
}
