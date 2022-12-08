package main

import (
	"fmt"
)

func main() {
	Menu := []string{"Pizza", "Chicken Nuggets", "French Fries"}
	Menu = append(Menu, "Hamburger", "Salad")
	for _, k := range Menu {
		fmt.Printf("Food : %s\n", k)
	}
}
