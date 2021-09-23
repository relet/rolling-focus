#!/bin/bash
set -x

adb shell input tap 350 1250
sleep 1
FILE=$(adb shell 'ls $EXTERNAL_STORAGE/DCIM/*.jpg | head -n1')
adb pull $FILE
adb shell "rm $FILE"

FILE=$(ls *.jpg|grep -v collection|head -n1)
mv $FILE collection/
