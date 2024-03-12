package main

import "fmt"

func arrays() {
	// 1
	hobbies := [3]string{"chess", "reading", "coding"}
	fmt.Println(hobbies)
	// 2
	firstHobby := hobbies[0]
	fmt.Println(firstHobby)
	otherHobbies := hobbies[1:]
	fmt.Println(otherHobbies)
	// 3
	firstAndSecondHobbies := hobbies[:2]
	fmt.Println(firstAndSecondHobbies)
	// 4
	mainHobbies := hobbies[1:cap(firstAndSecondHobbies)]
	fmt.Println(mainHobbies)
}
