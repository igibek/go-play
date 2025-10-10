package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "method is not supported", http.StatusNotFound)
		return
	}
	fmt.Fprint(w, "hello!")
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprint(w, err)
		return
	}
	fmt.Fprint(w, "POST request successful")

	name := r.FormValue("name")
	address := r.FormValue("address")

	fmt.Fprintf(w, "Name %s\n", name)
	fmt.Fprintf(w, "Address: %s\n", address)
}

func main() {
	fmt.Println("simple http server project")
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	log.Fatal(http.ListenAndServe(":8000", nil))
}
