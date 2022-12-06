# Problem 2:
`Write a script to remove duplicated lines from a file.`

## Solution 1: (Using sort and uniq)
* Time Complexity : O(nlob(n)) (loglinear complexity)

* How to run : Type `sh Code1.sh` in your bash terminal

* Explanation :
    * `sort` : Will sort the contents of the file alphabetically (using merge sort)
    * `uniq` : Separates out the uniq characters
    * `cat -n > output1.txt` : cat appends the uniq lines along with the line numbers inside the output file.

## Solution 2: (Using awk)
* Time Complexity : O(n) (linear complexity)
* Explanation :
    * `awk '!_[$0]++' dup.txt` : awk reads each line and !_[$0]++ first checks the first occurance of each line and then increements the count. 
    * `cat -n > output2.txt` : cat appends the uniq lines along with the line numbers inside the output file.
