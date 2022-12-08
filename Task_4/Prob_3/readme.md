# Problem 3:
```
Define an array of 5 items
Iterate over it and print for each item the following: This is <ITEM> and its index in the array is <INDEX>
```
## Solution : (Using echo and modifiers)
* Time Complexity : O(1) (Constant Complexity)

* How to run : Type `sh Code1.sh` in your bash terminal

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
