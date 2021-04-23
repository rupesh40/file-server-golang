package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("listening on port 8080")
	myDir := http.Dir("./public")
	handler := http.FileServer(myDir)
	http.Handle("/", handler)

	http.ListenAndServe(":8080", nil)
}
