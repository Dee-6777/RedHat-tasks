# Problem 1:
`Write a script to separate day, month, year, hour, minute and second values of the current system date and time.`

## Solution 1: (Using echo and modifiers)
* Time Complexity : O(1) (Constant Complexity)

* How to run : Type `sh Code1.sh` in your bash terminal

* Explanation :
This code been simply made using modifiers along with echo to get the desired output
    * Used Modifiers : 
        *  %A - Day
        *  %m - Month
        *  %Y - Year 
        *  %M - minute (00..59)
        *  %H - hour (00..23)
        *  %S - second (00..60)
    * Use of "-e" : Using the -e option allows us to use escape characters. These special characters make it easy to customize the output of the echo command.
