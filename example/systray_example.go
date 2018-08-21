package main

import "github.com/amkulikov/systray"
import "path/filepath"
import "fmt"

func main() {
	systray.Run(onReady, onExit)
}

func onReady() {
	iconpath, _ := filepath.Abs("iconwin.ico")
	systray.SetIconPath(iconpath)
	mnew := systray.AddMenuItem("New", "", 0)
	mi1 := systray.AddMenuItem("New2", "", 0)
	mi, _ := systray.AddSubMenu("Submenu", "profile_submenu")
	mi2 := systray.AddSubMenuItem("profile_submenu", "SubmenuItem", "", 0)
	bmpfile, _ := filepath.Abs("stop.bmp")
	mi.AddBitmap(bmpfile)
	fmt.Println(bmpfile)
	bmpstart, _ := filepath.Abs("start.bmp")

	go func() {
		for {
			<-mnew.OnClickCh()
			mi.AddBitmap(bmpstart)
		}
	}()

	go func() {
		for {
			<-mi2.OnClickCh()
			mi.AddBitmap(bmpfile)
		}
	}()

	go func() {
		for {
			<-mi1.OnClickCh()
			mi.AddBitmap(bmpfile)
		}
	}()
}

func onExit() {

}
