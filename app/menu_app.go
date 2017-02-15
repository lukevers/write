package main

import (
	"bytes"
	"github.com/murlokswarm/app"
)

// MenuApp is a menu at the application level.
type MenuApp struct{}

func init() {
	app.RegisterComponent(&MenuApp{})
}

// Render compiles the template for the MenuApp
func (m *MenuApp) Render() string {
	var content bytes.Buffer
	err := Templates.ExecuteTemplate(&content, "menu-app", m)
	if err != nil {
		Stderr.Println(err)
	}

	return content.String()
}

// OpenAboutWindow opens a new PageAbout in a new window if a window for it
// does not already exist.
func (m *MenuApp) OpenAboutWindow() {
	if aboutWindow == nil {
		aboutWindow = newAboutWindow()
	}
}
