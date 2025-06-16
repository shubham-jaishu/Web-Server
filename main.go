package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/hello"{
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}
	if r.Method != "GET"{
		http.Error(w, "Method not supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "Hello")
}

func formHandler(w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/form"{
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}
	if r.Method != "POST"{
		http.Error(w, "Method not supported", http.StatusNotFound)
		return
	}
	if err := r.ParseForm(); err != nil{
		fmt.Fprintf(w, "ParseForm() err: %v\n", err)
		return
	}
	fmt.Fprintf(w, "POST req success\n")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "name is %s\n", name)
	fmt.Fprintf(w, "address is %s\n", address)
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)
	fmt.Println("starting server at 8000")
	if err := http.ListenAndServe(":8000", nil); err != nil{
		log.Fatal(err)
	}
}