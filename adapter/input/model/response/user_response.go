package response

import "github.com/dexfs/go-twitter-clone/internal/core/domain"

type GetUserInfoResponse struct {
	ID        string `json:"id"`
	Username  string `json:"username"`
	CreatedAt string `json:"created_at"`
}

type GetUserFeedResponse struct {
	Items []*domain.Post `json:"items"`
}
