package GoRPC

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/hugolgst/rich-go/client"
	"image/color"
)

func appendIfValid(buttons []*client.Button, label, url string) []*client.Button {
	if label != "" && url != "" {
		buttons = append(buttons, &client.Button{Label: label, Url: url})
	}
	return buttons
}

func newBoundEntry(key string, prefs fyne.Preferences) *widget.Entry {
	entry := widget.NewEntry()
	entry.SetPlaceHolder(key)
	entry.SetText(prefs.StringWithFallback(key, ""))
	return entry
}

func separator() *fyne.Container {
	sep := canvas.NewRectangle(color.White)
	sep.SetMinSize(fyne.NewSize(510, 2))
	return container.NewPadded(sep)
}
