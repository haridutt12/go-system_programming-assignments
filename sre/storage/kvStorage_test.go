package storage

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
		val, err := m.Get(table.x)
		assert.Equal(t, err, "NULL")
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
		val, err := m.Get(table.x)
		assert.Equal(t, err, "NULL")
		assert.Nil(t, val)
	}

}
