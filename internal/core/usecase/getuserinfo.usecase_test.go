package usecase_test

import (
	"github.com/dexfs/go-twitter-clone/adapter/output/repository/inmemory"
	"github.com/dexfs/go-twitter-clone/internal/core/usecase"
	"github.com/dexfs/go-twitter-clone/tests/mocks"
	"testing"
)

func TestGetUserInfoUseCase_WithValidUsername_ReturnsUserInfo(t *testing.T) {
	TestMocks := mocks.GetTestMocks()
	inMemoryDb := TestMocks.MockDB
	usersSeed := TestMocks.MockUserSeed
	userRepo := inmemory.NewInMemoryUserRepository(inMemoryDb)

	getInfoUseCase, _ := usecase.NewGetUserInfoUseCase(userRepo)
	output, err := getInfoUseCase.Execute(usersSeed[0].Username)
	if err != nil {
		t.Errorf("error while executing getInfoUseCase: %v", err)
	}

	if output.ID != usersSeed[0].ID {
		t.Errorf("getInfoUseCase returned wrong user info, got %v, expected %v", output, usersSeed[0])
	}
}
func TestGetUserInfoUseCase_WithNonExistingUsername_ReturnsError(t *testing.T) {
	TestMocks := mocks.GetTestMocks()
	inMemoryDb := TestMocks.MockDB
	userRepo := inmemory.NewInMemoryUserRepository(inMemoryDb)

	getInfoUseCase, _ := usecase.NewGetUserInfoUseCase(userRepo)
	output, err := getInfoUseCase.Execute("")
	if err == nil {
		t.Errorf("should return error")
	}

	if output != nil {
		t.Errorf("should return empty user info")
	}
}
func TestGetUserInfoUseCase_WithNilUserRepository_ReturnsError(t *testing.T) {
	_, err := usecase.NewGetUserInfoUseCase(nil)

	if err == nil {
		t.Errorf("should return error")
	}

	if err.Error() != "userPort cannot be nil" {
		t.Errorf("should return 'userPort cannot be nil' got %v", err.Error())
	}
}
