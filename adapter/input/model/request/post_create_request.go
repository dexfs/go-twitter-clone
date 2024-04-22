package request

type CreatePostRequest struct {
	UserID  string `json:"user_id" binding:"required,uuid4"`
	Content string `json:"content" binding:"required,min=10,max=255"`
}
