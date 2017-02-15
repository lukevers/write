package main

import (
	"bytes"
	"github.com/murlokswarm/app"
)

type MenuApp struct {
	CustomTitle string
	Disabled    bool
}

func init() {
	app.RegisterComponent(&MenuApp{})
}

func (m *MenuApp) Render() string {
	var content bytes.Buffer
	err := Templates.ExecuteTemplate(&content, "menu-app", m)
	if err != nil {
		Stderr.Println(err)
	}

	return content.String()
}

func (m *MenuApp) OnCustomMenuClick() {
	Stdout.Println("OnCustomMenuClick")
}
