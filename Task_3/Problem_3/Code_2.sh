# Time Complexity : O(n)

# Sum of the numbers (float, int, -ve)

total=0

re='^-?[0-9]+[.]?[0-9]+$' 
for i in $(cat file2.txt); do
  if [[ $i =~ $re ]];
  then
    total=$(echo $total + $i | bc);  # bc : basic calculator 
  fi
done

echo "sum = ${total}" 
