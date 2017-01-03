package model

import "os"

// Gets the port for releasing the server
func GetPort() string {
    port := os.Getenv("PORT")

    if len(port) == 0 {
        port = "8000"
    }

    return ":" + port
}

// Gets the app web address
func GetAddress() string {
    address := "logey.herokuapp.com"
    port := os.Getenv("PORT")

    if len(port) == 0 {
        address = "localhost:8000"
    }

    return "http://" + address
}
