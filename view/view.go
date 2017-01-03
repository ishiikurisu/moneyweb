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

// TODO Create loading with args, where args is a `map[string]string` to be
// added to the `ViewModel` class that formats the page.
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

// Displays the home screen
func SayHello(writer io.Writer) {
    LoadFileWithoutArgs(writer, "assets/html/index.empty.gohtml")
}

func SayWelcome(writer io.Writer) {
    LoadFileWithoutArgs(writer, "assets/html/index.filled.gohtml")
}

func AddEntry(writer io.Writer, body map[string]string) {
    args := make(map[string]template.HTML)
    input := fmt.Sprintf("<input type=\"text\" name=\"description\" value=\"%s\"/>", body["Message"])
    args["Query"] = template.HTML(input)
    LoadFileWithArgs(writer, "assets/html/add.gohtml", args)
}
