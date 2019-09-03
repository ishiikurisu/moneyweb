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
    http.Handle("/", server.ProvideStaticFiles())
    // TODO add login and sign up routes
    // TODO add API for common actions
    http.HandleFunc("/api/hi", server.SayHi)

    return server
}

// Puts the webserver to, well, serve.
func (server *Server) Serve() {
    http.ListenAndServe(server.Port, nil)
}

/***************************
 * SERVER ROUTED FUNCTIONS *
 ***************************/

func (server *Server) ProvideStaticFiles() http.Handler {
    return http.FileServer(http.Dir(view.ListStaticFiles()))
}

// api test
func (server *Server) SayHi(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("hello!"))
}
