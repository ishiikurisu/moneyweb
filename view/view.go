package view

import "os"
import "io"
import "html/template"
import "fmt"
import "strings"

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

    entries := "<table rules='all'>\n<tr><th>Description</th><th>Value</th></tr>"
    for i := limit-1; i >= 0; i-- {
        tdClass := "positive"
        if !isPositive(values[i]) {
            tdClass = "negative"
        }
        entries = fmt.Sprintf("%s<tr><td>%s</td>", entries, descriptions[i])
        entries = fmt.Sprintf("%s<td class='%s'>%s$</td></p>\n", entries, tdClass, values[i])
    }
    entries = fmt.Sprintf("%s</table>\n", entries)
    args["Entries"] = template.HTML(entries)
    // TODO Make this balance prettier
    args["Balance"] = template.HTML(fmt.Sprintf("%s$", body["balance"]))

    LoadFileWithArgs(writer, "assets/html/index.filled.gohtml", args)
}

func isPositive(raw string) bool {
    var value float64
    fmt.Sscanf(raw, "%f", &value)
    return value >= 0
}

// Displays page to add entry
func AddEntry(writer io.Writer, body map[string]string) {
    // TODO Change this body map argument to the message string
    args := make(map[string]template.HTML)
    input := fmt.Sprintf("<input type=\"text\" name=\"description\" value=\"%s\"/>", body["Message"])
    args["Query"] = template.HTML(input)
    LoadFileWithArgs(writer, "assets/html/add.gohtml", args)
}

// Writes raw file to browser
func EnableData(writer io.Writer, data string) {
    io.WriteString(writer, data)
}

// Displays page to enable upload of previous user log
func UploadLog(writer io.Writer) {
    LoadFileWithoutArgs(writer, "assets/html/upload.gohtml")
}
