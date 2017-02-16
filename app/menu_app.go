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
