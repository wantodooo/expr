// This is free software. It comes without any warranty, to the extent
// permitted by applicable law. You can redistribute it and/or modify it under
// the termos of the Do What the Fuck You Want To Public License, Version 2,
// as published by Sam Hocevar. See COPYING for more details.

package main

import (
	"flag"
	"github.com/zephyrtronium/expr"
	"html/template"
	"log"
	"net/http"
)

var (
	templ *template.Template

	serveHost, serveRoute string
	templateFile string
)

func CalcServer(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	if q := req.Form.Get("q"); q != "" {
		if x, err := expr.EvalString(q); err == nil {
			req.Form.Set("q", x.String())
		} else {
			req.Form.Set("e", err.Error())
		}
	}
	req.Form.Set("r", serveRoute)
	templ.Execute(w, req.Form)
}

func main() {
	flag.StringVar(&serveHost, "http", ":3123", "http host")
	flag.StringVar(&serveRoute, "route", "/c", "route to serve")
	flag.StringVar(&templateFile, "template", "serve.html", "HTML template")
	flag.Parse()
	templ = template.Must(template.ParseFiles(templateFile))
	http.HandleFunc(serveRoute, CalcServer)
	err := http.ListenAndServe(serveHost, nil)
	if err != nil {
		log.Fatalln("ListenAndServe:", err)
	}
}
