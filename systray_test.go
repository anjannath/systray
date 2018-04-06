package systray

import (
	"log"
	"testing"
	"time"
)

func TestRun(t *testing.T) {
	onReady := func() {
		if err := SetIconPath("example/icon/iconwin.ico"); err != nil {
			t.Fatalf("Can't set icon: %s", err)
		}
		if err := SetTitle("Test title с кириллицей"); err != nil {
			t.Fatalf("Can't set title: %s", err)
		}

		bSomeBtn, err := AddMenuItem("Йа Кнопко", "")
		if err != nil {
			t.Fatalf("Can't add button: %s", err)
		}
		if err := bSomeBtn.Check(); err != nil {
			t.Fatalf("Can't check button: %s", err)
		}
		if err := AddSeparator(); err != nil {
			t.Fatalf("Can't add separator: %s", err)
		}
		bQuit, err := AddMenuItem("Quit", "Quit the whole app")
		if err != nil {
			t.Fatalf("Can't add button: %s", err)
		}
		go func() {
			<-bQuit.ClickedCh
			t.Log("Quit reqested")
			Quit()
		}()
	}

	time.AfterFunc(3*time.Second, Quit)
	Run(onReady, nil)
}

func ExampleRun() {
	onReady := func() {
		if err := SetIconPath("example/icon/iconwin.ico"); err != nil {
			log.Fatalf("Can't set icon: %s", err)
		}
		if err := SetTitle("Test title с кириллицей"); err != nil {
			log.Fatalf("Can't set title: %s", err)
		}

		bSomeBtn, _ := AddMenuItem("Йа Кнопко", "")
		if err := bSomeBtn.Check(); err != nil {
			log.Fatalf("Can't check button: %s", err)
		}
		if err := AddSeparator(); err != nil {
			log.Fatalf("Can't add separator: %s", err)
		}
		bQuit, _ := AddMenuItem("Quit", "Quit the whole app")
		go func() {
			<-bQuit.ClickedCh
			log.Println("Quit requested")
			Quit()
		}()
	}

	Run(onReady, nil)
}
