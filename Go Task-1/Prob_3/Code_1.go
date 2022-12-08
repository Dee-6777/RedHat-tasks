package main

import "fmt"

func main() {
	students := [5]string{"Aditya", "Deepti", "Sai", "Arpita", "Harsh"}
	for i, k := range students {
		fmt.Printf("The name of the students is %s and the position of the student in the array is %v\n", k, i)
	}
}
