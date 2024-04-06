package repository

import (
	"errors"
	"github.com/dexfs/go-twitter-clone/internal/user"
	"github.com/dexfs/go-twitter-clone/pkg/database"
)

type UserDb struct {
	db *database.InMemoryDB[user.User]
}

func NewUserInMemoryRepo(db *database.InMemoryDB[user.User]) *UserDb {
	return &UserDb{
		db: db,
	}
}

func (u *UserDb) ByUsername(username string) (*user.User, error) {
	for _, user := range u.db.GetAll() {
		if user.Username == username {
			return user, nil
		}
	}

	return nil, errors.New("user not found")
}

func (u *UserDb) Insert(item *user.User) {
	u.db.Insert(item)
}

func (u *UserDb) GetAll() []*user.User {
	return u.db.GetAll()
}

func (u *UserDb) Remove(item *user.User) {
	u.db.Remove(item)
}
