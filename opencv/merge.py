#!/usr/bin/env python3

import sys
import cv2 as cv
import numpy as np

try:
  _, infile, outfile, dilate, background = sys.argv
  dilate = int(dilate)
except:
  print("usage: mask.py infile outfile dilate background")
  sys.exit(1)

fg = cv.imread(infile)
bg = cv.imread(background)

# keep only the edge + dilation on the fg picture, set everything else to transparent black
edges = cv.Canny(fg, 10, 150)
kernel = np.ones((dilate, dilate), np.uint8)
eroded = cv.dilate(edges, kernel, iterations=1)
masked = cv.bitwise_and(fg, fg, mask=eroded)
r,g,b = cv.split(masked)
rgba = cv.merge((r,g,b,eroded))

# apply the inverse mask to the bg picture, set everything else to transparent black
mask = cv.bitwise_not(eroded)
bgmasked = cv.bitwise_and(bg, bg, mask=mask)
r,g,b = cv.split(bgmasked)
bgmask = cv.merge((r,g,b,mask))

# add the two together
merged = cv.addWeighted(bgmask, 1, rgba, 1, 0)
cv.imwrite(outfile, merged)
