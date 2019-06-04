package main

import "fmt"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Storage interface {
	Put(key string, value interface{})
	Get(key string) interface{}
	Delete(key string)
}

type Map map[string]interface{}

func (v Map) Put(key string, val interface{}) {
	v[key] = val
}

func (v Map) Get(key string) interface{} {
	val := v[key]
	return val
}

func (v Map) Delete(key string) {
	v[key] = nil

}

func main() {

	m := Map{}
	m.Put("a", 123)
	m.Put("b", 1.234)

	res := m.Get("a")
	// check(err)
	fmt.Println(res)

	m.Delete("b")

	res = m.Get("b")
	// check(err)
	fmt.Println(res)

}
