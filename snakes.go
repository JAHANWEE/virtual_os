package main

import (
	"image/color"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)
type snakepart struct{
	x,y float32
}

type movetype int
const(
	moveup movetype=iota
	movedown
	moveleft
	moveright
)
var move=moveup
var (
	snakeparts []snakepart
	game    *fyne.Container


)
func keytyped(e*fyne.KeyEvent){
	switch e.Name{
	case fyne.KeyUp:
		move=moveup
	case fyne.KeyDown:
		move=movedown
	case fyne.KeyLeft:
		move=moveleft
	case fyne.KeyRight:
		move=moveright
	}
}

func setupgame() *fyne.Container{
	var segments[]fyne.CanvasObject
	for i:=0;i<10;i++{
		r:=canvas.NewRectangle(&color.NRGBA{R:0 , G:255, B:0 , A:255})
		r.Resize(fyne.NewSize(10,10))
		r.Move(fyne.NewPos(90,float32(50+i*10)))
		seg:=snakepart{9,float32(5+i)}
		snakeparts=append(snakeparts, seg)
		segments=append(segments, r)

	}
	return container.NewWithoutLayout(segments...)
}
func refreshgame(){
	for i,seg:=range snakeparts{
		rect:=game.Objects[i]
		rect.Move(fyne.NewPos(seg.x*10,seg.y*10))

	}
	game.Refresh()
}
func rungame(){
	for{
		time.Sleep(time.Millisecond*250)    // for delay
		for i:=len(snakeparts)-1;i>=1;i--{
			snakeparts[i]=snakeparts[i-1]

		}
		switch move{
		case moveup:
			snakeparts[0].y--
		case movedown:
			snakeparts[0].y++
		case moveleft:
			snakeparts[0].x--
		case moveright:
			snakeparts[0].x++

		}
		refreshgame()
	}

}
func snake() {
	
	w := fyne.CurrentApp().NewWindow("Snake Game By Jahanwee")
	game=setupgame()
    
	w.SetContent(game)
	go rungame()
	w.Canvas().SetOnTypedKey(keytyped)
	w.Resize(fyne.NewSize(500,500))
	w.Show()
}