package model

import "net/http"
import "github.com/ishiikurisu/moneylog"
import "io"

func RegisterApi() {
    http.HandleFunc("/api/new", func (w http.ResponseWriter, r *http.Request) {
        log := moneylog.EmptyLog()
        io.WriteString(w, log.ToString())
    })
    http.HandleFunc("/api/add", func (w http.ResponseWriter, r *http.Request) {
        rawDescription := r.FormValue("description")
        rawValue := r.FormValue("value")
        rawLog := r.FormValue("log")
        log := moneylog.LogFromString(rawLog)
        description, value := stuffFromRaw(rawDescription, rawValue)
        log.Add(description, value)
        io.WriteString(w, log.ToString())
    })
    http.HandleFunc("/api/get/balance", func (w http.ResponseWriter, r *http.Request) {
        rawLog := r.FormValue("log")
        log := moneylog.LogFromString(rawLog)
        balance := log.CalculateBalance()
        io.WriteString(w, stuff2raw(balance))
    })
    http.HandleFunc("/api/get/descriptions", func (w http.ResponseWriter, r *http.Request) {
        rawLog := r.FormValue("log")
        log := moneylog.LogFromString(rawLog)
        descriptions := log.GetDescriptions()
        outlet := ""
        for _, description := range descriptions {
            outlet += description + ","
        }
        io.WriteString(w, outlet)
    })
    http.HandleFunc("/api/get/values", func (w http.ResponseWriter, r *http.Request) {
        rawLog := r.FormValue("log")
        log := moneylog.LogFromString(rawLog)
        values := log.GetValues()
        outlet := ""
        for _, value := range values {
            outlet += stuff2raw(value) + ","
        }
        io.WriteString(w, outlet)
    })
}
