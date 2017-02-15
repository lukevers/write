package main

import (
	"github.com/murlokswarm/app"
	"html/template"
)

var (
	// Templates contains all of the HTML templates used to build the
	// application. We load them up on boot instead of on render.
	Templates *template.Template
)

func init() {
	var err error
	Templates, err = template.ParseGlob(app.Storage().Resources() + "/html/*.html")
	if err != nil {
		Stderr.Println(err)
	}
}
