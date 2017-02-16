package main

import (
	"github.com/murlokswarm/app"
	_ "github.com/murlokswarm/mac"
)

func main() {
	app.OnLaunch = onLaunch
	app.OnReopen = onReopen
	app.OnFocus = onFocus
	app.OnBlur = onBlur
	app.OnFileOpen = onFileOpen
	app.OnTerminate = onTerminate
	app.OnFinalize = onFinalize
	app.Run()
}

// OnLaunch is a handler which is called when the app is initialized and ready.
func onLaunch() {
	menu := &Menu{}

	app.MenuBar().Mount(menu)
	app.Dock().Mount(menu)

	mainWindow = newWelcomeWindow()
}

// OnReopen is a handler which is called when the app is reopened.
func onReopen(hasVisibleWindow bool) {
	if mainWindow != nil {
		return
	}

	mainWindow = newWelcomeWindow()
}

// OnFocus is a handler which is called when the app became focused.
func onFocus() {
	//
}

// OnBlur is a handler which is called when the app lost the focus.
func onBlur() {
	//
}

// OnFileOpen is a handler which is called when a file is targeted to be opened
// with the app.
func onFileOpen(filename string) {
	//
}

// OnTerminate is a handler which is called when the app is requested to be
// terminated. Return false cancels the termination request.
func onTerminate() bool {
	return true
}

// OnFinalize is a handler which is called when the app is about to be
// terminated. It should be used to perform any final cleanup before the
// application terminates.
func onFinalize() {
	//
}
