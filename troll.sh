#!/bin/bash
set -x
stty -echoctl # hide ^C

# some settings
ERODE=5
SLEEP=1
VIEWER=eog

function background() {
    ./enfuse/auto.sh collection && {
      TARGET=target/$(date +%s.jpg)
      mv collection.jpg $TARGET
      rm *.jpg
      exit 0
    } || {
      echo "CRASHED"
      exit 1
    }
}

function cleanup () {
    background
}

trap cleanup SIGINT

mkdir -p collection
mkdir -p target
rm collection/*.jpg
rm nohup.out
adb shell rm /sdcard/DCIM/*.jpg
pkill $VIEWER

while true; do
  ./android/getpic.sh
  FILE=$(ls *.jpg|grep -v merged.png|head -n1)
  #if [ -f merged.png ]; then
  #  ./opencv/merge.py $FILE merged.png $ERODE merged.png
  mv $FILE collection/
  #else
  #  convert $FILE -rotate 90 merged.png
  #  mv $FILE collection/
  #  $VIEWER merged.png&
  #fi
done

