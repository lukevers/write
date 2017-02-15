package main

import (
	"bytes"
	"github.com/murlokswarm/app"
)

type PageAbout struct {
	Version string
	Commit  string
}

func init() {
	app.RegisterComponent(&PageAbout{})
}

func (p *PageAbout) Render() string {
	var content bytes.Buffer
	err := Templates.ExecuteTemplate(&content, "about", p)
	if err != nil {
		Stderr.Println(err)
	}

	return content.String()
}
