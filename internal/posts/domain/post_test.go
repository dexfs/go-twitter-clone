package domain

import (
	"github.com/dexfs/go-twitter-clone/internal/user/domain"
	"testing"
)

func TestShouldInitializeAPostCorrectly(t *testing.T) {
	user := domain.NewUser("user post 1")
	mockInput := NewPostInput{User: user, Content: "mock_content"}
	newPost := NewPost(mockInput)

	if newPost == nil {
		t.Errorf("Invalid instance of Post")
	}

	if newPost.User != user {
		t.Errorf("got %q want %q", newPost.User, user)
	}

	if newPost.IsQuote != false || newPost.IsRepost != false {
		t.Errorf("got IsQuote %v or IsRepost %v want false for both", newPost.IsQuote, newPost.IsRepost)
	}

	if newPost.OriginalPostID != "" ||
		newPost.OriginalPostContent != "" ||
		newPost.OriginalPostUserID != "" ||
		newPost.OriginalPostScreenName != "" {
		t.Errorf("One or more fields are filled when they shouldn't be")
	}
}

func TestShouldInitializeARepostCorrectly(t *testing.T) {
	mockUser := domain.NewUser("post_original_user")
	mockUserRepost := domain.NewUser("post_repost_user")
	mockPostInput := NewPostInput{
		User:    mockUser,
		Content: "post_original_content",
	}
	mockOriginalPost := NewPost(mockPostInput)
	mockInput := &NewRepostQuoteInput{
		User: mockUserRepost,
		Post: mockOriginalPost,
	}

	newRepost, err := NewRepost(mockInput)

	if err != nil {
		t.Errorf("Unexpected error. %v", err)
	}

	if newRepost.User.ID != mockUserRepost.ID {
		t.Errorf("Expected repost user ID to be 'user_id', but got '%s'", newRepost.User.ID)
	}
	if newRepost.IsRepost != true {
		t.Errorf("Expected IsRepost to be true, but got false")
	}
	if newRepost.OriginalPostID != mockOriginalPost.ID {
		t.Errorf("Expected OriginalPostID to be 'original_post_id', but got '%s'", newRepost.OriginalPostID)
	}
}

func TestShouldInitializeAQuoteCorrectly(t *testing.T) {
	mockePostUser := domain.NewUser("post_original_user")
	mockQuotePostUser := domain.NewUser("post_user_user")
	mockPostInput := NewPostInput{
		User:    mockePostUser,
		Content: "post_original_content",
	}
	mockOriginalPost := NewPost(mockPostInput)

	mockInput := &NewRepostQuoteInput{
		User:    mockQuotePostUser,
		Post:    mockOriginalPost,
		Content: "post_quote_content",
	}

	newQuotePost, err := NewQuote(mockInput)

	if err != nil {
		t.Errorf("Unexpected error. %v", err)
	}

	if newQuotePost.User.ID != mockQuotePostUser.ID {
		t.Errorf("Expected repost user ID to be 'user_id', but got '%s'", newQuotePost.User.ID)
	}

	if newQuotePost.IsQuote != true {
		t.Errorf("Expected IsRepost to be true, but got false")
	}

	if newQuotePost.Content != "post_quote_content" {
		t.Errorf("Expected Content to be 'post_quote_content', but got '%s'", newQuotePost.Content)
	}

	if newQuotePost.OriginalPostID != mockOriginalPost.ID {
		t.Errorf("Expected OriginalPostID to be '%s', but got '%s'", mockOriginalPost.ID, newQuotePost.OriginalPostID)
	}

	if newQuotePost.OriginalPostContent != mockOriginalPost.Content {
		t.Errorf("Expected OriginalPostContent to be '%s', but got '%s'", mockOriginalPost.Content, newQuotePost.OriginalPostContent)
	}

	if newQuotePost.OriginalPostUserID != mockOriginalPost.User.ID {
		t.Errorf("Expected OriginalPostUserID to be '%s', but got '%s'", mockOriginalPost.User.ID, newQuotePost.OriginalPostUserID)
	}

	if newQuotePost.OriginalPostScreenName != mockOriginalPost.User.Username {
		t.Errorf("Expected OriginalPostScreenName to be '%s', but got '%s'", mockOriginalPost.User.Username, newQuotePost.OriginalPostScreenName)
	}

}
