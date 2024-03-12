package main

import "fmt"

func pointers() {
	age := 32

	fmt.Println("Age:", age)

	editAgeToAdultYears(&age)
	fmt.Println("Adult years:", age)

}

func editAgeToAdultYears(age *int) {
	*age = *age - 18
}
