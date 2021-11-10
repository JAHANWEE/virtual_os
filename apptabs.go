package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func apptabs() {
	w := fyne.CurrentApp().NewWindow("App Store By Jahanwee")
	color_run:= widget.NewButton("color", func() {
		go colors()
	})

    snake_run:=widget.NewButton("snake",func() {
		go snake()
	})
	
	weather_run:=widget.NewButton("weather",func() {
		go weather()
	})
	sticky_run:= widget.NewButton("sticky",func() {
		go stickynotes()
	})
    tabs := container.NewAppTabs(
		container.NewTabItem("snake", snake_run),
		
		container.NewTabItem("weather", weather_run),

		container.NewTabItem("color", color_run),

		container.NewTabItem("sticky",sticky_run),
	)
	

	tabs.SetTabLocation(container.TabLocationTop)
   w.Resize(fyne.NewSize(500,100))
   w.CenterOnScreen()
	w.SetContent(tabs)
	w.Show()
}