package request

type QuoteRequest struct {
	UserID string `json:"user_id" binding:"required,ulid"`
	PostID string `json:"post_id" binding:"required,ulid"`
	Quote  string `json:"quote" binding:"required,min=10,max=100"`
}
