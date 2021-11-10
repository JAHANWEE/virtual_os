package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

func gallery() {
	



	w := fyne.CurrentApp().NewWindow("Gallery By Jahanwee")

	root_src := "C:\\Users\\jahan\\Desktop\\pictures"

	files, err := ioutil.ReadDir(root_src)
	if err != nil {
		log.Fatal(err)
	}
	tabs := container.NewAppTabs()

	for _, file := range files {
		fmt.Println(file.Name(), file.IsDir())
		if !file.IsDir() {
			extention := strings.Split(file.Name(), ".")[1]
			if extention == "png" || extention == "jpeg" {
				image := canvas.NewImageFromFile(root_src + "\\" + file.Name())
				image.FillMode = canvas.ImageFillOriginal
				tabs.Append(container.NewTabItem(file.Name(), image))
			}
		}
	}

	tabs.SetTabLocation(container.TabLocationLeading)
	
	w.Resize(fyne.NewSize(500, 500))
	w.SetContent(tabs)
	w.CenterOnScreen()
	w.Show()
}
