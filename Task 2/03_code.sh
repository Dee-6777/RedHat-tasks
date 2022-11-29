# with using associative arrays

#!/bin/bash
arr=( `cat "IP addresses.txt" `)
maxcount=0
for i in "${arr[@]}";
do
    count=0
    for j in "${arr[@]}";
    do
        if [[ $i = $j ]];
        then
            count=$((count+1))
        fi
    done

    if [[ $count > $maxcount ]];
    then
        maxcount=$count
        element_max_freq=$i
    fi
done

printf '%s\n' "$element_max_freq"
printf '%d\n' "$maxcount"