// Time Complexity : N*O(1) ; N is the no of times user wants to continue the loop
// Space Complexity : O(1)

package main /* The package main tells the Go compiler that the package should compile as an executable program instead of a shared library.
The main function in the main package is the entry point of the program*/

import (
	"fmt" // fmt stands for the Format package. This package allows to format basic strings, values, or anything and print
	"log" // writes to standard error and prints the date and time of each logged message
	"math/rand" // provides inbuilt support for generating random numbers of the specified type
	"time"      // provides current time and in rand.Seed it is initialised with the current time expressed in microseconds
)

func main() {
	start := time.Now()
	// Seeding - Go provides a method, Seed(see int64), that allows you to initialize this default sequence
	// Implementation is slow to make it faster rand.Seed(time.Now().UnixNano()) is added.
	// Seed is the current time, converted to int64 by UnixNano.
	// Gives constantly changing numbers
	rand.Seed(time.Now().UnixNano()) // Seed uses the provided seed value to initialize the default Source to a deterministic state, If Seed is not called, the generator behaves as if seeded by Seed(1).
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

		if check == false {
			fmt.Println("Thanks for using")
		}
	}
	log.Printf("main, execution time %s\n", time.Since(start))
}
