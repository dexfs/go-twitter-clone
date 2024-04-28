package mocks

import (
	"github.com/dexfs/go-twitter-clone/adapter/output/mappers"
	"github.com/dexfs/go-twitter-clone/adapter/output/repository/inmemory"
	inmemory_schema "github.com/dexfs/go-twitter-clone/adapter/output/repository/inmemory/schema"
	"github.com/dexfs/go-twitter-clone/internal/core/domain"
	"github.com/dexfs/go-twitter-clone/internal/core/port/output"
	"github.com/dexfs/go-twitter-clone/pkg/database"
	"log"
	"strconv"
)

// mocks
func MakeDb() *database.InMemoryDB {
	db := database.NewInMemoryDB()
	return db
}
func MakeInMemoryUserRepo(db *database.InMemoryDB) output.UserPort {
	repo := inmemory.NewInMemoryUserRepository(db)
	return repo
}
func MakeInMemoryPostRepo(db *database.InMemoryDB) output.PostPort {
	repo := inmemory.NewInMemoryPostRepository(db)
	return repo
}

func UserSeed(amount uint64) ([]*inmemory_schema.UserSchema, []*domain.User) {
	if amount <= 0 {
		amount = 1
	}
	seeds := make([]*inmemory_schema.UserSchema, amount)
	users := make([]*domain.User, amount)
	for i := 0; i < len(seeds); i++ {
		username := "user" + strconv.Itoa(i)
		newUser := domain.NewUser(username)
		users[i] = newUser
		seeds[i] = &inmemory_schema.UserSchema{
			ID:        newUser.ID,
			Username:  newUser.Username,
			CreatedAt: newUser.CreatedAt,
			UpdatedAt: newUser.UpdatedAt,
		}
	}
	return seeds, users
}

func PostSeed(user *inmemory_schema.UserSchema, amount int) ([]*inmemory_schema.PostSchema, []*domain.Post) {
	seeds := make([]*inmemory_schema.PostSchema, amount)
	posts := make([]*domain.Post, amount)
	for i := 0; i < len(seeds); i++ {
		newPostInput := domain.NewPostInput{
			User:    mappers.NewUserMapper().FromPersistence(user),
			Content: "post_" + strconv.Itoa(i),
		}
		newPost, _ := domain.NewPost(newPostInput)
		posts[i] = newPost
		seeds[i] = &inmemory_schema.PostSchema{
			ID:                     newPost.ID,
			UserID:                 newPost.User.ID,
			Content:                newPost.Content,
			CreatedAt:              newPost.CreatedAt,
			IsQuote:                newPost.IsQuote,
			IsRepost:               newPost.IsRepost,
			OriginalPostID:         newPost.OriginalPostID,
			OriginalPostContent:    newPost.OriginalPostContent,
			OriginalPostUserID:     newPost.OriginalPostUserID,
			OriginalPostScreenName: newPost.OriginalPostScreenName,
		}
	}
	return seeds, posts
}

type TestMocks struct {
	MockDB        *database.InMemoryDB
	MockUserSeed  []*domain.User
	MockPostsSeed []*domain.Post
}

func GetTestMocks() TestMocks {
	mockDB := MakeDb()
	userSeeds, mockUsers := UserSeed(1)
	postSeeds, mockPosts := PostSeed(userSeeds[0], 2)

	mockDB.RegisterSchema(inmemory.USER_SCHEMA_NAME, userSeeds)
	mockDB.RegisterSchema(inmemory.POST_SCHEMA_NAME, postSeeds)

	return TestMocks{
		MockDB:        mockDB,
		MockUserSeed:  mockUsers,
		MockPostsSeed: mockPosts,
	}
}

func InsertUserHelper(db *database.InMemoryDB, newItem *inmemory_schema.UserSchema) {
	existing, ok := db.Schemas[inmemory.USER_SCHEMA_NAME].([]*inmemory_schema.UserSchema)
	if !ok {
		log.Fatal("schema " + inmemory.USER_SCHEMA_NAME + " not found")
		return
	}

	updateSlice := append(existing, newItem)
	db.Schemas[inmemory.USER_SCHEMA_NAME] = updateSlice
}

func InsertPostHelper(db *database.InMemoryDB, newItem *inmemory_schema.PostSchema) {
	existing, ok := db.Schemas[inmemory.POST_SCHEMA_NAME].([]*inmemory_schema.PostSchema)
	if !ok {
		log.Fatal("schema " + inmemory.POST_SCHEMA_NAME + " not found")
		return
	}

	updateSlice := append(existing, newItem)
	db.Schemas[inmemory.POST_SCHEMA_NAME] = updateSlice
}
