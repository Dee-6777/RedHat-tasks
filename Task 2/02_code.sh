arr=( `cat "IP addresses.txt" `)
printf '%s\n' "${arr[@]}" | sort -n | uniq -c | sort -n | tail -1