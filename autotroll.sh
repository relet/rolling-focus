#!/bin/bash
set -x
stty -echoctl # hide ^C

DIRECTION="Z+"

STEPS=80
if [ -z "$2" ]; then
  DIRECTION="Z-"
fi
if [ "$2" == "0" ]; then
  DIRECTION="Z-"
fi
if [ -z "$1" ]; then
  STEPS=$1
fi

echo "MOVING IN DIRECTION $DIRECTION"

# some settings
ERODE=5
SLEEP=1
VIEWER=eog
STEP=0

function background() {
    ./enfuse/auto.sh collection && {
      TARGET=target/$(date +%s.jpg)
      mv collection.jpg $TARGET
      rm *.jpg
      echo "Interrupted after step $STEP"
      exit 0
    } || {
      echo "CRASHED"
      exit 1
    }

}

function cleanup () {
    background
}

function fwd () {
    echo $DIRECTION | minicom -b 9600 -D /dev/ttyUSB0
}

function bak() {
    echo "Z-" | minicom -b 9600 -D /dev/ttyUSB0
}

trap cleanup SIGINT

mkdir -p collection
mkdir -p target
rm collection/*.jpg
rm nohup.out
adb shell rm /sdcard/DCIM/*.jpg
adb shell rm /sdcard/DCIM/OpenCamera/*.jpg
pkill $VIEWER

for i in $(seq 1 75); do
  STEP=$i
  #./android/getpic.sh
  ./android/ochdr.sh
  FILE=$(ls *.jpg|grep -v merged.png|head -n1)
  #if [ -f merged.png ]; then
  #  ./opencv/merge.py $FILE merged.png $ERODE merged.png
  mv $FILE collection/
  #else
  #  convert $FILE -rotate 90 merged.png
  #  mv $FILE collection/
  #  $VIEWER merged.png&
  #fi
  fwd
  #sleep 2 #wait for the stepper to rotate
done
play /home/relet/Downloads/Emoji-trombone.wav

background

