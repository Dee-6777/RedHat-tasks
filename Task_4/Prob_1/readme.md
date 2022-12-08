# Problem 1:
```
Write a loop to ask the name , age and nationality from a user and print "Your name is ___  your age is ___ years old and your nationality is __"
This should be infinite loop
```
## Solution : 

* How to run : Type `go run Code1.go` in your bash terminal

* Explanation :
This code uses date modifiers along with echo to get the desired output
    * Used Modifiers : 
        *  %A - Day
        *  %m - Month
        *  %Y - Year 
        *  %M - minute (00..59)
        *  %H - hour (00..23)
        *  %S - second (00..60)
        
    * Use of "-e" : Using the -e option allows us to use escape characters. These special characters make it easy to customize the output of the echo command.

