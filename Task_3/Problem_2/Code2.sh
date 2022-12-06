# Time Complexity : O(n)

awk '!_[$0]++' dup.txt | cat -n > output2.txt

# Using Hashmaps