package main

import (
	"bytes"
	"github.com/murlokswarm/app"
)

type PageWelcome struct {
	Greeting string
}

func init() {
	app.RegisterComponent(&PageWelcome{})
}

func (p *PageWelcome) Render() string {
	var content bytes.Buffer
	err := Templates.ExecuteTemplate(&content, "welcome", p)
	if err != nil {
		Stderr.Println(err)
	}

	return content.String()
}

func (p *PageWelcome) OnInputChange(arg app.ChangeArg) {
	p.Greeting = arg.Value
	app.Render(p)
}

func (p *PageWelcome) OnContextMenu() {
	ctxmenu := app.NewContextMenu()
	ctxmenu.Mount(&MenuApp{})
}
