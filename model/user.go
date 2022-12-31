package model

type User struct {
	ID string
	PW string
}

func NewUser(i string, p string) User{
	return User{
		ID: i,
		PW: p,
	}
}