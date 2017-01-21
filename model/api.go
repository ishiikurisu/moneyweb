package model

import "net/http"
import "github.com/ishiikurisu/moneylog"
import "io"

func RegisterApi() {
    http.HandleFunc("/api/new", func (w http.ResponseWriter, r *http.Request) {
        log := moneylog.EmptyLog()
        io.WriteString(w, log.ToString())
    })
}
