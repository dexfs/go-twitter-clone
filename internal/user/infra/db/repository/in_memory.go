package repository

import (
	"errors"
	user_domain "github.com/dexfs/go-twitter-clone/internal/user/domain"
	"github.com/dexfs/go-twitter-clone/pkg/database"
)

type UserDb struct {
	db *database.InMemoryDB[user_domain.User]
}

func NewUserInMemoryRepo(db *database.InMemoryDB[user_domain.User]) *UserDb {
	return &UserDb{
		db: db,
	}
}

func (u *UserDb) ByUsername(username string) (*user_domain.User, error) {
	for _, user := range u.db.GetAll() {
		if user.Username == username {
			return user, nil
		}
	}

	return nil, errors.New("user not found")
}

func (u *UserDb) Insert(item *user_domain.User) {
	u.db.Insert(item)
}

func (u *UserDb) GetAll() []*user_domain.User {
	return u.db.GetAll()
}

func (u *UserDb) Remove(item *user_domain.User) {
	u.db.Remove(item)
}
