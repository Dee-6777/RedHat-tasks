# with using associative arrays
# Time Complexity: O(n)
# Space Complexity: O(n)
#!/bin/bash
arr=( `cat "IP addresses.txt" `)
declare -A freq
max_times=0
for i in "${arr[@]}"; 
do
  ((freq[$i]++))
  h=${freq[$i]}
  if [[ $h > $max_times ]]; 
  then
    max=$i
    max_times=$h
  fi
done

printf '%s\n' "$max"
printf '%d\n' "$max_times"
