package main

import (
	"fmt"
	"log"
	"net/http"
)


func helloHandler(w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/form"{
		http.Error(w,"404 not found", http.StatusNotFound)
		return
	}

	if r.Method != "GET"{
		http.Error(w,"Method not supported", http.StatusNotFound)
	}
	fmt.Fprintf(w, "hello")
}
func formHandler(w http.ResponseWriter, r *http.Request){
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "parse form failed: %v", err);
		return
	}
	fmt.Fprintf(w,"POST request Successfully")
	name := r.Form["name"];
	address := r.Form["address"];
	fmt.Fprintf(w, "name: %v, address: %v", name, address);
}


func main()  {
	fileServer := http.FileServer(http.Dir("/static"));
	http.Handle("/",fileServer)
	http.HandleFunc("/form",formHandler)
	http.HandleFunc("/hello",helloHandler)

	fmt.Printf("server is connected at port 8080");
	if err := http.ListenAndServe(":8080", nil); err != nil{
		log.Fatal(err)
	}
}