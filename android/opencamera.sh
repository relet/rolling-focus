# settings: exposure
adb shell input tap 350 50
## iso 60
adb shell input tap 234 900
## iso 150 --
# adb shell input tap 333 900
## iso 300 --
#adb shell input tap 390 900
## exposure 1/100 --
# adb shell input tap 400 950
## exposure 1/60 --
#adb shell input tap 450 950
## exposure 1/20 - takes crazy long
 adb shell input tap 520 950
# close settings: exposure
adb shell input tap 350 50

# settings: photo mode
adb shell input tap 550 50
# set hdr
adb shell input tap 400 400

# same menu again
adb shell input tap 550 50
# focus manual
adb shell input tap 350 250

# focus distance: 0.1m
adb shell input tap 50 1065 
# zoom: roughly 3x
adb shell input tap 10 1224
