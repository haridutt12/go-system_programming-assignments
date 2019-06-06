package storage

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStorer(t *testing.T) {

	m := Map{}
	tables := []struct {
		x string
		y interface{}
	}{
		{"a", 1},
		{"b", 2.0},
		{"c", "harry"},
	}

	t.Run("Should insert  key value ", func(t *testing.T) {
		for _, table := range tables {
			m.Put(table.x, table.y)
			val, _ := m.Get(table.x)
			assert.Equal(t, val, table.y)
		}

	})

	t.Run("Should return error for non existing key", func(t *testing.T) {

		for _, table := range tables {
			m.Put(table.x, table.y)
			_, err := m.Get(table.x)
			assert.Equal(t, err, "NULL")
		}

	})

	t.Run("Should return an error on non-existing key", func(t *testing.T) {

		for _, table := range tables {
			m.Put(table.x, table.y)
			m.Delete(table.x)
			_, err := m.Get(table.x)
			assert.Equal(t, err, "NULL")
		}

	})

	t.Run("Should delete existing key", func(t *testing.T) {

		for _, table := range tables {
			m.Put(table.x, table.y)
			m.Delete(table.x)
			val, _ := m.Get(table.x)
			assert.Nil(t, val)
		}

	})

}
