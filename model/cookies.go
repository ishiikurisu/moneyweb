package model

import "fmt"
import "net/url"
import "net/http"
import "net/http/cookiejar"
import "github.com/ishiikurisu/moneylog"
import "os"
import "mime/multipart"
import "bufio"

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

// Extracts stuff from raw strings
func stuffFromRaw(rawDescription, rawValue string) (string, float64) {
    var description string
    var value float64

    description = rawDescription
    fmt.Sscanf(rawValue, "%F", &value)

    return description, value
}

// Adds the given raw entry to the log.
func (storage *LocalStorage) AddEntryFromRaw(rawDescription, rawValue string) {
    description, value := stuffFromRaw(rawDescription, rawValue)
    storage.MoneyLog.Add(description, value)
}

// Saves the current money log on a cookie.
func (storage *LocalStorage) SaveLog(w http.ResponseWriter, r *http.Request) (http.ResponseWriter, *http.Request) {
    rawLog := storage.MoneyLog.ToString()
    cookie := http.Cookie {
        Name: "MoneyLog",
        Value: rawLog,
    }
    http.SetCookie(w, &cookie)
    return w, r
}


func (storage *LocalStorage) StoreLog(w http.ResponseWriter, r *http.Request) (http.ResponseWriter, *http.Request) {
    rawLog := storage.MoneyLog.ToString()
    fp, err := os.Create(storage.LogFile)

    if err != nil {
        panic(err)
    }
    defer fp.Close()

    _, err = fp.WriteString(rawLog)
    if err != nil {
        panic(err)
    }
    fp.Sync()

    return w, r
}

func (storage *LocalStorage) AddLogFromFile(mmf multipart.File) {
    buffer := bufio.NewReader(mmf)
    current := ReadField(buffer)
    outlet := ""

    for current != "...," {
        outlet += current
        outlet += ","
        current = ReadField(buffer)
    }

    outlet += current
    storage.MoneyLog = moneylog.LogFromString(outlet)
}

func ReadField(reader *bufio.Reader) string {
    raw := make([]byte, 0)
    raw, err := reader.ReadBytes(',')
    if err != nil {
        panic(err)
    }
    return string(raw)
}

func (storage *LocalStorage) AddEntryFromRawAndSaveLog(d, v string, w http.ResponseWriter, r *http.Request) (http.ResponseWriter, *http.Request) {
    description, value := stuffFromRaw(d, v)
    raw := storage.GetLog(w, r)
    log := moneylog.LogFromString(raw)
    log.Add(description, value)
    cookie := http.Cookie {
        Name: "MoneyLog",
        Value: log.ToString(),
    }
    http.SetCookie(w, &cookie)
    return w, r
}
