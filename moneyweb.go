package main

import "os"
import "net/http"
import "github.com/ishiikurisu/moneyweb/view"
import "github.com/ishiikurisu/moneyweb/model"

func hello(w http.ResponseWriter, r *http.Request) {
    user := model.GetUser(w, r)
    if len(user) > 0 {
        view.SayWelcome(w)
    } else {
        view.SayHello(w)
    }
}

func signUp(w http.ResponseWriter, r *http.Request) {
    view.SignUp(w)
}

func register(w http.ResponseWriter, r *http.Request) {
    username := r.FormValue("username")
    password := r.FormValue("password")
    if model.RegisterUser(username, password) {
        w, r = model.AddCookie(w, r)
    }
    http.Redirect(w, r, "/", http.StatusFound)
}

func main() {
    port := os.Getenv("PORT")
    if len(port) == 0 {
        port = "8000"
    }
    http.HandleFunc("/", hello)
    http.HandleFunc("/sign_up", signUp)
    http.HandleFunc("/register", register)
    http.ListenAndServe(":" + port, nil)
}
