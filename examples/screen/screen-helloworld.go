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
	fmt.Println("Initing body...")
	err := vbody.InitSpine()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	vbody.SetLEDs(0x000000, 0x000000, 0x000000)

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

	counter := 0
	for {
		// draw text
		counter++
		message := fmt.Sprintf("Hello world! (%d)", counter)
		scrnData := vscreen.CreateTextImage(message)
		vscreen.SetScreen(scrnData)

		// refreshing 30 Hz
		time.Sleep(time.Second / 30)
	}
}
