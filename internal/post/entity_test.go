package post

import (
	UserModule "github.com/dexfs/go-twitter-clone/internal/user"
	"github.com/google/uuid"
	"testing"
	"time"
)

// Post
func TestNewPost_WithValidInput_ReturnsOK(t *testing.T) {
	user := UserModule.NewUser("user post 1")
	mockInput := NewPostInput{User: user, Content: "mock_content"}
	newPost, _ := NewPost(mockInput)

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
func TestNewPost_WithEmptyInput_ReturnsError(t *testing.T) {
	mockInput := NewPostInput{}
	_, err := NewPost(mockInput)

	if err == nil {
		t.Errorf("Invalid instance of Post")
	}
}
func TestNewPost_WithNilUser_ReturnsError(t *testing.T) {
	mockInput := NewPostInput{
		User: nil,
	}
	_, err := NewPost(mockInput)

	if err == nil {
		t.Errorf("Invalid instance of Post")
	}

	if err.Error() != "no user provided" {
		t.Errorf("got %q want %q", err.Error(), "no user provided")
	}

}
func TestNewPost_WithEmptyPostContent_ReturnsError(t *testing.T) {
	mockUser := UserModule.NewUser("test_user")
	mockInput := NewPostInput{
		User: mockUser,
	}
	_, err := NewPost(mockInput)

	if err == nil {
		t.Errorf("Invalid instance of Post")
	}

	if "no content provided" != err.Error() {
		t.Errorf("got %q want %q", err.Error(), "no content provided")
	}
}

// Repost
func TestNewRepost_WithValidInput_ReturnsOK(t *testing.T) {
	mockUser := UserModule.NewUser("post_original_user")
	mockUserRepost := UserModule.NewUser("post_repost_user")
	mockPostInput := NewPostInput{
		User:    mockUser,
		Content: "post_original_content",
	}
	mockOriginalPost, _ := NewPost(mockPostInput)
	mockInput := NewRepostQuoteInput{
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
func TestNewRepost_WithRepostPost_ReturnsError(t *testing.T) {
	mockOriginalPost := GenerateOriginalPost()
	mockRepost := GenerateRepost(mockOriginalPost)
	newRepostInput := NewRepostQuoteInput{
		User:    mockRepost.User,
		Post:    mockRepost,
		Content: "repost in test",
	}
	_, err := NewRepost(newRepostInput)

	if err == nil {
		t.Errorf("NewRepost should have returned an error")
	}

	expectedMsgError := "it is not possible repost a repost post"
	if err.Error() != expectedMsgError {
		t.Errorf("Returned error is not correct. got '%s' want '%s'", err.Error(), expectedMsgError)
	}

}
func TestNewRepost_WithSameUserID_ReturnsError(t *testing.T) {
	mockOriginalPost := GenerateOriginalPost()
	newRepostInput := NewRepostQuoteInput{
		User:    mockOriginalPost.User,
		Post:    mockOriginalPost,
		Content: "repost with the same user",
	}
	_, err := NewRepost(newRepostInput)

	if err == nil {
		t.Errorf("NewRepost should have returned an error")
	}

	expectedMsgError := "it is not possible repost your own post"
	if expectedMsgError != err.Error() {
		t.Errorf("Returned error is not correct. got '%s' want '%s'", err.Error(), expectedMsgError)
	}
}
func TestNewRepost_WithEmptyInput_ReturnsError(t *testing.T)       {}
func TestNewRepost_WithEmptyPostContent_ReturnsError(t *testing.T) {}

// Quotepost
func TestNewQuotepost_WithValidInput_ReturnsOK(t *testing.T) {
	mockePostUser := UserModule.NewUser("post_original_user")
	mockQuotePostUser := UserModule.NewUser("post_user_user")
	mockPostInput := NewPostInput{
		User:    mockePostUser,
		Content: "post_original_content",
	}
	mockOriginalPost, _ := NewPost(mockPostInput)

	mockInput := NewRepostQuoteInput{
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
func TestNewQuotepost_WithSameUserID_ReturnsError(t *testing.T) {
	originalPost := GenerateOriginalPost()
	mockInput := NewRepostQuoteInput{
		User:    originalPost.User,
		Post:    originalPost,
		Content: "repost in test",
	}

	_, err := NewQuote(mockInput)

	if err == nil {
		t.Errorf("Invalid instance of QuotePost returned")
	}

	if "it is not possible quote your own post" != err.Error() {
		t.Errorf("Returned error is not correct. got '%s' want '%s'", err.Error(), "it is not possible quote your own post")
	}
}
func TestNewQuotepost_WithQuotepost_ReturnsError(t *testing.T) {
	originalPost := GenerateOriginalPost()
	quotePost := GenerateQuotepost(originalPost)
	mockInput := NewRepostQuoteInput{
		User:    originalPost.User,
		Post:    quotePost,
		Content: "repost in test",
	}

	_, err := NewQuote(mockInput)

	if err == nil {
		t.Errorf("Invalid instance of QuotePost returned")
	}

	if err.Error() != "it is not possible a quote post of a quote post" {
		t.Errorf("Returned error is not correct. got '%s' want '%s'", err.Error(), "it is not possible a quote post of a quote post")
	}
}
func TestNewQuotepost_WithEmptyPostContent_ReturnsError(t *testing.T) {
	originalPost := GenerateOriginalPost()
	mockInput := NewRepostQuoteInput{
		User: originalPost.User,
		Post: originalPost,
	}

	_, err := NewQuote(mockInput)

	if err == nil {
		t.Errorf("Invalid instance of QuotePost returned")
	}

	if err.Error() != "no content provided" {
		t.Errorf("Returned error is not correct. got '%s' want '%s'", err.Error(), "no content provided")
	}
}
func TestNewQuotepost_WithNilPost_ReturnsError(t *testing.T) {
	mockInput := NewRepostQuoteInput{}

	_, err := NewQuote(mockInput)

	if err == nil {
		t.Errorf("Invalid instance of QuotePost returned")
	}

	if err.Error() != "no post provided" {
		t.Errorf("Returned error is not correct. got '%s' want '%s'", err.Error(), "no post provided")
	}
}
func TestNewQuotepost_WithNilUser_ReturnsError(t *testing.T) {
	mockOriginalPost := GenerateOriginalPost()
	mockInput := NewRepostQuoteInput{
		Post: mockOriginalPost,
	}

	_, err := NewQuote(mockInput)

	if err == nil {
		t.Errorf("Invalid instance of QuotePost returned")
	}

	if err.Error() != "no user provided" {
		t.Errorf("Returned error is not correct. got '%s' want '%s'", err.Error(), "no user provided")
	}
}

// in memory seeders
func GenerateOriginalPost() *Post {
	mockePostUser := UserModule.NewUser("post_original_user")
	mockPostInput := NewPostInput{
		User:    mockePostUser,
		Content: "post_original_content",
	}
	newPost, _ := NewPost(mockPostInput)
	return newPost
}
func GenerateRepost(anOriginalPost *Post) *Post {
	return &Post{
		ID:                     uuid.NewString(),
		User:                   anOriginalPost.User,
		Content:                anOriginalPost.Content,
		CreatedAt:              time.Now(),
		IsQuote:                false,
		IsRepost:               true,
		OriginalPostID:         anOriginalPost.ID,
		OriginalPostContent:    anOriginalPost.Content,
		OriginalPostUserID:     anOriginalPost.User.ID,
		OriginalPostScreenName: anOriginalPost.User.Username,
	}
}
func GenerateQuotepost(anOriginalPost *Post) *Post {
	return &Post{
		ID:                     uuid.NewString(),
		User:                   anOriginalPost.User,
		Content:                anOriginalPost.Content,
		CreatedAt:              time.Now(),
		IsQuote:                true,
		IsRepost:               false,
		OriginalPostID:         anOriginalPost.ID,
		OriginalPostContent:    anOriginalPost.Content,
		OriginalPostUserID:     anOriginalPost.User.ID,
		OriginalPostScreenName: anOriginalPost.User.Username,
	}
}
