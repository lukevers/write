package main

import (
	"github.com/murlokswarm/app"
)

var (
	win app.Contexter
)

func newMainWindow() app.Contexter {
	// Creates a window context.
	win := app.NewWindow(app.Window{
		Title:          "Write",
		Width:          1280,
		Height:         720,
		TitlebarHidden: true,
		OnClose: func() bool {
			win = nil
			return true
		},
	})

	welcome := &PageWelcome{}
	win.Mount(welcome)
	return win
}
