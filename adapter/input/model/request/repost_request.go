package request

type RepostRequest struct {
	UserID string `json:"user_id" binding:"required,uuid4"`
	PostID string `json:"post_id" binding:"required,uuid4"`
}
