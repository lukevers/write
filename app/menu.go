package main

import (
	"bytes"
	"github.com/murlokswarm/app"
)

// MenuApp is a menu at the global level that contains all other menus.
type Menu struct{}

func init() {
	app.RegisterComponent(&Menu{})
}

// Render compiles the template for the Menu
func (m *Menu) Render() string {
	var content bytes.Buffer
	err := Templates.ExecuteTemplate(&content, "menu", m)
	if err != nil {
		Stderr.Println(err)
	}

	return content.String()
}
