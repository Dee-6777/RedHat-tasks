// Time Complexity : N*O(1) ; N is the no of times user wants to continue the loop
// Space Complexity : O(1)

package main /* The package main tells the Go compiler that the package should compile as an executable program instead of a shared library.
The main function in the main package is the entry point of the program*/

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
			fmt.Println("It's closer to 100")

		} else if randomNumber < 50 {
			fmt.Println("It's closer to 0")

		}

		fmt.Println("The generated random number: ", randomNumber)

		fmt.Println("Want to continue?\nThen type 1 else type 0")

		fmt.Scanln(&check)
	}
}
