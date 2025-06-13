package main

import (
	"fmt"
	"os"
	"time"
	"image"
	
	"github.com/kercre123/vector-gobot/pkg/vbody"
	"github.com/kercre123/vector-gobot/pkg/vscreen"
	"github.com/fogleman/gg"
)

var isMidas bool

func main() {
	// init vector body module
	fmt.Println("Initing body...")
	err := vbody.InitSpine()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	vbody.SetLEDs(0x000000, 0x000000, 0x000000)

	// init vector screen module
	fmt.Print("Initing screen...")
	vscreen.InitLCD()
	isMidas, err = vscreen.IsMidas()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	var W int
	var H int
	if isMidas {
		fmt.Println(" Vector 2.0 (midas)")
		W = 160
		H = 80
	} else {
		fmt.Println(" Vector 1.0")
		W = 184
		H = 96
	}
	vbody.SetLEDs(0xA500FF, 0xA500FF, 0xA500FF)

	// initialize image buffer target and gg context
	img := image.NewRGBA(image.Rect(0, 0, W, H))
	dc := gg.NewContextForRGBA(img)
	var cpt float64 = 10.0

	// main loop
	for {
		cpt += 1.0
		if int(cpt) > 50 {
			cpt = 0;
		}

		// clear black
		dc.SetRGB(0, 0, 0)
		dc.Clear()

		// draw a green circle
		dc.SetRGB(0, 1, 0)
		var x float64 = 10 + cpt
		var y float64 = 40
		var radius = 15 + cpt / 3
		dc.DrawCircle(x, y, radius)
		dc.Fill()

		// convert the color format from RGBA to RGB565
		pixels := make([]uint16, W*H)
		for y := 0; y < H; y++ {
			for x := 0; x < W; x++ {
				r, g, b, _ := img.At(x, y).RGBA()
				pixel := (r>>8&0xF8)<<8 | (g>>8&0xFC)<<3 | b>>8>>3
				pixels[y*W+x] = uint16(pixel)
			}
		}

		// update to screen
		vscreen.SetScreen(pixels)

		// refreshing 30 Hz
		time.Sleep(time.Second / 30)
	}
}
