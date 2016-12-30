package model

import "net/http"
import "time"
import "fmt"
// import "net/http/cookiejar"

func GetUser(w http.ResponseWriter, r *http.Request) string {
    outlet := ""
    cookie, err := r.Cookie("UserName")

    if err == nil {
        outlet = cookie.Value
    } else {
        fmt.Println(err)
    }

    return outlet
}

func RegisterUser(username, password string) bool {
    // TODO Implement database
    return true
}

func AddCookie(w http.ResponseWriter, r *http.Request) (http.ResponseWriter, *http.Request) {
    username, _ := getUserAndPassword(r)
    cookie := http.Cookie{
        Name: "UserName",
        Value: username,
        Expires: time.Now().Add(86400),
        HttpOnly: false,
    }
    http.SetCookie(w, &cookie)
    return w, r
}

func getUserAndPassword(r *http.Request) (string, string) {
    username := r.FormValue("username")
    password := r.FormValue("password")
    return username, password
}
