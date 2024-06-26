package app

import (
	"github.com/dexfs/go-twitter-clone/internal/domain"
	"github.com/dexfs/go-twitter-clone/internal/domain/interfaces"
	"github.com/dexfs/go-twitter-clone/internal/infra/repository/inmemory"
	"github.com/dexfs/go-twitter-clone/pkg/database"
	"strconv"
	"testing"
)

func TestGetUserInfoUseCase_WithValidUsername_ReturnsUserInfo(t *testing.T) {
	inMemoryDb := MakeDb()
	usersSeed := UserSeed(inMemoryDb)
	userRepo := MakeRepoInstance(inMemoryDb)

	getInfoUseCase, _ := NewGetUserInfoUseCase(userRepo)
	output, err := getInfoUseCase.Execute(usersSeed[0].Username)
	if err != nil {
		t.Errorf("error while executing getInfoUseCase: %v", err)
	}

	if output.User != usersSeed[0] {
		t.Errorf("getInfoUseCase returned wrong user info, got %v, expected %v", output, usersSeed[0])
	}
}
func TestGetUserInfoUseCase_WithNonExistingUsername_ReturnsError(t *testing.T) {
	inMemoryDb := MakeDb()
	userRepo := MakeRepoInstance(inMemoryDb)

	getInfoUseCase, _ := NewGetUserInfoUseCase(userRepo)
	output, err := getInfoUseCase.Execute("")
	if err == nil {
		t.Errorf("should return error")
	}

	if output.User != nil {
		t.Errorf("should return empty user info")
	}
}
func TestGetUserInfoUseCase_WithNilUserRepository_ReturnsError(t *testing.T) {
	_, err := NewGetUserInfoUseCase(nil)
	if err == nil {
		t.Errorf("should return error")
	}

	if err.Error() != "userRepo cannot be nil" {
		t.Errorf("should return 'userRepo cannot be nil' got %v", err.Error())
	}
}

// mocks
func MakeDb() *database.InMemoryDB[domain.User] {
	return &database.InMemoryDB[domain.User]{}
}
func MakeRepoInstance(db *database.InMemoryDB[domain.User]) interfaces.UserRepository {
	repo := inmemory.NewInMemoryUserRepo(db)
	return repo
}
func UserSeed(db *database.InMemoryDB[domain.User]) []*domain.User {
	users := make([]*domain.User, 5)
	for i := 0; i < 5; i++ {
		username := "user" + strconv.Itoa(i)
		newUser := domain.NewUser(username)
		db.Insert(newUser)
		users[i] = newUser
	}
	return users
}
