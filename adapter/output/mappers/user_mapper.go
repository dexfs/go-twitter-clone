package mappers

import (
	inmemory_schema "github.com/dexfs/go-twitter-clone/adapter/output/repository/inmemory/schema"
	"github.com/dexfs/go-twitter-clone/internal/core/domain"
)

type userMapper struct{}

func NewUserMapper() *userMapper {
	return &userMapper{}
}

func (m *userMapper) ToPersistence(aUser *domain.User) *inmemory_schema.UserSchema {
	return &inmemory_schema.UserSchema{
		ID:        aUser.ID,
		Username:  aUser.Username,
		CreatedAt: aUser.CreatedAt,
		UpdatedAt: aUser.UpdatedAt,
	}
}

func (m *userMapper) FromPersistence(aUserSchema *inmemory_schema.UserSchema) *domain.User {
	return &domain.User{
		ID:        aUserSchema.ID,
		Username:  aUserSchema.Username,
		CreatedAt: aUserSchema.CreatedAt,
		UpdatedAt: aUserSchema.UpdatedAt,
	}
}
