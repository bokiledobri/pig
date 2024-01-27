package main

import (
	"log"
	"net/http"
	"text/template"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	data := struct{Name string}{
		Name: "somename",
	}
    if r.Method != "GET" {
        w.Header().Set("Allow", "GET")
        http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
        return
    }
    if r.URL.Path !="/"{
        http.NotFound(w, r)
        return
    }
	ts, err := template.ParseFiles("ui/html/pages/home.tmpl")
	if err!=nil{
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	err = ts.Execute(w, data)
	if err!=nil{
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

