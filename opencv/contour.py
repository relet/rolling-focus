#!/usr/bin/env python3

import sys
import cv2 as cv
import numpy as np

try:
  _, infile, outfile, dilate = sys.argv
  dilate = int(dilate)
except:
  print("usage: mask.py infile outfile dilate")
  sys.exit(1)

src = cv.imread(infile)
kernel = np.ones((dilate, dilate), np.uint8)

# detect darkness
#gray = cv.cvtColor(src, cv.COLOR_BGR2GRAY)
#blurred = cv.GaussianBlur(gray, (159,159), 25)
#blurred = cv.dilate(gray, kernel, iterations=1)
#darkness = cv.adaptiveThreshold(blurred, 255, cv.ADAPTIVE_THRESH_GAUSSIAN_C, cv.THRESH_BINARY, 55, 1)
#blur2 = cv.GaussianBlur(darkness, (159,159), 25)
#darkness = cv.adaptiveThreshold(blur2, 255, cv.ADAPTIVE_THRESH_GAUSSIAN_C, cv.THRESH_BINARY, 55, 1)
#darkness = cv.resize(darkness, (int(darkness.shape[0]/4), int(darkness.shape[1]/4)))
#cv.imshow('test', darkness)
#cv.waitKey(0)

# detect edges
edges = cv.Canny(src, 10, 60)
eroded = cv.dilate(edges, kernel, iterations=1)
masked = cv.bitwise_and(src, src, mask=eroded)

r,g,b = cv.split(masked)
mask = cv.bitwise_not(eroded)

rgba = cv.merge((r,g,b,eroded))

preview = cv.resize(rgba, (int(rgba.shape[0]/4), int(rgba.shape[1]/4)))
cv.imshow('test', preview)
cv.waitKey(0)

#cv.imwrite(outfile, rgba)
