package user

type Repository interface {
	Create(user *User) error
	FindByEmail(email string) (*User, error)
}
