package main

import (
	"ascii-art/ascii-art-web/server"
	"fmt"
	"log"
	"net/http"
)

func main() {

	//handeling the root
	http.HandleFunc("/", server.HomeHandler)
	http.HandleFunc("/ascii-art", server.Submit)
	fmt.Printf("http://localhost:8000/\n")
	//check if there is an error in serving this port
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
