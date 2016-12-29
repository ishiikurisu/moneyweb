package main

import "os"
import "net/http"
import "github.com/ishiikurisu/moneyweb/view"
import "fmt"

func hello(w http.ResponseWriter, r *http.Request) {
    cookies := r.Header.Get("username")
    username := r.FormValue("username")

    if len(cookies) == 0 && len(username) == 0 {
        view.SayHello(w)
    } else if len(cookies) == 0 && len(username) > 0 {
        password := r.FormValue("password")
        rawCookie := fmt.Sprintf("%s=%s", username, password)
        cookie := http.Cookie{
            Name: "username",
            Value: username,
            MaxAge: 86400,
            Secure: true,
            HttpOnly: false,
            Raw: rawCookie,
            Unparsed: []string{rawCookie},
        }
        http.SetCookie(w, &cookie)
        view.SayWelcome(w)
    } else {
        view.SayWelcome(w)
    }

}

func signUp(w http.ResponseWriter, r *http.Request) {
    view.SignUp(w)
}

func submit(w http.ResponseWriter, r *http.Request) {
    
}

func main() {
    port := os.Getenv("PORT")
    if len(port) == 0 {
        port = "8000"
    }
    http.HandleFunc("/", hello)
    http.HandleFunc("/sign_up", signUp)
    http.HandleFunc("/submit", submit)
    http.ListenAndServe(":" + port, nil)
}
