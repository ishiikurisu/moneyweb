package view

import "os"
import "io"
import "html/template"
import "fmt"

/* SERVER STUFF */

// Gets the current PWD
func GetPwd() string {
    codePath := "./src/github.com/ishiikurisu/moneyweb/"
    port := os.Getenv("PORT")

    if len(port) != 0 {
        codePath = os.Getenv("HOME") + "/"
    }

    return codePath
}

func LoadFileWithoutArgs(writer io.Writer, path string) {
    htmlPath := GetPwd() + path
    templ, err := template.ParseFiles(htmlPath)
    err = templ.Execute(writer, nil)
    if err != nil {
        fmt.Printf("%#v\n", err)
    }
}

/* VIEWS */

// Displays the home screen
func SayHello(writer io.Writer) {
    LoadFileWithoutArgs(writer, "viewmodel/index.empty.gohtml")
}

// Displays the logged screen
func SayWelcome(writer io.Writer) {
    LoadFileWithoutArgs(writer, "viewmodel/index.logged.gohtml")
}

// Loads sign up screen
func SignUp(writer io.Writer) {
    LoadFileWithoutArgs(writer, "viewmodel/sign_up.gohtml")
}

func Login(writer io.Writer) {
    LoadFileWithoutArgs(writer, "viewmodel/login.gohtml")
}
