package main

import (
	"fmt"
	"image/color"
	"os"
	"os/exec"

	"gocv.io/x/gocv"
)

const (
	FOCUS_THRESHOLD = 1000 // Threshold for focus detection
)

var (
	colorRed    = color.RGBA{255, 0, 0, 0}
	colorGreen  = color.RGBA{0, 255, 0, 0}
	colorYellow = color.RGBA{255, 255, 0, 0}
)

func IsImageInFocus(img gocv.Mat) bool {
	if img.Empty() {
		return false
	}

	// Simple focus detection logic: check if canny edges are detected
	edges := gocv.NewMat()
	defer edges.Close()
	gocv.Canny(img, &edges, 100, 200)
	edgeCount := gocv.CountNonZero(edges)

	return edgeCount > FOCUS_THRESHOLD
}

func Fuse(stack *ImageStack) error {
	// create temporary folder
	if stack == nil || len(stack.Images) == 0 {
		return fmt.Errorf("no images to fuse")
	}
	os.Mkdir("tmp_fuse", 0755)
	os.Mkdir("collection", 0755)

	// write all Images as jpeg to disk
	for _, imgPos := range stack.Images {
		filename := fmt.Sprintf("tmp_fuse/image_x%d_y%d_z%d.jpg", stack.x, stack.y, imgPos.z)
		gocv.IMWrite(filename, imgPos.img)
	}

	// fuse images using enfuse
	_, err := exec.Command(`enfuse -o collection/fused_x` + fmt.Sprint(stack.x) + `_y` + fmt.Sprint(stack.y) + `.jpg \
 		--exposure-weight=0 \
 		--saturation-weight=0 \
 		--contrast-weight=1 \
 		--hard-mask \
 		--contrast-window-size=9 \
 		tmp_fuse/*.jpg`).Output()
	if err != nil {
		fmt.Println("Error running enfuse:", err)
		return err
	}

	return nil
}
