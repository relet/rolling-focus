echo "IF THATS NOT WHAT YOU WANT ABORT NOW"
read
rm -r collection/* 2>/dev/null
rm target/*        2>/dev/null
echo "LAST PROJECT HAD $(cat counter) PHOTOS"
echo 0 > counter
