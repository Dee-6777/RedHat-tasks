arr=( `cat "IP addresses.txt" `)
printf '%s' "${arr[@]}" | sort -n | uniq -c | sort -n | tail -1