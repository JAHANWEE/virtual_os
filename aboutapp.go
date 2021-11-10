package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

func aboutapp() {
	 // current app returns current application to the process
	w := fyne.CurrentApp().NewWindow("About")
	image := canvas.NewImageFromFile("./img/aboutme.png")


	w.SetContent(container.NewMax(image,
		
	))
  w.Resize(fyne.NewSize(400,500))
	w.Show()
}