package main

type Storage interface {
	Put(key string, value interface{}) error
	Get(key string) (error, interface{})
	Delete(key string) error
}

type Vars map[string]interface{}


func (v *Vars)Put(key string, val interface{}) error{
     return map[string]interface{}(*v)[key] = val
}

func (v *Vars)Get(key string) (error, interface{}) {
	ok, val := map[string]interface{}(*v)[key]
}

func (v *Vars)Delete(key string) error {
	map[string]interface{}(*v)[key] = nil
}

func main() {
	

	

}
