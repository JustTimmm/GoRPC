package GoRPC

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/hugolgst/rich-go/client"
)

func SetupUI(a fyne.App, w fyne.Window) {
	prefs := a.Preferences()

	entries := map[string]*widget.Entry{}
	fieldKeys := []string{
		"idEntry", "detailsEntry", "stateEntry",
		"largeImageKeyEntry", "largeImageTextEntry",
		"smallImageKeyEntry", "smallImageTextEntry",
		"button1Label", "button1Url", "button2Label", "button2Url",
	}

	for _, key := range fieldKeys {
		entries[key] = newBoundEntry(key, prefs)
	}

	progress := widget.NewProgressBar()
	startButton := widget.NewButton("START", func() {
		StartRPC(RPC{
			ClientID:       entries["idEntry"].Text,
			Details:        entries["detailsEntry"].Text,
			State:          entries["stateEntry"].Text,
			LargeImageKey:  entries["largeImageKeyEntry"].Text,
			LargeImageText: entries["largeImageTextEntry"].Text,
			SmallImageKey:  entries["smallImageKeyEntry"].Text,
			SmallImageText: entries["smallImageTextEntry"].Text,
			Button1Label:   entries["button1Label"].Text,
			Button1URL:     entries["button1Url"].Text,
			Button2Label:   entries["button2Label"].Text,
			Button2URL:     entries["button2Url"].Text,
		})
	})
	startButton.Disable()
	stopButton := widget.NewButton("STOP", func() { client.Logout() })

	updateProgress := func() {
		filled := 0
		for _, e := range entries {
			if e.Text != "" {
				filled++
			}
		}
		progress.SetValue(float64(filled) / float64(len(entries)))
		if filled == len(entries) {
			startButton.Enable()
		} else {
			startButton.Disable()
		}
	}

	for key, entry := range entries {
		entry.OnChanged = func(k string, e *widget.Entry) func(string) {
			return func(val string) {
				prefs.SetString(k, val)
				updateProgress()
			}
		}(key, entry)
	}

	box := container.NewVBox(
		widget.NewLabelWithStyle("DiscordRCP Custom", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		entries["idEntry"],
		separator(),
		entries["detailsEntry"], entries["stateEntry"],
		separator(),
		container.NewGridWithColumns(2,
			entries["largeImageKeyEntry"],
			entries["largeImageTextEntry"],
		),
		container.NewGridWithColumns(2,
			entries["smallImageKeyEntry"],
			entries["smallImageTextEntry"],
		),
		separator(),
		container.NewGridWithColumns(2,
			entries["button1Label"],
			entries["button1Url"],
		),
		container.NewGridWithColumns(2,
			entries["button2Label"],
			entries["button2Url"],
		),
		separator(),
		progress,
		container.NewGridWithColumns(2,
			startButton,
			stopButton,
		),
	)

	w.SetContent(box)
	updateProgress()
}
