package request

type QuoteRequest struct {
	UserID string `json:"user_id" binding:"required,uuid4"`
	PostID string `json:"post_id" binding:"required,uuid4"`
	Quote  string `json:"quote" binding:"required,min=10,max=100"`
}
