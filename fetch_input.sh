#!/bin/bash
UA="Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.3"
YEAR=$(date +'%Y')

# Check if a day parameter is provided, otherwise use today's date
if [ -n "$1" ]; then
    DAY="$1"
else
    DAY=$(date +'%-d')
fi 
# Fetch input using curl
curl -s --header "Cookie: $SESSION" --header "User-Agent: $UA" "https://adventofcode.com/$YEAR/day/$DAY/input"

