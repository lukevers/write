package main

import (
	"bytes"
	"github.com/murlokswarm/app"
)

// MenuWindow is a menu at the window level.
type MenuWindow struct{}

func init() {
	app.RegisterComponent(&MenuWindow{})
}

// Render compiles the template for the MenuWindow
func (m *MenuWindow) Render() string {
	var content bytes.Buffer
	err := Templates.ExecuteTemplate(&content, "menu-window", m)
	if err != nil {
		Stderr.Println(err)
	}

	return content.String()
}
