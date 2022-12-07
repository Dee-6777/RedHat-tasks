# Time Complexity : O(n)

# Sum of the numbers (float, int, -ve)

total=0
rexprr='^-?[0-9]+[.]?[0-9]+$' 
for i in $(cat file2.txt); do 
  if [[ $i =~ $rexprr ]];   
  then
    total=$(echo $total + $i | bc);  # bc : basic calculator 
  fi
done

echo "sum = ${total}" 

#Expln:
# Finds and discards the strings which are found to be unsuitable for addition
# Time Complexity will be O(n) where n represents the size of the input or the length of the array.
# Addition occurs for the remaining suitable strings.




# Few additional terms :
# additional expression added for covering corner casses such as alpha-numeric characters, special characters etc..
# Here, ^ : matches the empty string at the beginning of a line
# - : represents the range 
# + : The preceding item will be matched one or more times [1,]
# ? : The preceding item is optional and will be matched, at most, once [0,1]
# $ : Matches the empty string at the end of a line
