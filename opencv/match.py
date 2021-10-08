#!/usr/bin/env python3

import cv2 as cv
import numpy as np
import sys

img1 = cv.imread('target/1632825572.jpg')

orb = cv.ORB_create()
kp = orb.detect(img1, None)
kp, des = orb.compute(img1, kp)

img1kp = cv.drawKeypoints(img1, kp, None, color=(255,0,0), flags=0)
cv.imshow('Test', img1kp)
cv.waitKey(0)

#img2 = cv.imread('target/1632825897.jpg')

