package view

import "os"
import "io"
import "html/template"
import "fmt"
import "strings"

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

// Displays the home screen
func SayHello(writer io.Writer) {
    LoadFileWithoutArgs(writer, "assets/html/index.empty.gohtml")
}

// Displays a money log
func BeUseful(writer io.Writer, body map[string]string) {
    args := make(map[string]template.HTML)
    descriptions := strings.Split(body["descriptions"], "\n")
    values := strings.Split(body["values"], "\n")
    limit := len(descriptions)

    entries := ""
    for i := limit-1; i >= 0; i-- {
        // TODO Make these entries prettier
        entries = fmt.Sprintf("%s<p>%s: %s$</p>\n", entries, descriptions[i], values[i])
    }
    args["Entries"] = template.HTML(entries)
    args["Balance"] = template.HTML(fmt.Sprintf("%s$", body["balance"]))

    LoadFileWithArgs(writer, "assets/html/index.filled.gohtml", args)
}

// Displays page to add entry
func AddEntry(writer io.Writer, body map[string]string) {
    // TODO Change this body map argument to the message string
    args := make(map[string]template.HTML)
    input := fmt.Sprintf("<input type=\"text\" name=\"description\" value=\"%s\"/>", body["Message"])
    args["Query"] = template.HTML(input)
    LoadFileWithArgs(writer, "assets/html/add.gohtml", args)
}
