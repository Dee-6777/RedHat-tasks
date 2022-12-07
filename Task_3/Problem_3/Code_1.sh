# Time Complexity : O(n)


# Finds the sum of the integers
sum=0
for i in $(cat file1.txt);
do
    sum=$((sum+$i))
done

echo "sum = $sum" 
# limits only to +ve integers
