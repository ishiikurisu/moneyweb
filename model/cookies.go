package model

import "fmt"
import "net/url"
import "net/http"
import "net/http/cookiejar"
import "github.com/ishiikurisu/moneylog"

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
        }
        storage = &s
    } else {
        fmt.Println(shit)
    }

    return storage, oops
}

// Adds a cookie to the cookiejar and saves it on the user session.
// Cookies added here will last forever because of the serverless nature
// of this app.
func (storage *LocalStorage) AddCookie(w http.ResponseWriter, r *http.Request) (http.ResponseWriter, *http.Request) {
    username, _ := getUserAndPassword(r)
    cookie := http.Cookie {
        Name: "UserName",
        Value: username,
    }
    storage.CookieJar.SetCookies(storage.Url, append(storage.CookieJar.Cookies(storage.Url), &cookie))
    http.SetCookie(w, &cookie)
    return w, r
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

// Extracts the current user based on the store cookies. If there is no user,
// a string of length 0 is returned.
func (storage *LocalStorage) GetUser(w http.ResponseWriter, r *http.Request) string {
    outlet := ""
    cookie, err := r.Cookie("UserName")

    if err == nil {
        outlet = cookie.Value
    } else {
        fmt.Println(err)
    }

    return outlet
}

// Adds the given raw entry to the log.
func (storage *LocalStorage) AddEntryFromRaw(rawDescription, rawValue string) {
    var description string
    var value float64

    description = rawDescription
    fmt.Sscanf(rawValue, "%F", &value)

    storage.MoneyLog.Add(description, value)
}

// Saves the current money log on a cookie.
func (storage *LocalStorage) SaveLog(w http.ResponseWriter, r *http.Request) (http.ResponseWriter, *http.Request) {
    rawLog := storage.MoneyLog.ToString()
    cookie := http.Cookie {
        Name: "MoneyLog",
        Value: rawLog,
    }
    storage.CookieJar.SetCookies(storage.Url, append(storage.CookieJar.Cookies(storage.Url), &cookie))
    http.SetCookie(w, &cookie)
    return w, r
}
