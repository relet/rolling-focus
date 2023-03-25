#!/bin/bash
set -x

COUNT=$(cat counter)

adb shell input tap 350 1250

##1/100
#sleep 3.0

##1/60
sleep 5.0

##1/30
#sleep 10.0

while true; do
  FILE=$(adb shell 'ls $EXTERNAL_STORAGE/DCIM/OpenCamera/*.jpg | head -n1')
  if [ -z "$FILE" ]; then break; fi
  adb pull $FILE
  adb shell "rm $FILE"
done

echo $(( $COUNT + 3 )) > counter

