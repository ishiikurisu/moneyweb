package model

import "net/http"

func RegisterUser(username, password string) bool {
    // TODO Implement database
    return true
}

func getUserAndPassword(r *http.Request) (string, string) {
    username := r.FormValue("username")
    password := r.FormValue("password")
    return username, password
}
