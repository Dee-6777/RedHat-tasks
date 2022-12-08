# Problem 1:
```
Write a loop to ask the name , age and nationality from a user and print "Your name is ___  your age is ___ years old and your nationality is __"
This should be infinite loop
```
## Solution : 

* How to run : Type `go run Code1.go` in your bash terminal
* Space Complexity : O(1) [constant]
* Explanation :
This code has been made using functions IsNumber() and IsLetter() and using an infinite user controlled loop. 
    * `func IsLetter(s string) bool{}` : Checks for name and nationality i.e if the string contains any character other than alphabets.
    * `func IsNumber(s string) bool {}` : Checks for age i.e if the string contains any character other than numbers.

## Output :
![Screenshot from 2022-12-09 02-07-06](https://user-images.githubusercontent.com/73513838/206564040-b99d9226-7b2b-46a0-b9e9-edb173ebf067.png)
