package main

import "os"
import "net/http"
import "github.com/ishiikurisu/moneyweb/view"

func hello(w http.ResponseWriter, r *http.Request) {
    view.SayHello(w)
}

func signUp(w http.ResponseWriter, r *http.Request) {
    view.SignUp(w)
}

func main() {
    port := os.Getenv("PORT")
    if len(port) == 0 {
        port = "8000"
    }
    http.HandleFunc("/", hello)
    http.HandleFunc("/sign_up", signUp)
    http.ListenAndServe(":" + port, nil)
}
