package database

import "slices"

type InMemoryDB[T any] struct {
	data []*T
}

func (db *InMemoryDB[T]) Insert(item *T) {
	db.data = slices.Insert(db.data, len(db.data), item)
}

func (db *InMemoryDB[T]) GetAll() []*T {
	return db.data
}

func (db *InMemoryDB[T]) Remove(item *T) {
	for i, v := range db.data {
		if v == item {
			db.data = slices.Delete(db.data, i, i+1)
			return
		}
	}
}
