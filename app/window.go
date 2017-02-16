package main

import (
	"github.com/murlokswarm/app"
)

var (
	mainWindow app.Contexter
)

func newWelcomeWindow() app.Contexter {
	w := app.NewWindow(app.Window{
		Title:          "Write",
		Width:          400,
		Height:         580,
		FixedSize:      true,
		TitlebarHidden: true,

		OnClose: func() bool {
			mainWindow = nil
			return true
		},
	})

	w.Mount(&PageWelcome{})

	return w
}
