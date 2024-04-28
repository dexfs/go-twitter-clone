package mappers

import (
	inmemory_schema "github.com/dexfs/go-twitter-clone/adapter/output/repository/inmemory/schema"
	"github.com/dexfs/go-twitter-clone/internal/core/domain"
)

type postMapper struct{}

func NewPostMapper() *postMapper {
	return &postMapper{}
}

func (m *postMapper) ToPersistence(aPost *domain.Post) *inmemory_schema.PostSchema {
	return &inmemory_schema.PostSchema{
		ID:                     aPost.ID,
		UserID:                 aPost.User.ID,
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

func (m *postMapper) FromPersistence(aPost *inmemory_schema.PostSchema, aUser *domain.User) *domain.Post {
	return &domain.Post{
		ID:                     aPost.ID,
		User:                   aUser,
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
