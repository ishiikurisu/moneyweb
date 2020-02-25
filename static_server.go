package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	go func() {
		log.Fatal(http.ListenAndServe(":8000", http.FileServer(http.Dir("./static"))))
	}()
	fmt.Scanf("\n")
}
