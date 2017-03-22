package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"log"
	"net/http"
)

func renderTemplate(name string, w http.ResponseWriter) (err error) {
	tmpl, err := template.New(name).ParseFiles(name + ".html")
	if err != nil {
		return err
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		return err
	}

	return nil
}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	err := renderTemplate("index", w)
	if err != nil {
		fmt.Fprint(w, err.Error())
	}
}

func About(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	err := renderTemplate("about", w)
	if err != nil {
		fmt.Fprint(w, err.Error())
	}
}

func Contact(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	err := renderTemplate("contact", w)
	if err != nil {
		fmt.Fprint(w, err.Error())
	}
}

func main() {
	fmt.Println("Starting web server.")
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/about", About)
	router.GET("/contact", Contact)
	log.Fatal(http.ListenAndServe(":80", router))
}
