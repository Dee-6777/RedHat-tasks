# %A   Day
# %B   Month
# %Y   Year 
# %M   minute (00..59)
# %H   hour (00..23)
# %S   second (00..60)

# Time Complexity : O(1)
# Space Complexity : O(1)

echo -e "Current Day : $(date +%A)\nCurrent Month : $(date +%B)\nCurrent Year : $(date +%Y)\nHours : $(date +%H)\nMinutes : $(date +%M)\nSeconds : $(date +%S)" 

# Using echo command and modifiers 
# -e : By using echo -e we're making echo to enable interpret backslash escapes. 
