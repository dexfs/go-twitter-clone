package repository

import (
	postEntity "github.com/dexfs/go-twitter-clone/internal/posts"
	userEntity "github.com/dexfs/go-twitter-clone/internal/user"
	"github.com/dexfs/go-twitter-clone/pkg/database"
	"math/rand"
	"testing"
)

func TestShouldInsertAPost(t *testing.T) {
	//Given
	userTest := userEntity.NewUser("post_user_test")
	db := &database.InMemoryDB[postEntity.Post]{}
	postRepo := NewPostInMemory(db)
	newPostInput := postEntity.NewPostInput{
		User:    userTest,
		Content: "post test",
	}
	newPost := postEntity.NewPost(newPostInput)
	// When
	postRepo.Insert(newPost)
	posts := postRepo.GetAll()

	// Then
	if len(posts) <= 0 {
		t.Errorf("got %v want 1", len(posts))
	}
}

func TestShouldFindAPostByID(t *testing.T) {
	//Given
	userTest := userEntity.NewUser("post_user_test")
	db := &database.InMemoryDB[postEntity.Post]{}
	postRepo := NewPostInMemory(db)
	newPostInput := postEntity.NewPostInput{
		User:    userTest,
		Content: "post test",
	}
	newPostInput2 := postEntity.NewPostInput{
		User:    userTest,
		Content: "post2 test",
	}
	newPost := postEntity.NewPost(newPostInput)
	postEntity.NewPost(newPostInput2)
	// When
	postRepo.Insert(newPost)
	post, err := postRepo.FindByID(newPost.ID)

	// Then

	if err != nil {
		t.Errorf("got %v want no empty", post)
	}

	if post.ID != newPost.ID {
		t.Errorf("got %v want %v", post.ID, newPost.ID)
	}
}

func TestShouldRemoveAPost(t *testing.T) {
	//Given
	userTest := userEntity.NewUser("post_user_test")
	db := &database.InMemoryDB[postEntity.Post]{}
	postRepo := NewPostInMemory(db)
	newPostInput := postEntity.NewPostInput{
		User:    userTest,
		Content: "post test",
	}
	newPostInput2 := postEntity.NewPostInput{
		User:    userTest,
		Content: "post2 test",
	}
	newPost := postEntity.NewPost(newPostInput)
	newPost2 := postEntity.NewPost(newPostInput2)

	// When
	postRepo.Insert(newPost)
	postRepo.Insert(newPost2)
	postRepo.Remove(newPost)
	posts := postRepo.GetAll()

	// Then
	expected := 1
	if len(posts) != expected {
		t.Errorf("got %v want %v", len(posts), expected)
	}
}

func TestShoulCountPostsPerUser(t *testing.T) {
	//Given
	userTest := userEntity.NewUser("post_user_test")
	db := &database.InMemoryDB[postEntity.Post]{}
	postRepo := NewPostInMemory(db)
	newPostInput := postEntity.NewPostInput{
		User:    userTest,
		Content: "post test",
	}
	newPostInput2 := postEntity.NewPostInput{
		User:    userTest,
		Content: "post2 test",
	}
	newPost := postEntity.NewPost(newPostInput)
	newPost2 := postEntity.NewPost(newPostInput2)

	// When
	postRepo.Insert(newPost)
	postRepo.Insert(newPost2)
	countPosts := postRepo.CountByUser(userTest.ID)

	// Then
	expected := Count(2)
	if countPosts != expected {
		t.Errorf("got %v want %v", countPosts, expected)
	}
}

func TestShouldValidateHasReachedPostingLimitDay(t *testing.T) {
	//Given
	userTest := userEntity.NewUser("post_user_test")
	db := &database.InMemoryDB[postEntity.Post]{}
	postRepo := NewPostInMemory(db)
	count := 5
	for i := 0; i < count; i++ {
		postRepo.Insert(postEntity.NewPost(postEntity.NewPostInput{
			User:    userTest,
			Content: GenerateRandomString(10),
		}))
	}

	// When
	hasReached := postRepo.HasReachedPostingLimitDay(userTest.ID, uint64(count))

	expect := true

	if hasReached != expect {
		t.Errorf("got %v want %v", hasReached, expect)
	}
}

func GenerateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}
