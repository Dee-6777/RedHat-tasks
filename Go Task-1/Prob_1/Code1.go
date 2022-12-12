// Space Complexity : O(1) i.e only variables are used

// This program uses infinite but user controlled loop and functions to check characters those are alphabets or digits


package main	/* The package main tells the Go compiler that the package should compile as an executable program instead of a shared library. 
				The main function in the main package is the entry point of the program*/

import "fmt"  	/*fmt stands for the Format package. This package allows to format basic strings, values, or anything and print them or collect user input from the console, 
				or write into a file using a writer or even print customized fancy error messages.*/

func IsLetter(s string) bool {    // Function to check if the string contains any character other than alphabets.
	for _, r := range s {
		if (r < 'a' || r > 'z') && (r < 'A' || r > 'Z') {     
			return false
		}
	}
	return true
}

func IsNumber(s string) bool {      // Function to check if the string contains any character other than numbers.
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
		
		for (IsLetter(name1)) == false {      // Here it sends 'name1' to the above defined IsLetter function to check whether it is alphabet or not
			fmt.Println("Enter again :")      // If the return is false it will ask to enter again until correct input is recieved
			fmt.Scanln(&name1)
		}
		fmt.Println("Enter your Surname :")
		fmt.Scanln(&name2)
		for (IsLetter(name2)) == false {     // Here it sends 'name2' to the above defined IsLetter function to check whether it is alphabet or not
			fmt.Println("Enter again :")	 
			fmt.Scanln(&name2)
		}
		fmt.Println("Enter your age :")      // Here it sends 'age' to the above defined IsNumber function to check whether it is numeric or not
		fmt.Scanln(&age)
		for (IsNumber(age)) == false {
			fmt.Println("Enter again :")
			fmt.Scanln(&age)
		}
		fmt.Println("Enter your Nationality :")  // Here also we check for Nationality
		fmt.Scanln(&nationality)
		for (IsLetter(nationality)) == false {
			fmt.Println("Enter again :")
			fmt.Scanln(&nationality)
		}
		fmt.Println("Your name is", name1, name2, "your age is", age, "years old and your nationality is", nationality) // Printing as per in the format asked in the question
		
		fmt.Println("Want to continue?\nThen type 1 else type 0")
		fmt.Scanln(&check)
	}
}
