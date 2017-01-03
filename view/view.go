package view

import "os"
import "io"
import "io/ioutil"
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

// Extracts the CSS path
func loadCss() string {
    pwd := GetPwd()
    css, err := ioutil.ReadFile(pwd + "assets/app.css")

    if err != nil {
        fmt.Println(err)
        css = []byte { }
    }

    return string(css)
}

// Standard procedure to load a HTML file that does not need any customization.
func LoadFileWithoutArgs(writer io.Writer, path string) {
    htmlPath := GetPwd() + path
    templ, err := template.ParseFiles(htmlPath)
    css := template.CSS(loadCss())
    err = templ.Execute(writer, css)
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
