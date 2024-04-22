package mappers

import (
	"github.com/dexfs/go-twitter-clone/adapter/output/repository/inmemory"
	"github.com/dexfs/go-twitter-clone/core/domain"
)

type postMapper struct{}

func NewPostMapper() *postMapper {
	return &postMapper{}
}

func (m *postMapper) ToPersistence(aPost *domain.Post) *inmemory.PostSchema {
	return &inmemory.PostSchema{
		ID:                     aPost.ID,
		UserId:                 aPost.UserID,
		Content:                aPost.Content,
		CreatedAt:              aPost.CreatedAt,
		IsQuote:                aPost.IsQuote,
		IsRepost:               aPost.IsRepost,
		OriginalPostID:         aPost.OriginalPostID,
		OriginalPostContent:    aPost.OriginalPostContent,
		OriginalPostUserID:     aPost.OriginalPostUserID,
		OriginalPostScreenName: aPost.OriginalPostScreenName,
	}
}
