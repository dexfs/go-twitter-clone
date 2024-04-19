package mocks

import (
	"github.com/dexfs/go-twitter-clone/internal/domain"
	"github.com/dexfs/go-twitter-clone/internal/infra/repository/in_memory"
	"github.com/dexfs/go-twitter-clone/pkg/database"
	"strconv"
)

// mocks
func MakeDb[T any]() *database.InMemoryDB[T] {
	return &database.InMemoryDB[T]{}
}
func MakeInMemoryUserRepo(db *database.InMemoryDB[domain.User]) domain.UserRepository {
	repo := in_memory.NewInMemoryUserRepo(db)
	return repo
}
func MakeInMemoryPostRepo(db *database.InMemoryDB[domain.Post]) domain.PostRepository {
	repo := in_memory.NewInMemoryPostRepo(db)
	return repo
}
func UserSeed(db *database.InMemoryDB[domain.User], amount int) []*domain.User {
	if amount <= 0 {
		amount = 1
	}
	users := make([]*domain.User, amount)
	for i := 0; i < len(users); i++ {
		username := "user" + strconv.Itoa(i)
		newUser := domain.NewUser(username)
		db.Insert(newUser)
		users[i] = newUser
	}
	return users
}
func PostSeed(db *database.InMemoryDB[domain.Post], user *domain.User, amount int) []*domain.Post {
	posts := make([]*domain.Post, amount)
	for i := 0; i < len(posts); i++ {
		newPostInput := domain.NewPostInput{
			User:    user,
			Content: "post_" + strconv.Itoa(i),
		}
		newPost, _ := domain.NewPost(newPostInput)
		db.Insert(newPost)
		posts[i] = newPost
	}
	return posts
}

type TestMocks struct {
	MockUserDB    *database.InMemoryDB[domain.User]
	MockUserSeed  []*domain.User
	MockPostDB    *database.InMemoryDB[domain.Post]
	MockPostsSeed []*domain.Post
}

func GetTestMocks() TestMocks {
	mockUserDB := MakeDb[domain.User]()
	mockPostDB := MakeDb[domain.Post]()
	mockUserSeed := UserSeed(mockUserDB, 1)
	mockPostsSeed := PostSeed(mockPostDB, mockUserSeed[0], 2)

	return TestMocks{
		MockUserDB:    mockUserDB,
		MockUserSeed:  mockUserSeed,
		MockPostDB:    mockPostDB,
		MockPostsSeed: mockPostsSeed,
	}
}
