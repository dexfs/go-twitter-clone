package in_memory

import (
	"github.com/dexfs/go-twitter-clone/internal/domain"
	"github.com/dexfs/go-twitter-clone/pkg/database"
	"strconv"
	"testing"
)

func TestShouldReturnInsertedUser(t *testing.T) {
	userTest := domain.NewUser("usuarion_test_1")

	db := &database.InMemoryDB[domain.User]{}
	userRepo := NewInMemoryUserRepo(db)

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
	userToFind := domain.NewUser("user_to_find")
	db := &database.InMemoryDB[domain.User]{}
	for i := 0; i < 5; i++ {
		newUser := domain.NewUser("username_" + strconv.Itoa(i))
		db.Insert(newUser)
	}
	db.Insert(userToFind)

	userRepo := NewInMemoryUserRepo(db)

	foundUser, _ := userRepo.ByUsername(userToFind.Username)

	if foundUser.Username != userToFind.Username {
		t.Errorf("got %v want %v", foundUser, userToFind.Username)
	}
}

func TestShouldRemoveUserByID(t *testing.T) {
	userToDelete := domain.NewUser("user_to_find")
	db := &database.InMemoryDB[domain.User]{}
	for i := 0; i < 5; i++ {
		newUser := domain.NewUser("username_" + strconv.Itoa(i))
		db.Insert(newUser)
	}
	db.Insert(userToDelete)

	userRepo := NewInMemoryUserRepo(db)

	userRepo.Remove(userToDelete)

	findByUserRemoved, err := userRepo.ByUsername(userToDelete.Username)

	if err == nil {
		t.Errorf("got %v want nil", findByUserRemoved)
	}
}
