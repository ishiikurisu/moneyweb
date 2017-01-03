package controller

import "net/http"
import "github.com/ishiikurisu/moneyweb/model"
import "github.com/ishiikurisu/moneyweb/view"

// The definitions of a server.
type Server struct {
    // This storage is responsible for dealing with Cookies.
    Storage *model.LocalStorage

    // This is the port this webserver will serve upon.
    Port string

    //TODO Create database access
}

// Create a webserver for this app.
func CreateServer() Server {
    // Create server structure
    storage, oops := model.NewLocalStorage()
    if oops != nil {
        panic(oops)
    }

    server := Server{
        Storage: storage,
        Port: model.GetPort(),
    }

    // Route stuff
    http.HandleFunc("/", server.SayHello)

    return server
}

// Puts the webserver to, well, serve.
func (server *Server) Serve() {
    http.ListenAndServe(server.Port, nil)
}

// SERVER ROUTED FUNCTIONS

// Function for index function
func (server *Server) SayHello(w http.ResponseWriter, r *http.Request) {
    user := server.Storage.GetLog(w, r)
    if len(user) > 0 {
        view.SayWelcome(w)
    } else {
        view.SayHello(w)
    }
}

// TODO Create serverless actions
