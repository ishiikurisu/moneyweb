package model

import "fmt"
import "strings"
import "net/http"
import "github.com/ishiikurisu/logey"

/**************************
 * CEMETERY OF PROCEDURES *
 **************************/

func RegisterUser(username, password string) bool {
    // TODO Implement database
    return true
}

func getUserAndPassword(r *http.Request) (string, string) {
    username := r.FormValue("username")
    password := r.FormValue("password")
    return username, password
}

/*****************************
 * MONEY LOG TRANSFORMATIONS *
 *****************************/

// Turns a raw money log string into a map of useful information
func LogToMap(raw string) map[string]string {
    outlet := make(map[string]string)
    log := logey.LogFromString(raw)

    numericValues := log.GetValues()
    literalValues := make([]string, len(numericValues))
    for i, value := range numericValues {
        literalValues[i] = fmt.Sprintf("%.2F", value)
    }

    outlet["values"] = strings.Join(literalValues, "\n")
    outlet["descriptions"] = strings.Join(log.GetDescriptions(), "\n")
    outlet["balance"] = fmt.Sprintf("%.2F", log.CalculateBalance())

    return outlet
}
