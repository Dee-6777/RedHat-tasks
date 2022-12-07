# Time Complexity : O(n)

# Sum of the numbers (float, int, -ve)

total=0
rexprr='^-?[0-9]+[.]?[0-9]+$' #?
for i in $(cat file2.txt); do 
  if [[ $i =~ $rexprr ]];   # this is a tilde symbol which is used for the expressions ?
  then
    total=$(echo $total + $i | bc);  # bc : basic calculator 
  fi
done

echo "sum = ${total}" 

# additional expression added for covering corner casses such as alpha-numeric characters, special characters etc..
# Here, ^ means it matches the empty string at the beginning of a line
# [0-9] : represents the range 
# + : The preceding item will be matched one or more times [1,]
# ? : The preceding item is optional and will be matched, at most, once [0,1]
# $ : Matches the empty string at the end of a line
