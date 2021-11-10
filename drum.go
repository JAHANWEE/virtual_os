package main

import (
	"os"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

func drum() {
	
	w := fyne.CurrentApp().NewWindow("Notepad By Jahanwee")

    image := canvas.NewImageFromFile("./drum.png")   // canvas package lets us add images
	


	r, _ := fyne.LoadResourceFromPath("./img/home.png")
	
	button1 := widget.NewButtonWithIcon(" ", r, func() {
		f, _ := os.Open("./sounds/kick-bass.mp3")
	streamer, format, _ := mp3.Decode(f)
	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	speaker.Play(streamer)
		
	})

	button1.Resize(fyne.NewSize(100, 100))
	button1.Move(fyne.NewPos(200, 200))

	r1, _ := fyne.LoadResourceFromPath("./img/home.png")
	button2:= widget.NewButtonWithIcon(" ", r1, func() {
		f, _ := os.Open("./sounds/crash.mp3")
	streamer, format, _ := mp3.Decode(f)
	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	speaker.Play(streamer)
		
	})

	button2.Resize(fyne.NewSize(100, 100))
	button2.Move(fyne.NewPos(400, 200))
	
	r3, _ := fyne.LoadResourceFromPath("./img/home.png")
	button3 := widget.NewButtonWithIcon(" ", r3, func() {
		f, _ := os.Open("./sounds/snare.mp3")
	streamer, format, _ := mp3.Decode(f)
	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	speaker.Play(streamer)
		
	})

	button3.Resize(fyne.NewSize(100, 100))
	button3.Move(fyne.NewPos(650, 200))
	

	r4, _ := fyne.LoadResourceFromPath("./img/home.png")
	button4 := widget.NewButtonWithIcon(" ", r4, func() {
		f, _ := os.Open("./sounds/tom-1.mp3")
	streamer, format, _ := mp3.Decode(f)
	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	speaker.Play(streamer)
		
	})

	button4.Resize(fyne.NewSize(100, 100))
	button4.Move(fyne.NewPos(850, 200))
	
	r5, _ := fyne.LoadResourceFromPath("./img/home.png")
	button5 := widget.NewButtonWithIcon(" ", r5, func() {
		f, _ := os.Open("./sounds/tom-2.mp3")
	streamer, format, _ := mp3.Decode(f)
	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	speaker.Play(streamer)
		
	})

	button5.Resize(fyne.NewSize(100, 100))
	button5.Move(fyne.NewPos(1100, 200))

	r6, _ := fyne.LoadResourceFromPath("./img/home.png")
	button6 := widget.NewButtonWithIcon(" ", r6, func() {
		f, _ := os.Open("./sounds/tom-3.mp3")
	streamer, format, _ := mp3.Decode(f)
	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	speaker.Play(streamer)
		
	})

	button6.Resize(fyne.NewSize(100, 100))
	button6.Move(fyne.NewPos(500, 350))

	r7, _ := fyne.LoadResourceFromPath("./img/home.png")
	button7 := widget.NewButtonWithIcon(" ", r7, func() {
		f, _ := os.Open("./sounds/tom-4.mp3")
	streamer, format, _ := mp3.Decode(f)
	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	speaker.Play(streamer)
		
	})

	button7.Resize(fyne.NewSize(100, 100))
	button7.Move(fyne.NewPos(730, 350))
	
	
	w.SetContent(container.NewMax(
		image,
		container.NewWithoutLayout(
       button1, button2 ,button3, button4,button5, button6,button7,
		),
	),
)
    w. Resize(fyne.NewSize(1700, 700))
	w.Show()
}