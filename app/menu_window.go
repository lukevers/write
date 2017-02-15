package main

import (
	"bytes"
	"github.com/murlokswarm/app"
)

type MenuWindow struct{}

func init() {
	app.RegisterComponent(&MenuWindow{})
}

func (m *MenuWindow) Render() string {
	var content bytes.Buffer
	err := Templates.ExecuteTemplate(&content, "menu-window", m)
	if err != nil {
		Stderr.Println(err)
	}

	return content.String()
}
