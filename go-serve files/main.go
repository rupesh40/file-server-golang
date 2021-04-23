package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var tpl *template.Template

func main() {
	fmt.Println("hii")
	myDir := http.Dir("./public")
	handler := http.FileServer(myDir)
	http.Handle("/", handler)

	tpl, _ = tpl.ParseGlob("public/*.html")

	http.HandleFunc("/upload", uploadFile)

	http.ListenAndServe(":8080", nil)
}
func uploadFile(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.html", nil)
}
