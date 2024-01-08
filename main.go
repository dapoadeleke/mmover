package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
	"github.com/go-vgo/robotgo"
	"time"
)

var status = "INACTIVE"

func main() {
	a := app.New()
	w := a.NewWindow("mMover")
	w.Resize(fyne.NewSize(300, 100))

	statusText := binding.NewString()
	statusText.Set("Status: " + status)

	statusLabel := widget.NewLabelWithData(statusText)

	startButton := widget.NewButton("Start", func() {
		status = "ACTIVE"
		statusText.Set("Status: " + status)
		go moveMouseEvery5Minutes()
		//log.Println("started ", status)
	})

	stopButton := widget.NewButton("Stop", func() {
		stopMovingMouse()
		statusText.Set("Status: " + status)
		//log.Println("stopped ", status)
	})

	w.SetContent(container.NewVBox(
		statusLabel,
		startButton,
		stopButton,
	))

	w.ShowAndRun()
}

func moveMouseEvery5Minutes() {
	sx, sy := robotgo.GetScreenSize()
	for {
		if status == "INACTIVE" {
			break
		}
		time.Sleep(4 * time.Minute)
		//log.Println("Moving mouse")
		sx = sx - 1
		robotgo.MoveMouse(sx, sy)
		//log.Println("move mouse", sx, sy)
	}
}

func stopMovingMouse() {
	status = "INACTIVE"
}
