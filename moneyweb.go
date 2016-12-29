package main

import "os"
import "net/http"
import "io"

func hello(w http.ResponseWriter, r *http.Request) {
    io.WriteString(w, "Hello from Money Log!")
}

func main() {
    port := os.Getenv("PORT")
    if len(port) == 0 {
        port = "8000"
    }
    http.HandleFunc("/", hello)
    http.ListenAndServe(":" + port, nil)
}
