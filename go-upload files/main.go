package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
)

var tpl *template.Template

func uploadFile(w http.ResponseWriter, r *http.Request) {
	fmt.Println("r.method : ", r.Method)
	if r.Method == "GET" { //if method is get then load form
		tpl.ExecuteTemplate(w, "fileUpload.html", nil)
		return
	}

	r.ParseMultipartForm(10) //parsing from form upto 10mb file
	file, fileHeader, err := r.FormFile("myFile")
	if err != nil {
		fmt.Println("err getting file : ", err)
		return
	}

	defer file.Close()
	fmt.Printf("fileHeader.filename : %v\n", fileHeader.Filename)
	fmt.Printf("fileHeader.Size : %v\n", fileHeader.Size)
	fmt.Printf("fileHeader.Header : %v\n", fileHeader.Header)

	contentType := fileHeader.Header["Content-Type"][0]
	fmt.Printf("content-type : %v\n", contentType)

	// defining  fileType var
	var osFile *os.File

	// creating the empty file according to Content-type of selected File
	if contentType == "image/jpeg" {
		osFile, err = ioutil.TempFile("images", "*.jpg")
	} else if contentType == "application/pdf" {
		osFile, err = ioutil.TempFile("PDFs", "*.pdf")
	} else if contentType == "text/javascript" {
		osFile, err = ioutil.TempFile("js", "*.js")
	}
	fmt.Println("err : ", err)
	defer osFile.Close()

	fileBytes, err := ioutil.ReadAll(file) // reading the selected file and creating into bytes and storing into fileBytes
	if err != nil {
		fmt.Println("err : ", err)
	}

	osFile.Write(fileBytes) // writing the data into OsFile

	fmt.Fprintf(w, "your file was successfully uploaded")
}
func main() {
	tpl, _ = tpl.ParseGlob("templates/*.html")
	http.HandleFunc("/upload", uploadFile)
	http.HandleFunc("/", handleHomepage)
	http.ListenAndServe(":8080", nil)
}

func handleHomepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to homepage  (go to '/upload') ")
}
