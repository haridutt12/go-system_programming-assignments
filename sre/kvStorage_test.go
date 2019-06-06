package main

import "testing"

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
		val := m.Get(table.x)
		if val != table.y {
			t.Errorf("Fetched incorrect value for key %s , got: %d, want: %d.", table.x, val, table.y)
		}
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
		val := m.Get(table.x)
		if val != nil {
			t.Errorf("Fetched incorrect value for key %s , got: %d, want: %s.", table.x, val, "null")
		}
	}

}
