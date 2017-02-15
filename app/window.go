package main

import (
	"github.com/murlokswarm/app"
)

var (
	mainWindow  app.Contexter
	aboutWindow app.Contexter
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

func newAboutWindow() app.Contexter {
	w := app.NewWindow(app.Window{
		Width:          200,
		Height:         100,
		FixedSize:      true,
		MinimizeHidden: true,

		OnClose: func() bool {
			aboutWindow = nil
			return true
		},
	})

	w.Mount(&PageAbout{
		Version: Version,
		Commit:  Commit[:8],
	})

	return w
}
