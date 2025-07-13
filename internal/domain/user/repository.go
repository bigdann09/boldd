package user

type IUserRepository interface {
	Create(user *User) error
	Find(id int) (interface{}, error)
}
