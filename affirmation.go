package main

import (
	"time"

	"fyne.io/fyne/v2"
)


func showNotification(a fyne.App) {
	
	time.Sleep(time.Second * 2)
	a.SendNotification(fyne.NewNotification("Affirmations", "You are enough."))   // sends notifications

	time.Sleep(time.Second * 6)

	a.SendNotification(fyne.NewNotification("Affirmations", "You are best "))

	time.Sleep(time.Second * 10)  // for delay

	a.SendNotification(fyne.NewNotification("Affirmations", "You are best "))

	

}