package controller

import "net/http"
import "github.com/ishiikurisu/moneyweb/model"
import "github.com/ishiikurisu/moneyweb/view"

/*********************
 * SERVER DEFINITION *
 *********************/

// The definitions of a server.
type Server struct {
    // This storage is responsible for dealing with Cookies.
    Storage *model.LocalStorage

    // This is the port this webserver will serve upon.
    Port string
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
    http.HandleFunc("/add", server.AddEntry)
    http.HandleFunc("/register", server.Register)
    http.HandleFunc("/raw", server.DownloadLog)
    http.HandleFunc("/upload", server.UploadLog)
    http.HandleFunc("/uploading", server.UploadingLog)

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
    log := server.Storage.GetLog(w, r)
    if len(log) > 0 {
        view.BeUseful(w, model.LogToMap(log))
    } else {
        view.SayHello(w)
    }
}

// Function for adding a entry
func (server *Server) AddEntry(w http.ResponseWriter, r *http.Request) {
    rawLog := server.Storage.GetLog(w, r)
    msg := model.GetRandomMessage()

    if len(rawLog) == 0 {
        msg = "First log!"
    }

    data := make(map[string]string)
    data["Message"] = msg
    view.AddEntry(w, data)
}

// Saves entry to cookies and to model format
func (server *Server) Register(w http.ResponseWriter, r *http.Request) {
    // Extract data from request
    description := r.FormValue("description")
    value := r.FormValue("value")

    // Save entry to current log
    w, r = server.Storage.AddEntryFromRawAndSaveLog(description, value, w, r)

    // Redirecting to correct page
    http.Redirect(w, r, "/", http.StatusFound)
}

// Displays raw log on browser
func (server *Server) DownloadLog(w http.ResponseWriter, r *http.Request) {
    view.EnableData(w, server.Storage.GetLog(w, r))
}

// Displays the page to download log
func (server *Server) UploadLog(w http.ResponseWriter, r *http.Request) {
    view.UploadLog(w)
}

// Loads file from user's computer to use on logging
func (server *Server) UploadingLog(w http.ResponseWriter, r *http.Request) {
    // Extract data from request
    rawLog, _, err := r.FormFile("file")
    if err != nil {
        panic(err)
    }

    // Save entry to current log
    server.Storage.AddLogFromFile(rawLog)
    w, r = server.Storage.SaveLog(w, r)

    // Redirecting to correct page
    http.Redirect(w, r, "/", http.StatusFound)
}
