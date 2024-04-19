package mocks

import (
	"github.com/dexfs/go-twitter-clone/internal/infra/post_repo"
	"github.com/dexfs/go-twitter-clone/internal/infra/user_repo"
	"github.com/dexfs/go-twitter-clone/internal/post"
	"github.com/dexfs/go-twitter-clone/internal/user"
	"github.com/dexfs/go-twitter-clone/pkg/database"
	"strconv"
)

// mocks
func MakeDb[T any]() *database.InMemoryDB[T] {
	return &database.InMemoryDB[T]{}
}
func MakeInMemoryUserRepo(db *database.InMemoryDB[user.User]) user.UserRepository {
	repo := user_repo.NewInMemoryUserRepo(db)
	return repo
}
func MakeInMemoryPostRepo(db *database.InMemoryDB[post.Post]) post.PostRepository {
	repo := post_repo.NewInMemoryPostRepo(db)
	return repo
}
func UserSeed(db *database.InMemoryDB[user.User], amount int) []*user.User {
	if amount <= 0 {
		amount = 1
	}
	users := make([]*user.User, amount)
	for i := 0; i < len(users); i++ {
		username := "user" + strconv.Itoa(i)
		newUser := user.NewUser(username)
		db.Insert(newUser)
		users[i] = newUser
	}
	return users
}
func PostSeed(db *database.InMemoryDB[post.Post], user *user.User, amount int) []*post.Post {
	posts := make([]*post.Post, amount)
	for i := 0; i < len(posts); i++ {
		newPostInput := post.NewPostInput{
			User:    user,
			Content: "post_" + strconv.Itoa(i),
		}
		newPost, _ := post.NewPost(newPostInput)
		db.Insert(newPost)
		posts[i] = newPost
	}
	return posts
}

type TestMocks struct {
	MockUserDB    *database.InMemoryDB[user.User]
	MockUserSeed  []*user.User
	MockPostDB    *database.InMemoryDB[post.Post]
	MockPostsSeed []*post.Post
}

func GetTestMocks() TestMocks {
	mockUserDB := MakeDb[user.User]()
	mockPostDB := MakeDb[post.Post]()
	mockUserSeed := UserSeed(mockUserDB, 1)
	mockPostsSeed := PostSeed(mockPostDB, mockUserSeed[0], 2)

	return TestMocks{
		MockUserDB:    mockUserDB,
		MockUserSeed:  mockUserSeed,
		MockPostDB:    mockPostDB,
		MockPostsSeed: mockPostsSeed,
	}
}
