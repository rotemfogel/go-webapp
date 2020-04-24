package main

import (
	"html/template"
	"log"
	"net/http"
	"strings"
)

func main() {
	http.Handle("/img/", http.FileServer(http.Dir("public")))
	http.Handle("/css/", http.FileServer(http.Dir("public")))

	templates := populateTemplates()
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		f := request.URL.Path[1:]
		if len(f) == 0 {
			f = "home"
		}
		var t *template.Template
		if strings.HasSuffix(f, "html") {
			t = templates.Lookup(f)
		} else {
			t = templates.Lookup(f + ".html")
		}
		if t != nil {
			err := t.Execute(writer, nil)
			if err != nil {
				log.Println(err)
			}
		} else {
			writer.WriteHeader(http.StatusNotFound)
		}
	})
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

const (
	basePath = "templates"
)

func populateTemplates() *template.Template {
	result := template.New("templates")
	template.Must(result.ParseGlob(basePath + "/*.html"))
	return result
}
