package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPutandGet(t *testing.T) {

	m := Map{}
	tables := []struct {
		x string
		y interface{}
	}{
		{"a", 1},
		{"b", 2.0},
		{"c", "harry"},
	}
	for _, table := range tables {
		m.Put(table.x, table.y)
		val, _ := m.Get(table.x)
		// assert.EqualError(t, err, nil)
		assert.Equal(t, val, table.y)
	}

}
func TestDelete(t *testing.T) {
	m := Map{}
	tables := []struct {
		x string
		y interface{}
	}{
		{"a", 1},
		{"b", 2.0},
		{"c", "harry"},
	}
	for _, table := range tables {
		m.Put(table.x, table.y)
		m.Delete(table.x)
		val, _ := m.Get(table.x)
		// assert.EqualErrorf(assert.TestingT, err, "value not found in the map")
		assert.Nil(t, val)
	}

}
