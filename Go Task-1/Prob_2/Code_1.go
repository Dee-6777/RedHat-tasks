// Time Complexity : O(N) ; where N is the length of the array 
// Space Complexity : O(N)

// This program uses inbuilt append function and a single for loop to print the elements 


package main	/*The package main tells the Go compiler that the package should compile as an executable program instead of a shared library. 
				The main function in the main package is the entry point of the program*/

import "fmt"  	/*fmt stands for the Format package. This package allows to format basic strings, values, or anything and print them or collect user input from the console, 
				or write into a file using a writer or even print customized fancy error messages.*/

func main() {
	Menu := []string{"Pizza", "Chicken Nuggets", "French Fries"}  // Declared an array of type string an initialised it with 3 items
	Menu = append(Menu, "Hamburger", "Salad")     // Appendednthe two items asked i.e "Hamburger" and "Salad"
	for _, k := range Menu {
		fmt.Printf("Food : %s\n", k)
	}
}
