package model

type User struct {
	FirstName string
	LastName  string
	Email     string
}

var u *User

func NewUser(firstName, lastName, email string) {
	u = &User{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
	}
}

func GetUser() *User {
	return u
}
