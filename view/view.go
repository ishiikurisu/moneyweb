package view

import (
    "os"
    "io"
    "html/template"
    "fmt"
)

/* SERVER STUFF */

// Gets the current PWD
func GetPwd() string {
    codePath := "./src/github.com/ishiikurisu/logeyweb/"
    port := os.Getenv("PORT")

    if len(port) != 0 {
        codePath = os.Getenv("HOME") + "/"
    }

    return codePath
}

// Standard procedure to load a HTML file that does not need any customization.
func LoadFileWithoutArgs(writer io.Writer, path string) {
    htmlPath := GetPwd() + path
    templ, err := template.ParseFiles(htmlPath)
    viewModel := NewViewModel()
    err = templ.Execute(writer, viewModel)
    if err != nil {
        fmt.Printf("%#v\n", err)
    }
}

// Standard procedure to load a HTML file that needs some customization.
func LoadFileWithArgs(writer io.Writer, path string, args map[string]template.HTML) {
    htmlPath := GetPwd() + path
    templ, err := template.ParseFiles(htmlPath)
    viewModel := NewViewModel()
    viewModel.AddBody(args)
    err = templ.Execute(writer, viewModel)
    if err != nil {
        fmt.Printf("%#v\n", err)
    }
}

/* VIEWS */


func ListStaticFiles() string {
    return fmt.Sprintf("%sstatic", GetPwd())
}


// Displays the home screen
func SayHello(writer io.Writer) {
    LoadFileWithoutArgs(writer, "assets/html/index.gohtml")
}
