package repository

import (
	"github.com/dexfs/go-twitter-clone/internal/user"
	"github.com/dexfs/go-twitter-clone/pkg/database"
	"strconv"
	"testing"
)

func TestShouldReturnInsertedUser(t *testing.T) {
	userTest := user.NewUser("usuarion_test_1")

	db := &database.InMemoryDB[user.User]{}
	userRepo := NewUserInMemoryRepo(db)

	userRepo.Insert(userTest)

	users := userRepo.GetAll()

	if len(users) > 1 || len(users) < 1 {
		t.Errorf("got %d want 1", len(users))
	}

	if users[0].Username != userTest.Username {
		t.Errorf("got %v want %v", users[0], userTest)
	}
}

func TestShouldReturnUserByUsername(t *testing.T) {
	userToFind := user.NewUser("user_to_find")
	db := &database.InMemoryDB[user.User]{}
	for i := 0; i < 5; i++ {
		user := user.NewUser("username_" + strconv.Itoa(i))
		db.Insert(user)
	}
	db.Insert(userToFind)

	userRepo := NewUserInMemoryRepo(db)

	foundUser, _ := userRepo.ByUsername(userToFind.Username)

	if foundUser.Username != userToFind.Username {
		t.Errorf("got %v want %v", foundUser, userToFind.Username)
	}
}

func TestShouldRemoveUserByID(t *testing.T) {
	userToDelete := user.NewUser("user_to_find")
	db := &database.InMemoryDB[user.User]{}
	for i := 0; i < 5; i++ {
		user := user.NewUser("username_" + strconv.Itoa(i))
		db.Insert(user)
	}
	db.Insert(userToDelete)

	userRepo := NewUserInMemoryRepo(db)

	userRepo.Remove(userToDelete)

	findByUserRemoved, err := userRepo.ByUsername(userToDelete.Username)

	if err == nil {
		t.Errorf("got %v want nil", findByUserRemoved)
	}
}
