package main
// import fyne
import (
    "image/color"
    "math/rand"
    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/canvas"
    "fyne.io/fyne/v2/container"
    "fyne.io/fyne/v2/widget"
)
func colors() {

	w := fyne.CurrentApp().NewWindow("Color picker By Jahanwee")
    w.Resize(fyne.NewSize(400, 400))

    
    colorx := color.NRGBA{R: 0, G: 0, B: 0, A: 255}
    rect1 := canvas.NewRectangle(colorx)
    rect1.SetMinSize(fyne.NewSize(200, 200))


    // Btn for color change here
    btn1 := widget.NewButton("Random color", func() {
        

        rect1.FillColor = color.NRGBA{R: uint8(rand.Intn(255)),        // for generating random colors
            G: uint8(rand.Intn(255)), B: uint8(rand.Intn(255)), A: 255}
        rect1.Refresh() 
    })
    //  RED COLOR
    btnRed := widget.NewButton("RED", func() {
        
        rect1.FillColor = color.NRGBA{R: uint8(rand.Intn(255)),
            G: 0, B: 0, A: 255}
        rect1.Refresh() 
    })
    // GREEN COLOR
    btnGreen := widget.NewButton("Green", func() {
       
        rect1.FillColor = color.NRGBA{R: 0,
            G: uint8(rand.Intn(255)), B: 0, A: 255}
        rect1.Refresh() 
    })
    //  Blue COLOR
    btnBlue := widget.NewButton("Blue", func() {
        




        rect1.FillColor = color.NRGBA{R: 0,
            G: 0, B: uint8(rand.Intn(255)), A: 255}
        rect1.Refresh() 
    })
    w.SetContent(
        container.NewVBox(
            rect1,
            btn1,
            btnRed,
            btnGreen,
            btnBlue,
        ),
    )
    w.Show()
}