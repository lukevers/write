package main

import (
	"github.com/murlokswarm/app"
)

func newWelcomeWindow() app.Contexter {
	w := app.NewWindow(app.Window{
		Title:          "Write",
		Width:          400,
		Height:         580,
		FixedSize:      true,
		TitlebarHidden: true,
		OnClose:        onWindowClose,
	})

	page := &PageWelcome{}
	w.Mount(page)
	return w
}

func onWindowClose() bool {
	window = nil
	return true
}
