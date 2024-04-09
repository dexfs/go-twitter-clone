package mocks

import (
	postEntity "github.com/dexfs/go-twitter-clone/internal/posts"
	postDomainInterfaces "github.com/dexfs/go-twitter-clone/internal/posts/domain/interfaces"
	postInfraImpl "github.com/dexfs/go-twitter-clone/internal/posts/infra/db/repository"
	userEntity "github.com/dexfs/go-twitter-clone/internal/user"
	"github.com/dexfs/go-twitter-clone/internal/user/domain/interfaces"
	"github.com/dexfs/go-twitter-clone/internal/user/infra/db/repository"
	"github.com/dexfs/go-twitter-clone/pkg/database"
	"strconv"
)

// mocks
func MakeDb[T any]() *database.InMemoryDB[T] {
	return &database.InMemoryDB[T]{}
}
func MakeInMemoryUserRepo(db *database.InMemoryDB[userEntity.User]) interfaces.UserRepository {
	repo := repository.NewInMemoryUserRepo(db)
	return repo
}
func MakeInMemoryPostRepo(db *database.InMemoryDB[postEntity.Post]) postDomainInterfaces.PostRepository {
	repo := postInfraImpl.NewInMemoryPostRepo(db)
	return repo
}
func UserSeed(db *database.InMemoryDB[userEntity.User], amount int) []*userEntity.User {
	if amount <= 0 {
		amount = 1
	}
	users := make([]*userEntity.User, amount)
	for i := 0; i < len(users); i++ {
		username := "user" + strconv.Itoa(i)
		newUser := userEntity.NewUser(username)
		db.Insert(newUser)
		users[i] = newUser
	}
	return users
}
func PostSeed(db *database.InMemoryDB[postEntity.Post], user *userEntity.User, amount int) []*postEntity.Post {
	posts := make([]*postEntity.Post, amount)
	for i := 0; i < len(posts); i++ {
		newPostInput := postEntity.NewPostInput{
			User:    user,
			Content: "post_" + strconv.Itoa(i),
		}
		newPost, _ := postEntity.NewPost(newPostInput)
		db.Insert(newPost)
		posts[i] = newPost
	}
	return posts
}

type TestMocks struct {
	MockUserDB    *database.InMemoryDB[userEntity.User]
	MockUserSeed  []*userEntity.User
	MockPostDB    *database.InMemoryDB[postEntity.Post]
	MockPostsSeed []*postEntity.Post
}

func GetTestMocks() TestMocks {
	mockUserDB := MakeDb[userEntity.User]()
	mockPostDB := MakeDb[postEntity.Post]()
	mockUserSeed := UserSeed(mockUserDB, 1)
	mockPostsSeed := PostSeed(mockPostDB, mockUserSeed[0], 2)

	return TestMocks{
		MockUserDB:    mockUserDB,
		MockUserSeed:  mockUserSeed,
		MockPostDB:    mockPostDB,
		MockPostsSeed: mockPostsSeed,
	}
}
