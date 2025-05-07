package main

import (
	"GoRPC/GoRPC"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	a := app.NewWithID("GoRPC.preferences")
	w := a.NewWindow("GoRPC - By JustTimmm")
	w.Resize(fyne.NewSize(600, 0))
	w.SetFixedSize(true)
	GoRPC.SetupUI(a, w)
	w.ShowAndRun()
}
