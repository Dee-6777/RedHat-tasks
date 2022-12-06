# Time Complexity : O(n)

# Generates Random numbers
for((i=0;i<10000;i++));
do
    echo $RANDOM >> file1.txt     #redirects the output of the command or data to end of file file1.txt
done

# Finds the sum of the integers
sum=0
for i in $(cat file1.txt);
do
    sum=$((sum+$i))
done

echo "sum = $sum" > output1.txt
