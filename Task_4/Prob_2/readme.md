# Problem 2:
```
Define the following array "Menu"}
Append to it the following item: "hamburger"
Append to it the following item: "salad"
Iterate over the list and print for each item Food: <Food name>. Make sure to replace <Food name> with item from the array
```
## Solution : 

* How to run : Type `go run Code_1.go` in your bash terminal
* Time Complexity : O(N) [Linear time Complexity]
* Explanation :
* In this code I've first defined an array named "Menu" and then appended the two items using the inbuild append function. And after that printed as per the given format.
    * `Menu := []string{"Pizza", "Chicken Nuggets", "French Fries"}` : Declared an array called Menu and added 3 elements initially.
    * `Menu = append(Menu, "Hamburger", "Salad")` : appended the required elements
    * `for _, k := range Menu {}` : looping over the elements inside Menu array.
