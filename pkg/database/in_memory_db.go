package database

type InMemoryDB[T any] struct {
	data []*T
}

func (db *InMemoryDB[T]) Insert(item *T) {
	db.data = append(db.data, item)
}

func (db *InMemoryDB[T]) GetAll() []*T {
	return db.data
}

func (db *InMemoryDB[T]) Remove(item *T) {
	for i, v := range db.data {
		if v == item {
			db.data = append(db.data[:i], db.data[i+1:]...)
			return
		}
	}
}
