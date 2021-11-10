package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func stickynotes() {
    

    w := fyne.CurrentApp().NewWindow("Sticky Notes By Jahanwee")
    image := canvas.NewImageFromFile("./img/black.png")
    largeText := widget.NewMultiLineEntry()
    largeText.SetText("Write your tasks here")
    largeText.SetPlaceHolder("Type here")

    form := &widget.Form{
        Items: []*widget.FormItem{
        },
        OnCancel: func() {
            fmt.Println("Cancelled")
        },
        OnSubmit: func() {
            fmt.Println("Form submitted")

            
            // EVENT TO APPEND TO MULTILINE
			largeText := widget.NewMultiLineEntry()
    largeText.SetText("Enter your today's task")
    originalText := largeText.Text
    fmt.Println(originalText)
    newText := originalText + "appending new text"
    largeText.SetText(newText)
        },
    }

    w.SetContent(container.NewMax(image,
        container.NewVBox(form,largeText)),
    )

    w.Show()
}