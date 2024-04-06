package database

type BaseDb[T any] interface {
	Insert(entity *T) *T
	FindByID(id string) *T
	Update(entity *T)
}
