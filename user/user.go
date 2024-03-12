package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

type Admin struct {
	email    string
	password string
	User
}

func NewAdmin(email, password, firstName, lastName, birthDate string) (*Admin, error) {
	user, err := New(firstName, lastName, birthDate)
	if err != nil {
		return nil, err
	}

	return &Admin{
		email:    email,
		password: password,
		User:     *user,
	}, nil
}

// method of user (u user) is receiver argument
func (u User) OutputUserDetails() {
	fmt.Println("First Name:", u.firstName)
	fmt.Println("Last Name:", u.lastName)
	fmt.Println("Birth Date:", u.birthDate)
	fmt.Println("Created At:", u.createdAt)
}

func (u *User) ClearUserDetails() {
	u.firstName = ""
	u.lastName = ""
}

func New(firstName, lastName, birthDate string) (*User, error) {
	if firstName == "" || lastName == "" || birthDate == "" {
		return nil, errors.New("Invalid user data")
	}

	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthDate,
		createdAt: time.Now(),
	}, nil
}
