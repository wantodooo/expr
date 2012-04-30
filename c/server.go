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
