package request

type CreatePostRequest struct {
	UserID  string `json:"user_id" binding:"required,ulid"`
	Content string `json:"content" binding:"required,min=10,max=255"`
}
