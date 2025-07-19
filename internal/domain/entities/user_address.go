package entities

type UserAddress struct {
	UserID  uint
	State   string
	City    string
	Address string
	ZipCode string
}

func NewUserAddress(userID uint, state, city, zipcode, address string) *UserAddress {
	return &UserAddress{
		UserID:  userID,
		State:   state,
		City:    city,
		Address: address,
		ZipCode: zipcode,
	}
}
