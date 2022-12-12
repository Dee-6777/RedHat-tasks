// Time Complexity : N*O(1) ; N is the no of times user wants to continue the loop
// Space Complexity : O(1)

package main /* The package main tells the Go compiler that the package should compile as an executable program instead of a shared library.
The main function in the main package is the entry point of the program*/

import (
	"crypto/rand"
	"fmt" // fmt stands for the Format package. This package allows to format basic strings, values, or anything and print
	"log"
	"math/big"
	"time"
)

func genRandNum(min, max int64) int64 {
	// will get maximum
	bg := big.NewInt(max - min)

	// get big.Int between 0 and bg
	// in this case 0 to 101
	n, err := rand.Int(rand.Reader, bg)
	if err != nil {
		panic(err)
	}

	// add n to min to support the passed in range
	return n.Int64() + min
}

func main() {
	start := time.Now()
	var check bool = true

	for check == true {
		randomNumber := genRandNum(1, 100)
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
