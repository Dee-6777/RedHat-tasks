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
    * `Menu := []string{"Pizza", "Chicken Nuggets", "French Fries"}` : Declared a slice named Menu and added 3 elements initially.
    * `Menu = append(Menu, "Hamburger", "Salad")` : appended the required elements
    * `for i := 0; i < len(Menu); i++ {}` : looping over the elements inside Menu array.

## Output
![Screenshot from 2022-12-09 02-18-48](https://user-images.githubusercontent.com/73513838/206564818-e00559f9-64a2-4c70-bb3f-2a6b037fdac5.png)
