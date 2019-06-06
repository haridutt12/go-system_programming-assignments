package main

import (
	"fmt"
)

type Storage interface {
	Put(key string, value interface{})
	Get(key string) (interface{}, string)
	Delete(key string)
}

type Map map[string]interface{}

func (v Map) Put(key string, val interface{}) {
	v[key] = val
}

func (v Map) Get(key string) (interface{}, string) {
	val := v[key]
	val, ok := v[key]
	if !ok {
		return nil, "NULL"
	}
	return val, "NULL"
}

func (v Map) Delete(key string) {
	_, ok := v[key]
	// checks if key exists. and deletes content only if key exists. else ignores.
	if ok {
		delete(v, key)
	}

}

func main() {

	m := Map{}
	m.Put("a", 123)
	m.Put("b", 1.234)

	res, _ := m.Get("a")
	fmt.Println(res)

	res, _ = m.Get("b")

	fmt.Println(res)

	m.Delete("b")

	res, _ = m.Get("b")

	fmt.Println(res)

}
