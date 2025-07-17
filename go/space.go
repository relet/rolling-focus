package main

import (
	"gocv.io/x/gocv"
)

type ImagePos struct {
	z   int
	img gocv.Mat
}

type ImageStack struct {
	x, y       int
	zmin, zmax int
	Images     []ImagePos
}

type ScanPosition struct {
	x, y, z int
}

const (
	StateInit = iota
	StateScanZUp
	StateRolling
	StateMove
)

const (
	XY_STEP = 10
)

type ImageSpace struct {
	cx, cy, cz             int
	state                  int
	xmin, ymin, xmax, ymax int
	zmin, zmax             int
	Stacks                 []ImageStack
	currentStack           *ImageStack

	scanPositions []ScanPosition
	scanned       []ScanPosition
}

func NewImageSpace() *ImageSpace {
	return &ImageSpace{
		cx:            0,
		cy:            0,
		cz:            0,
		state:         StateScanZUp,
		xmin:          0,
		ymin:          0,
		xmax:          0,
		ymax:          0,
		zmin:          0,
		zmax:          0,
		Stacks:        make([]ImageStack, 0),
		currentStack:  nil,
		scanPositions: make([]ScanPosition, 0),
		scanned:       make([]ScanPosition, 0),
	}
}

func (space *ImageSpace) AddScanPositions(x, y, z int) {
	// Add scan positions around the current position
	for dx := -XY_STEP; dx <= XY_STEP; dx += XY_STEP {
		for dy := -XY_STEP; dy <= XY_STEP; dy += XY_STEP {
			if dx != 0 || dy != 0 { // Skip the current position
				// Ensure we don't add positions that are already scanned
				for _, pos := range space.scanned {
					if pos.x == x+dx && pos.y == y+dy {
						continue
					}
				}
				// Add the new scan position
				space.scanPositions = append(space.scanPositions, ScanPosition{x: x + dx, y: y + dy, z: z})
			}
		}
	}
}

func (space *ImageSpace) GetClosestScanPosition(x, y int) *ScanPosition {
	if len(space.scanPositions) == 0 {
		return nil // No scan positions available
	}
	min_dist := 1<<31 - 1
	var pos ScanPosition
	for _, p := range space.scanPositions {
		dist := (p.x-x)*(p.x-x) + (p.y-y)*(p.y-y)
		if dist < min_dist {
			min_dist = dist
			pos = p
		}
	}
	space.scanned = append(space.scanned, pos)
	for i, p := range space.scanPositions {
		if p.x == pos.x && p.y == pos.y {
			space.scanPositions = append(space.scanPositions[:i], space.scanPositions[i+1:]...)
			break
		}
	}

	return &pos
}
