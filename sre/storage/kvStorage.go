package storage

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
