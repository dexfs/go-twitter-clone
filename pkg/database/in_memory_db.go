package database

type InMemoryDB struct {
	Schemas map[string]interface{}
}

func NewInMemoryDB() *InMemoryDB {
	return &InMemoryDB{Schemas: make(map[string]interface{})}
}

func (db *InMemoryDB) GetSchema(key string) any {
	existing, ok := db.Schemas[key]
	if !ok {
		return nil
	}

	return existing
}

func (db *InMemoryDB) RegisterSchema(key string, value interface{}) {
	db.Schemas[key] = value
}

func (db *InMemoryDB) DropSchema(key string) {
	delete(db.Schemas, key)
}

//func (db *InMemoryDB) Insert(key string, newItem interface{}) {
//	existing, ok := db.Schemas[key].([]interface{})
//	if !ok {
//		log.Fatal("schema " + key + " not found")
//		return
//	}
//
//	updateSlice := append(existing, newItem.(*interface{}))
//	db.Schemas[key] = updateSlice
//}
