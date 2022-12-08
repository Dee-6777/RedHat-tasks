# Problem 3:
```
Define an array of 5 items
Iterate over it and print for each item the following: This is <ITEM> and its index in the array is <INDEX>
```
## Solution :
* How to run : Type `go run Code_1.go` in your bash terminal
* Time Complexity : O(N) [Linear time Complexity]
* Space Complexity : O(N) [Linear time Complexity]
* Explanation :
* In this code I've first defined an array named "Students" having 5 elements and then printed as per the given format.
    * `students := [5]string{"Aditya", "Deepti", "Sai", "Arpita", "Harsh"}` : Declared an array called Students of type string and initialised with 5 elements.
    * `for i, k := range Menu {}` : looping over the elements inside Menu array.
    * `fmt.Printf("The name of the students is %s and the position of the student in the array is %v\n", k, i)`: Printed the name of the students first and after that their position in the array.
