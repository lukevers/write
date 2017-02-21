package main

import (
	"bytes"
	"fmt"
	"github.com/murlokswarm/app"
	"time"
)

// PageWelcome is a component page that is the first page seen to the user.
type PageWelcome struct {
	Page    int
	Animate bool
}

func init() {
	app.RegisterComponent(&PageWelcome{})
}

// Render compiles the template for to be displayed
func (p *PageWelcome) Render() string {
	var content bytes.Buffer
	page := fmt.Sprintf("welcome-%d", p.Page)
	err := Templates.ExecuteTemplate(&content, page, p)
	if err != nil {
		Stderr.Println(err)
	}

	return content.String()
}

// ClickGetStarted is called when the user clicks the get started button. It
// sets the animate flag to true and re-renders welcome-0. After waiting for a
// second, it then renders welcome-1.
func (p *PageWelcome) ClickGetStarted() {
	p.Animate = true
	app.Render(p)
	p.Page += 1
	time.Sleep(time.Second)
	app.Render(p)
}

// OnContextMenu is the function called when a user right clicks on this
// component.
func (p *PageWelcome) OnContextMenu() {
	ctxmenu := app.NewContextMenu()
	ctxmenu.Mount(&Menu{})
}
