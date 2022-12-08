package main

import (
	"fmt"
)

func IsLetter(s string) bool {
	for _, r := range s {
		if (r < 'a' || r > 'z') && (r < 'A' || r > 'Z') {
			return false
		}
	}
	return true
}

func IsNumber(s string) bool {
	for _, c := range s {
		if c < '0' || c > '9' {
			return false
		}
	}
	return true
}

func main() {
	var check bool = true
	for check == true {
		var name1, name2 string
		var age string
		var nationality string
		fmt.Println("Enter your First Name :")
		fmt.Scanln(&name1)
		for (IsLetter(name1)) == false {
			fmt.Println("Enter again :")
			fmt.Scanln(&name1)
		}
		fmt.Println("Enter your Surname :")
		fmt.Scanln(&name2)
		for (IsLetter(name2)) == false {
			fmt.Println("Enter again :")
			fmt.Scanln(&name2)
		}
		fmt.Println("Enter your age :")
		fmt.Scanln(&age)
		for (IsNumber(age)) == false {
			fmt.Println("Enter again :")
			fmt.Scanln(&age)
		}
		fmt.Println("Enter your Nationality :")
		fmt.Scanln(&nationality)
		for (IsLetter(nationality)) == false {
			fmt.Println("Enter again :")
			fmt.Scanln(&nationality)
		}

		fmt.Printf("Your name is %s %s your age is %v years old and your nationality is %s\n", name1, name2, age, nationality)
		fmt.Println("Want to continue?\nThen type 1 else type 0")
		fmt.Scanln(&check)
	}
}
