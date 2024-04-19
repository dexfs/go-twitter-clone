package post_repo

import (
	"github.com/dexfs/go-twitter-clone/internal/post"
	"github.com/dexfs/go-twitter-clone/internal/user"
	"github.com/dexfs/go-twitter-clone/pkg/database"
	"math/rand"
	"testing"
)

func TestShouldInsertAPost(t *testing.T) {
	//Given
	userTest := user.NewUser("post_user_test")
	db := &database.InMemoryDB[post.Post]{}
	postRepo := NewInMemoryPostRepo(db)
	newPostInput := post.NewPostInput{
		User:    userTest,
		Content: "post test",
	}
	newPost, _ := post.NewPost(newPostInput)
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
	userTest := user.NewUser("post_user_test")
	db := &database.InMemoryDB[post.Post]{}
	postRepo := NewInMemoryPostRepo(db)
	newPostInput := post.NewPostInput{
		User:    userTest,
		Content: "post test",
	}
	newPostInput2 := post.NewPostInput{
		User:    userTest,
		Content: "post2 test",
	}
	newPost, _ := post.NewPost(newPostInput)
	post.NewPost(newPostInput2)

	postRepo.Insert(newPost)
	post, err := postRepo.FindByID(newPost.ID)

	if err != nil {
		t.Errorf("got %v want no empty", post)
	}

	if post == nil {
		t.Errorf("got nil want post")
	}

	if newPost.ID != post.ID {
		t.Errorf("got %v want %v", post.ID, newPost.ID)
	}

	post, err = postRepo.FindByID("not_found_id")
	if err == nil {

		t.Errorf("got %v want empty", post)
	}

}

func TestShouldRemoveAPost(t *testing.T) {
	//Given
	userTest := user.NewUser("post_user_test")
	db := &database.InMemoryDB[post.Post]{}
	postRepo := NewInMemoryPostRepo(db)
	newPostInput := post.NewPostInput{
		User:    userTest,
		Content: "post test",
	}
	newPostInput2 := post.NewPostInput{
		User:    userTest,
		Content: "post2 test",
	}
	newPost, _ := post.NewPost(newPostInput)
	newPost2, _ := post.NewPost(newPostInput2)

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
	userTest := user.NewUser("post_user_test")
	db := &database.InMemoryDB[post.Post]{}
	postRepo := NewInMemoryPostRepo(db)
	newPostInput := post.NewPostInput{
		User:    userTest,
		Content: "post test",
	}
	newPostInput2 := post.NewPostInput{
		User:    userTest,
		Content: "post2 test",
	}
	newPost, _ := post.NewPost(newPostInput)
	newPost2, _ := post.NewPost(newPostInput2)

	// When
	postRepo.Insert(newPost)
	postRepo.Insert(newPost2)
	countPosts := postRepo.CountByUser(userTest.ID)

	// Then
	expected := post.Count(2)
	if countPosts != expected {
		t.Errorf("got %v want %v", countPosts, expected)
	}
}

func TestShouldValidateHasReachedPostingLimitDay(t *testing.T) {
	//Given
	userTest := user.NewUser("post_user_test")
	db := &database.InMemoryDB[post.Post]{}
	postRepo := NewInMemoryPostRepo(db)
	count := 5
	for i := 0; i < count; i++ {
		newPost, _ := post.NewPost(post.NewPostInput{
			User:    userTest,
			Content: generateRandomString(10),
		})
		postRepo.Insert(newPost)
	}

	// When
	hasReached := postRepo.HasReachedPostingLimitDay(userTest.ID, uint64(count))

	if !hasReached {
		t.Errorf("got %v want %v", hasReached, true)
	}

	hasReached = postRepo.HasReachedPostingLimitDay(userTest.ID, uint64(10))

	if hasReached {
		t.Errorf("got %v want %v", hasReached, false)
	}
}

func TestShouldVerifyIfAPostIsEligibleForRepost(t *testing.T) {
	//Given
	mockUser := user.NewUser("post_user_test")
	mockRepostUser := user.NewUser("repost_user_test")
	mockOrigionalPostInput := post.NewPostInput{
		User:    mockUser,
		Content: "original_post",
	}
	mockOriginalPost, _ := post.NewPost(mockOrigionalPostInput)
	mockRepostInput := post.NewRepostQuoteInput{
		User:    mockRepostUser,
		Post:    mockOriginalPost,
		Content: "repost",
	}
	mockRepost, _ := post.NewRepost(mockRepostInput)

	db := &database.InMemoryDB[post.Post]{}
	db.Insert(mockOriginalPost)
	db.Insert(mockRepost)
	postRepo := NewInMemoryPostRepo(db)

	// When
	// Then
	hasPostBeenRepostedByUserRepost := postRepo.HasPostBeenRepostedByUser(mockOriginalPost.ID, mockRepostUser.ID)

	if !hasPostBeenRepostedByUserRepost {
		t.Errorf("got %v want %v", hasPostBeenRepostedByUserRepost, true)
	}

	hasPostBeenRepostedByUser := postRepo.HasPostBeenRepostedByUser(mockOriginalPost.ID, mockUser.ID)

	if hasPostBeenRepostedByUser {
		t.Errorf("got %v want %v", hasPostBeenRepostedByUserRepost, false)
	}

}

// utils
func generateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}
