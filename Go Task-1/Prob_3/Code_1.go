// Time Complexity : O(N) ; where N is the length of the array 
// Space Complexity : O(N)

// This program uses a single for loop to print the elements 


package main	/* The package main tells the Go compiler that the package should compile as an executable program instead of a shared library. 
				The main function in the main package is the entry point of the program */

import "fmt"  	/*fmt stands for the Format package. This package allows to format basic strings, values, or anything and print them or collect user input from the console, 
				or write into a file using a writer or even print customized fancy error messages. */

func main() {
	students := [5]string{"Aditya", "Deepti", "Sai", "Arpita", "Harsh"}  // Declared an array of type string an initialised it with 5 Names of type string

	for i, k := range students {
		fmt.Printf("The name of the students is %s and the position of the student in the array is %v\n", k, i)  /* Here i represents the position of the 
	}																										element inside the array and k is element */ 			
}
