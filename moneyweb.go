package main

import "os"
import "net/http"
import "github.com/ishiikurisu/moneyweb/view"
// import "fmt"

func hello(w http.ResponseWriter, r *http.Request) {
    cookies := r.Cookies()
    username := r.FormValue("username")

    if len(cookies) == 0 && len(username) == 0 {
        view.SayHello(w)
    } else {
        view.SayWelcome(w)
    }

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
