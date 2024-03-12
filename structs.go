package main

import (
	"fmt"

	"example.com/learngo/user"
)

func structs() {

	userFirstName := getUserData("Enter your first name: ")
	userLastName := getUserData("Enter your last name: ")
	userBirthDate := getUserData("Enter your birth date: ")

	var appUser *user.User

	// appUser = user{
	// 	firstName: userFirstName,
	// 	lastName:  userLastName,
	// 	birthDate: userBirthDate,
	// 	createdAt: time.Now(),
	// }

	admin, err1 := user.NewAdmin("jl@email.com", "xxx", "JL", "D", "12/12/12")

	if err1 != nil {
		fmt.Println("Error creating user:", err1)
		return
	}

	appUser, err2 := user.New(userFirstName, userLastName, userBirthDate)

	if err2 != nil {
		fmt.Println("Error creating user:", err1)
		return
	}

	admin.OutputUserDetails()
	admin.ClearUserDetails()
	admin.OutputUserDetails()
	appUser.OutputUserDetails()
	appUser.ClearUserDetails()
	appUser.OutputUserDetails()

}

func getUserData(promptText string) string {
	var userInput string
	fmt.Print(promptText)
	fmt.Scanln(&userInput)
	return userInput
}
