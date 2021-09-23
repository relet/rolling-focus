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
edges = cv.Canny(src, 10, 200)

kernel = np.ones((dilate, dilate), np.uint8)
eroded = cv.dilate(edges, kernel, iterations=1)
masked = cv.bitwise_and(src, src, mask=eroded)

r,g,b = cv.split(masked)
mask = cv.bitwise_not(eroded)

rgba = cv.merge((r,g,b,eroded))

cv.imwrite(outfile, rgba)
