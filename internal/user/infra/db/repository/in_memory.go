package repository

import (
	"errors"
	userEntity "github.com/dexfs/go-twitter-clone/internal/user"
	"github.com/dexfs/go-twitter-clone/pkg/database"
)

type UserDb struct {
	db *database.InMemoryDB[userEntity.User]
}

func NewUserInMemoryRepo(db *database.InMemoryDB[userEntity.User]) *UserDb {
	return &UserDb{
		db: db,
	}
}

func (u *UserDb) ByUsername(username string) (*userEntity.User, error) {
	for _, currentUser := range u.db.GetAll() {
		if currentUser.Username == username {
			return currentUser, nil
		}
	}

	return nil, errors.New("currentUser not found")
}

func (u *UserDb) Insert(item *userEntity.User) {
	u.db.Insert(item)
}

func (u *UserDb) GetAll() []*userEntity.User {
	return u.db.GetAll()
}

func (u *UserDb) Remove(item *userEntity.User) {
	u.db.Remove(item)
}
