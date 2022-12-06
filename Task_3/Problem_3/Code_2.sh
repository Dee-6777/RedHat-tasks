# Time Complexity : O(n)

# Sum of the numbers (float, int, -ve)

total=0

for i in $(cat file2.txt); do
  total=$(echo $total + $i | bc);  # bc : basic calculator 
done

echo "sum = ${total}" > output2.txt

