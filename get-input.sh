n=$1

curl -H "Cookie: session=$(cat session_cookie)" https://adventofcode.com/2018/day/$n/input > ./$n/input.txt
