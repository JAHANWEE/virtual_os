package main

import (
	
	"os/exec"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func secondw() {
	

	w := fyne.CurrentApp().NewWindow("OS By Jahanwee")

	image := canvas.NewImageFromFile("./img/backgrounds.png")


	r, _ := fyne.LoadResourceFromPath("./img/home.png")
	button1 := widget.NewButtonWithIcon(" ", r, func() {
		go apptabs()
		go clock()
	})
	

	button1.Resize(fyne.NewSize(50, 50))
	button1.Move(fyne.NewPos(530, 705))

    r1, _ := fyne.LoadResourceFromPath("./img/home.png")
	button2 := widget.NewButtonWithIcon(" ", r1, func() {
		go Notepad()
	})

	button2.Resize(fyne.NewSize(50, 50))
	button2.Move(fyne.NewPos(622, 705))
	r2, _ := fyne.LoadResourceFromPath("./img/home.png")
	button3 := widget.NewButtonWithIcon(" ", r2, func() {
		go gallery()
	})

	button3.Resize(fyne.NewSize(50, 50))
	button3.Move(fyne.NewPos(712,705))
	r3, _ := fyne.LoadResourceFromPath("./img/home.png")
	button4 := widget.NewButtonWithIcon(" ", r3, func() {
		go calculator()
	})

	button4.Resize(fyne.NewSize(50, 50))
	button4.Move(fyne.NewPos(800, 705))
	r4, _ := fyne.LoadResourceFromPath("./img/home.png")
	button5 := widget.NewButtonWithIcon(" ", r4, func() {
		go drum()
	})

	button5.Resize(fyne.NewSize(50, 50))
	button5.Move(fyne.NewPos(883, 705)) 
	 


	 
	 // header buttons .............



	 h, _ := fyne.LoadResourceFromPath("./img/home.png")
	hbutton1 := widget.NewButtonWithIcon("Github", h, func() {
		exec.Command("rundll32", "url.dll,FileProtocolHandler", "https://github.com/JAHANWEE").Start()
	})

	hbutton1.Resize(fyne.NewSize(100, 30))
	hbutton1.Move(fyne.NewPos(5, 5))
	h1, _ := fyne.LoadResourceFromPath("./img/home.png")
	hbutton2 := widget.NewButtonWithIcon("Whatsapp", h1, func() {
		exec.Command("rundll32", "url.dll,FileProtocolHandler", "https://web.whatsapp.com/").Start()
	})

	hbutton2.Resize(fyne.NewSize(100, 30))
	hbutton2.Move(fyne.NewPos(115, 5))
	h2, _ := fyne.LoadResourceFromPath("./img/home.png")
	hbutton3 := widget.NewButtonWithIcon("Youtube ", h2, func() {
		exec.Command("rundll32", "url.dll,FileProtocolHandler", "https://www.youtube.com").Start()
	})

	hbutton3.Resize(fyne.NewSize(100, 30))
	hbutton3.Move(fyne.NewPos(225, 5))
	h3, _ := fyne.LoadResourceFromPath("./img/home.png")
	hbutton4 := widget.NewButtonWithIcon("Gmail ", h3, func() {
		exec.Command("rundll32", "url.dll,FileProtocolHandler", "https://mail.google.com").Start()
	})

	hbutton4.Resize(fyne.NewSize(100, 30))
	hbutton4.Move(fyne.NewPos(335, 5))
	  
	h4, _ := fyne.LoadResourceFromPath("./img/home.png")
	hbutton5 := widget.NewButtonWithIcon("Google ", h4, func() {
		exec.Command("rundll32", "url.dll,FileProtocolHandler", "https://www.google.com").Start()  // handles url and opens url in the browser
	})

	hbutton5.Resize(fyne.NewSize(100, 30))
	hbutton5.Move(fyne.NewPos(445, 5))

	h5, _ := fyne.LoadResourceFromPath("./img/home.png")
	hbutton6 := widget.NewButtonWithIcon("Linklden ", h5, func() {
		exec.Command("rundll32", "url.dll,FileProtocolHandler", "https://www.linkedin.com").Start()
	})

	hbutton6.Resize(fyne.NewSize(100, 30))
	hbutton6.Move(fyne.NewPos(555, 5))
	  
	h6, _ := fyne.LoadResourceFromPath("./img/home.png")
	hbutton7 := widget.NewButtonWithIcon("Pepcoding", h6, func() {
		exec.Command("rundll32", "url.dll,FileProtocolHandler", "https://pepcoding.com").Start()
	})

	hbutton7.Resize(fyne.NewSize(100, 30))
	hbutton7.Move(fyne.NewPos(665,5))

	h7, _ := fyne.LoadResourceFromPath("./img/home.png")
	hbutton8 := widget.NewButtonWithIcon("GFG ", h7, func() {
		exec.Command("rundll32", "url.dll,FileProtocolHandler", "https://www.geeksforgeeks.org").Start()
	})

	hbutton8.Resize(fyne.NewSize(100, 30))
	hbutton8.Move(fyne.NewPos(775, 5))

	h8, _ := fyne.LoadResourceFromPath("./img/home.png")
	hbutton9 := widget.NewButtonWithIcon("Codechef ", h8, func() {
		exec.Command("rundll32", "url.dll,FileProtocolHandler", "https://www.codechef.com").Start()
	})

	hbutton9.Resize(fyne.NewSize(100, 30))
	hbutton9.Move(fyne.NewPos(885, 5))

	h9, _ := fyne.LoadResourceFromPath("./img/home.png")
	hbutton10 := widget.NewButtonWithIcon("Twitter ", h9, func() {
		exec.Command("rundll32", "url.dll,FileProtocolHandler", "https://twitter.com").Start()
	})

	hbutton10.Resize(fyne.NewSize(100, 30))
	hbutton10.Move(fyne.NewPos(995, 5))

	h10, _ := fyne.LoadResourceFromPath("./img/home.png")
	hbutton11 := widget.NewButtonWithIcon("Canva ", h10, func() {
		exec.Command("rundll32", "url.dll,FileProtocolHandler", "https://www.canva.com").Start()
	})

	hbutton11.Resize(fyne.NewSize(100, 30))
	hbutton11.Move(fyne.NewPos(1105, 5))

	h11, _ := fyne.LoadResourceFromPath("./img/home.png")
	hbutton12 := widget.NewButtonWithIcon("Fyne ", h11, func() {
		exec.Command("rundll32", "url.dll,FileProtocolHandler", "https://fyne.io").Start()
	})

	hbutton12.Resize(fyne.NewSize(100, 30))
	hbutton12.Move(fyne.NewPos(1220, 5))

	h12, _ := fyne.LoadResourceFromPath("./img/home.png")
	hbutton13 := widget.NewButtonWithIcon(" About", h12, func() {
		go aboutapp()
	})

	hbutton13.Resize(fyne.NewSize(100, 30))
	hbutton13.Move(fyne.NewPos(1330, 5))

	

	// header buttons end here ..........


	

	w.SetContent(container.NewMax(
		image,
		container.NewWithoutLayout(
       button1, button2, button3, button4, button5 ,hbutton1 , hbutton2, hbutton3,hbutton4, hbutton5, hbutton6,hbutton7,hbutton8,
	   hbutton9, hbutton10, hbutton11, hbutton12, hbutton13,
		),
	))
    

	w.CenterOnScreen()
	 w. Resize(fyne.NewSize(1485, 768))
	w.Show()
}