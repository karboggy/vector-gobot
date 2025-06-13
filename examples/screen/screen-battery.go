package main

import (
	"fmt"
	"os"
	"time"

	"github.com/kercre123/vector-gobot/pkg/vbody"
	"github.com/kercre123/vector-gobot/pkg/vscreen"
)

var isMidas bool

func main() {
	fmt.Print("Initing body...")
	err := vbody.InitSpine()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	vbody.SetLEDs(0x000000, 0x000000, 0x000000)

	frameChan := vbody.GetFrameChan()
	fmt.Println(" FrameChan is OK")
	vbody.SetLEDs(0xA500FF, 0x000000, 0x000000)

	fmt.Print("Initing screen...")
	vscreen.InitLCD()
	isMidas, err = vscreen.IsMidas()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if isMidas {
		fmt.Println(" Vector 2.0 (midas)")
	} else {
		fmt.Println(" Vector 1.0")
	}
	vbody.SetLEDs(0xA500FF, 0xA500FF, 0xA500FF)

	fpsCount := 0
	fps := 0

	// update FPS value each second
	ticker := time.NewTicker(1 * time.Second)
	go func() {
	for range ticker.C {
		fps = fpsCount
		fpsCount = 0
	}
    }()

	for frame := range frameChan {
		// draw text
		fpsCount++
		batteryVotage := 0.00136719 * float32(frame.BattVoltage)
		message := fmt.Sprintf("Hello world!    FPS:%d   BATTERY: %02.2fV (raw: %d)", fps, batteryVotage, frame.BattVoltage)
		scrnData := vscreen.CreateTextImage(message)
		vscreen.SetScreen(scrnData)
	}
}
