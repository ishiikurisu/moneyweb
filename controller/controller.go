package controller

import "net/http"
import "github.com/ishiikurisu/moneyweb/model"
import "github.com/ishiikurisu/moneyweb/view"

type Server struct {
    Storage *model.LocalStorage
    Port string
}

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
    http.HandleFunc("/sign_up", server.SignUp)
    http.HandleFunc("/register", server.Register)
    http.HandleFunc("/login", server.LogIn)

    return server
}

func (server *Server) Serve() {
    http.ListenAndServe(server.Port, nil)
}

// SERVER ROUTED FUNCTIONS

// Function for index function
func (server *Server) SayHello(w http.ResponseWriter, r *http.Request) {
    user := server.Storage.GetUser(w, r)
    if len(user) > 0 {
        view.SayWelcome(w)
    } else {
        view.SayHello(w)
    }
}

// Function for sign up page
func (server *Server) SignUp(w http.ResponseWriter, r *http.Request) {
    view.SignUp(w)
}

// Function for login screen
func (server *Server) LogIn(w http.ResponseWriter, r *http.Request) {
    view.Login(w)
}

func (server *Server) Register(w http.ResponseWriter, r *http.Request) {
    username := r.FormValue("username")
    password := r.FormValue("password")

    // TODO Use storage to store these cookies
    // TODO Create a database object to store these
    if model.RegisterUser(username, password) {
        w, r = server.Storage.AddCookie(w, r)
    }
    http.Redirect(w, r, "/", http.StatusFound)
}
