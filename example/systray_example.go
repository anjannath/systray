package main

import (
	"fmt"
	"path/filepath"

	"github.com/amkulikov/systray"
)

var (
	menuTitles          = []string{"minishift", "kuberenets", "kubedash", "kvirt"}
	submenus            = make(map[string]*systray.MenuItem)
	submenusToMenuItems = make(map[string]MenuAction)
)

func main() {
	systray.Run(onReady, onExit)
}

type MenuAction struct {
	start *systray.MenuItem
	stop  *systray.MenuItem
}

func onReady() {
	iconpath, _ := filepath.Abs("minishift.ico")
	systray.SetIconPath(iconpath)
	exit := systray.AddMenuItem("Exit", "", 0)
	systray.AddSeparator()
	for _, menuTitle := range menuTitles {
		submenu := systray.AddSubMenu(menuTitle)
		startMenu := submenu.AddSubMenuItem("Start", "", 0)
		stopMenu := submenu.AddSubMenuItem("Stop", "", 0)
		iconpath, _ := filepath.Abs("doesnotexist.bmp")
		submenu.AddBitmap(iconpath)
		submenus[menuTitle] = submenu
		submenusToMenuItems[menuTitle] = MenuAction{start: startMenu, stop: stopMenu}
	}

	go func() {
		<-exit.OnClickCh()
		systray.Quit()
	}()

	iconStart, _ := filepath.Abs("running.bmp")
	iconStop, _ := filepath.Abs("stopped.bmp")

	for k, v := range submenusToMenuItems {
		fmt.Println(k)
		go func(iconpath, submenu string, v MenuAction) {
			for {
				<-v.start.OnClickCh()
				submenus[submenu].AddBitmap(iconpath)
			}
		}(iconStart, k, v)

		go func(iconpath, submenu string, v MenuAction) {
			for {
				<-v.stop.OnClickCh()
				submenus[submenu].AddBitmap(iconpath)
			}
		}(iconStop, k, v)
	}

}

func onExit() {

}
