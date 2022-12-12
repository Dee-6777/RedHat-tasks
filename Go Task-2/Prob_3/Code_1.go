// Time Complexity : N*O(1) ; N is the no of times user wants to continue the loop
// Space Complexity : O(1)

package main

import (
	"fmt"       //fmt stands for the Format package. This package allows to format basic strings, values, or anything and print
	"math/rand" //  provides inbuilt support for generating random numbers of the specified type
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	var check bool = true
	for check == true {

		randomNumber := rand.Intn(99) + 1

		if randomNumber > 50 {
			if randomNumber%2 == 0 {
				fmt.Println("It's closer to 100 and it's even!")
			}
			fmt.Println("It's closer to 100")
		} else if randomNumber < 50 {
			fmt.Println("It's closer to 0")
		} else {
			fmt.Println("It's 50!")
		}

		fmt.Println("The generated number is : ", randomNumber)

		fmt.Println("Want to continue?\nThen type 1 else type 0")
		fmt.Scanln(&check)
	}
}
