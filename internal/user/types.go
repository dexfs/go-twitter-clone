package user

type UserRepository interface {
	ByUsername(username string) (*User, error)
	FindByID(id string) (*User, error)
}
