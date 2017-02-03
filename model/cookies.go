package model

import "fmt"
import "net/url"
import "net/http"
import "net/http/cookiejar"
import "github.com/ishiikurisu/moneylog"
import "mime/multipart"
import "bufio"
import "time"
import "math"

/*****************************
 * LOCAL STORAGE DEFINITIONS *
 *****************************/

// This is the structure that will deal with the session's cookies. It will
// need to know which URL these cookies relate to, and its logic is already
// coded to understand the app's behaviour.
type LocalStorage struct {
    // The actual data. The logic of this application. The reason we are here.
    MoneyLog moneylog.Log

    // The structure that will deal with our cookies.
    CookieJar *cookiejar.Jar

    // This is the URL these cookies refer to.
    Url *url.URL

    // This is where the current log is stored on local memory
    LogFile string
}

// Creates a cookie storage structure.
func NewLocalStorage() (*LocalStorage, error) {
    var storage *LocalStorage = nil
    jar, oops := cookiejar.New(nil)
    url, shit := url.Parse(GetAddress())

    if oops == nil && shit == nil {
        url.Scheme = "http"
        url.Host = "heroku.com"
        s := LocalStorage {
            // TODO Implement actual cookie jar
            MoneyLog: moneylog.EmptyLog(),
            CookieJar: jar,
            Url: url,
            LogFile: "log.txt",
        }
        storage = &s
    } else {
        fmt.Println(shit)
    }

    return storage, oops
}

// Extracts the current money log based on the store cookies. If there is no
// log, a string of length 0 is returned.
func (storage *LocalStorage) GetLog(w http.ResponseWriter, r *http.Request) string {
    outlet := ""
    cookie, err := r.Cookie("MoneyLog")

    if err == nil {
        outlet = cookie.Value
    } else {
        fmt.Println(err)
    }

    return outlet
}

// Extracts stuff from raw strings
func stuffFromRaw(rawDescription, rawValue string) (string, float64) {
    var description string
    var value float64

    description = rawDescription
    fmt.Sscanf(rawValue, "%F", &value)

    return description, value
}

// Converts a float64 to a string
func stuff2raw(x float64) string {
    y := fmt.Sprintf("%.2F", x)
    return y
}

// Extracts the log from a multipart file included in a HTTP request.
func (storage *LocalStorage) AddLogFromFile(mmf multipart.File) string {
    buffer := bufio.NewReader(mmf)
    current := ReadField(buffer)
    outlet := ""

    for current != "...," {
        outlet += current
        outlet += ","
        current = ReadField(buffer)
    }

    outlet += current
    return outlet
}

// Reads a field as specified on the money log API from the reader buffer.
func ReadField(reader *bufio.Reader) string {
    raw := make([]byte, 0)
    raw, err := reader.ReadBytes(',')
    if err != nil {
        panic(err)
    }
    return string(raw)
}

// Saves a log on memory in the form of a cookie
func (storage *LocalStorage) SaveCookie(log moneylog.Log, w http.ResponseWriter) http.ResponseWriter {
    cookie := http.Cookie {
        Name: "MoneyLog",
        Value: log.ToString(),
        Expires: time.Date(2020, time.May, 25, 23, 0, 0, 0, time.UTC),
        MaxAge: math.MaxInt32,
        HttpOnly: false,
    }
    http.SetCookie(w, &cookie)
    return w
}

// Adds a entry to current log and store the changes onto a cookie.
func (storage *LocalStorage) AddEntryFromRawAndSaveLog(d, v string, w http.ResponseWriter, r *http.Request) (http.ResponseWriter, *http.Request) {
    description, value := stuffFromRaw(d, v)
    raw := storage.GetLog(w, r)
    log := moneylog.LogFromString(raw)
    log.Add(description, value)
    w = storage.SaveCookie(log, w)
    return w, r
}

// Stores the given log onto a cookie.
func (storage *LocalStorage) AddLogFromRawAndSaveLog(raw string, w http.ResponseWriter, r *http.Request) (http.ResponseWriter, *http.Request) {
    log := moneylog.LogFromString(raw)
    w = storage.SaveCookie(log, w)
    return w, r
}
