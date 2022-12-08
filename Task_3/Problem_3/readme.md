# Problem 3:
`Write a script to find the sum of all numbers in a file in Linux.`

## Solution 1: (Using $RANDOM and for loop)
* Time Complexity : O(n) (Linear Complexity)

* How to run : Type `sh Code_1.sh` in your bash terminal

* Explanation :
    * `for ((i=0; i<1000000; i++)) ; do echo $RANDOM >> file1.txt; done` : Generates 10000 random numbers and append each element inside file1.txt.
    
    * `for i in $(cat file1.txt);do sum=$((sum+$i)) done` : For loop to generate sum of all the elements in the fle file1.txt.
    
    
    
    
## Solution 2: (Using for loop, basic calculator(bc) and regular expression )
* Time Complexity : O(n) (linear complexity)

* How to run : Type `sh Code_2.sh` in your bash terminal

* Explanation :
    * `total=$(echo $total + $i | bc);` : Computes sum of all kinds of numbers such as integers and floating point values. The use of bc helps to calculate floating point values as it provides the functionality of a scientific calculator within a Bash script.
    
    * `rexprr='^-?[0-9]+([.][0-9]?)+$'` : This expression will help us discard each string which contains alpha-numeric characters, special characters etc..
