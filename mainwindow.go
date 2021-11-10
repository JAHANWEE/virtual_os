package main

import (
	"fmt"
	

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	  a := app.New()
	  
	  w := a.NewWindow("OS By Jahanwee")
    
	image := canvas.NewImageFromFile("./img/mainbackground.png")
       
      

	r, _ := fyne.LoadResourceFromPath("./img/home.png")
	button1 := widget.NewButtonWithIcon(" ", r, func() {
		fmt.Println("Home")
		go clock()
		go showNotification(a)
		
	})

	button1.Resize(fyne.NewSize(50, 50))
	button1.Move(fyne.NewPos(525, 700))

    r1, _ := fyne.LoadResourceFromPath("./img/home.png")
	button2 := widget.NewButtonWithIcon(" ", r1, func() {
		
		go Notepad()
	})

	button2.Resize(fyne.NewSize(50, 50))
	button2.Move(fyne.NewPos(620, 700))
	r2, _ := fyne.LoadResourceFromPath("./img/home.png")
	button3 := widget.NewButtonWithIcon(" ", r2, func() {
          go gallery()
	})

	button3.Resize(fyne.NewSize(50, 50))
	button3.Move(fyne.NewPos(680,700))
	r3, _ := fyne.LoadResourceFromPath("./img/home.png")
	button4 := widget.NewButtonWithIcon(" ", r3, func() {
		go calculator()
	})

	button4.Resize(fyne.NewSize(50, 50))
	button4.Move(fyne.NewPos(780, 700))
	r4, _ := fyne.LoadResourceFromPath("./img/home.png")
	button5 := widget.NewButtonWithIcon(" ", r4, func() {
		go stickynotes()
	})

	button5.Resize(fyne.NewSize(50, 50))
	button5.Move(fyne.NewPos(860, 700)) 
	r5, _ := fyne.LoadResourceFromPath("./img/home.png")
	button6 := widget.NewButtonWithIcon(" ", r5, func() {
		go secondw()
	})

	button6.Resize(fyne.NewSize(50, 55))
	button6.Move(fyne.NewPos(1394, 300)) 
	
	
	w.SetContent(container.NewMax(
		image,
		container.NewWithoutLayout(
       button1, button2, button3, button4, button5 , button6,
		),
	))

    

	w.CenterOnScreen()
	 w. Resize(fyne.NewSize(1485, 768))
	w.ShowAndRun()
}