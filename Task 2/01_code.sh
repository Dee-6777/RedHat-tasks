arr=( `cat "IP addresses.txt" `)
declare -A hash
max_times=0
for i in "${arr[@]}"; do
  ((hash[$i]++))
  h=${hash[$i]}
  if [[ $h > $max_times ]]; then
    max=$i
    max_times=$h
  fi
done

printf '%s\n' "$max"
printf '%d\n' "$max_times"