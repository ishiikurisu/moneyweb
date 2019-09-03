package controller

import (
    "net/http"
    "github.com/ishiikurisu/logeyweb/view"
)

/*********************
 * SERVER DEFINITION *
 *********************/

// The definitions of a server.
type Server struct {
    // This is the port this webserver will serve upon.
    Port string
}

// Create a webserver for this app.
func CreateServer() Server {
    // Create server structure
    server := Server {
        Port: GetPort(),
    }

    // Route stuff
    http.HandleFunc("/", server.SayHello)
    // TODO add login and sign up routes
    // TODO add API for common actions

    return server
}

// Puts the webserver to, well, serve.
func (server *Server) Serve() {
    http.ListenAndServe(server.Port, nil)
}

/***************************
 * SERVER ROUTED FUNCTIONS *
 ***************************/

// Function for index function
func (server *Server) SayHello(w http.ResponseWriter, r *http.Request) {
    view.SayHello(w)
}
