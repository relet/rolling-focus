package main

import (
	"flag"
	"fmt"
	"image"

	"go.bug.st/serial"
	"gocv.io/x/gocv"
)

func main() {
	deviceid := flag.Int("i", 0, "Device ID of the camera")
	serialPort := flag.String("p", "/dev/ttyACM0", "Serial port for motor control")
	flag.Parse()

	fmt.Printf("Opening camera with device ID: %d\n", *deviceid)

	webcam, _ := gocv.OpenVideoCapture(*deviceid)
	window := gocv.NewWindow("Rolling Focus")
	img := gocv.NewMat()
	defer webcam.Close()
	defer window.Close()
	defer img.Close()

	port, err := serial.Open(*serialPort, serialMode)
	if err != nil {
		fmt.Println("Error opening serial port:", err)
		//return
	} else {
		defer port.Close()
	}

	space := NewImageSpace()
	init_fail := 0
	var target *ScanPosition
	statusColor := colorGreen

	for {
		webcam.Read(&img)

		switch space.state {

		case StateInit:
			if IsImageInFocus(img) {
				statusColor = colorGreen
				space.state = StateScanZUp
			} else {
				statusColor = colorRed
				init_fail += 1
				if init_fail > 10 {
					fmt.Println("No object in focus, exiting...")
					return
				}
				motorMove(port, AxisZ, DirUp)
			}

		case StateScanZUp:
			if IsImageInFocus(img) {
				statusColor = colorGreen
				motorMove(port, AxisZ, DirUp)

				if space.currentStack == nil {
					space.currentStack = &ImageStack{
						x:    space.cx,
						y:    space.cy,
						zmin: space.cz,
						zmax: space.cz,
					}
				}

				space.cz -= 1
			} else {
				statusColor = colorRed
				space.state = StateRolling

				if space.zmin > space.cz {
					space.zmin = space.cz
				}
				if space.cz == space.zmin {
					// only if we have reached the minimum known z position
					// init a stack and add scan positions around the current position
				} else {
					statusColor = colorYellow
					// keep moving
					motorMove(port, AxisZ, DirUp)
					space.cz -= 1
				}
			}

		case StateRolling:
			if IsImageInFocus(img) {
				statusColor = colorGreen
				space.currentStack.Images = append(space.currentStack.Images, ImagePos{z: space.cz, img: img.Clone()})
				if space.currentStack.zmax < space.cz {
					space.currentStack.zmax = space.cz
				}
				if space.currentStack.zmin > space.cz {
					space.currentStack.zmin = space.cz
				}
				motorMove(port, AxisZ, DirDown)
				space.cz += 1
			} else {
				statusColor = colorRed
				space.Stacks = append(space.Stacks, *space.currentStack)

				if space.zmax < space.cz {
					space.zmax = space.cz
				}

				if space.cz == space.zmax {
					// add scan positions around the current position, using average z
					avg_z := space.currentStack.zmin + (space.currentStack.zmax-space.currentStack.zmin)/2
					space.AddScanPositions(space.cx, space.cy, avg_z)
					space.currentStack = nil

					target = space.GetClosestScanPosition(space.cx, space.cy)
					if target == nil {
						fmt.Println("No more scan positions available, exiting...")
						return
					}
					space.state = StateMove
				} else {
					statusColor = colorYellow
					motorMove(port, AxisZ, DirDown)
					space.cz += 1
				}
			}

		case StateMove:
			statusColor = colorYellow
			if target.x > space.cx {
				motorMove(port, AxisX, DirUp)
				space.cx += 1
			} else if target.x < space.cx {
				motorMove(port, AxisX, DirDown)
				space.cx -= 1
			} else if target.y > space.cy {
				motorMove(port, AxisY, DirUp)
				space.cy += 1
			} else if target.y < space.cy {
				motorMove(port, AxisY, DirDown)
				space.cy -= 1
			} else if target.z > space.cz {
				motorMove(port, AxisZ, DirUp)
				space.cz += 1
			} else if target.z < space.cz {
				motorMove(port, AxisZ, DirDown)
				space.cz -= 1
			} else {
				// Reached the target position
				space.state = StateScanZUp
			}

		}
		gocv.Rectangle(&img, image.Rect(0, 0, img.Cols(), img.Rows()), statusColor, 20)
		window.IMShow(img)

		window.WaitKey(1)
	}
}
