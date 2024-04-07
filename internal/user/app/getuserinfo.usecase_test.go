package app

import (
	userEntity "github.com/dexfs/go-twitter-clone/internal/user"
	"github.com/dexfs/go-twitter-clone/internal/user/domain/interfaces"
	"github.com/dexfs/go-twitter-clone/internal/user/infra/db/repository"
	"github.com/dexfs/go-twitter-clone/pkg/database"
	"strconv"
	"testing"
)

func TestGetUserInfoUseCase_WithValidUsername_ReturnsUserInfo(t *testing.T) {
	inMemoryDb := MakeDb()
	usersSeed := UserSeed(inMemoryDb)
	userRepo := MakeRepoInstance(inMemoryDb)

	getInfoUseCase, _ := NewGetUserFeedUseCase(userRepo)
	output, err := getInfoUseCase.Execute(usersSeed[0].Username)
	if err != nil {
		t.Errorf("error while executing getInfoUseCase: %v", err)
	}

	if output.info != usersSeed[0] {
		t.Errorf("getInfoUseCase returned wrong user info, got %v, expected %v", output.info, usersSeed[0])
	}
}
func TestGetUserInfoUseCase_WithNonExistingUsername_ReturnsError(t *testing.T) {
	inMemoryDb := MakeDb()
	userRepo := MakeRepoInstance(inMemoryDb)

	getInfoUseCase, _ := NewGetUserFeedUseCase(userRepo)
	output, err := getInfoUseCase.Execute("")
	if err == nil {
		t.Errorf("should return error")
	}

	if output.info != nil {
		t.Errorf("should return empty user info")
	}
}
func TestGetUserInfoUseCase_WithNilUserRepository_ReturnsError(t *testing.T) {
	_, err := NewGetUserFeedUseCase(nil)
	if err == nil {
		t.Errorf("should return error")
	}

	if err.Error() != "userRepo cannot be nil" {
		t.Errorf("should return 'userRepo cannot be nil' got %v", err.Error())
	}
}

// helpers
func MakeDb() *database.InMemoryDB[userEntity.User] {
	return &database.InMemoryDB[userEntity.User]{}
}
func MakeRepoInstance(db *database.InMemoryDB[userEntity.User]) interfaces.UserRepository {
	repo := repository.NewInMemoryUserRepo(db)
	return repo
}
func UserSeed(db *database.InMemoryDB[userEntity.User]) []*userEntity.User {
	users := make([]*userEntity.User, 5)
	for i := 0; i < 5; i++ {
		username := "user" + strconv.Itoa(i)
		newUser := userEntity.NewUser(username)
		db.Insert(newUser)
		users[i] = newUser
	}
	return users
}