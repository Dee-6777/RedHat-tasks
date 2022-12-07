# Problem 2:
`Write a script to remove duplicated lines from a file.`

## Solution 1: (Using sort and uniq)
* Time Complexity : O(nlog(n)) (loglinear complexity)

* How to run : Type `sh Code1.sh` in your bash terminal

* Explanation :
    * `sort` : Will sort the contents of the file alphabetically (using merge sort)
    * `uniq` : Separates out the uniq characters

## Solution 2: (Using awk)
* Time Complexity : O(n) (linear complexity)

* How to run : Type `sh Code2.sh` in your bash terminal

* Explanation :
    * `awk '!_[$0]++' dup.txt` : `awk !_[$0]++` maintains a hashmap for all the visited lines and checks if it is not visited.
