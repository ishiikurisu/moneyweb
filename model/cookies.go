package model

import "fmt"
import "net/url"
import "net/http"
import "net/http/cookiejar"

type LocalStorage struct {
    CookieJar *cookiejar.Jar
    Url *url.URL
}

func NewLocalStorage() (*LocalStorage, error) {
    var storage *LocalStorage = nil
    jar, oops := cookiejar.New(nil)
    url, shit := url.Parse(GetAddress())

    if oops == nil && shit == nil {
        url.Scheme = "http"
        url.Host = "heroku.com"
        s := LocalStorage{
            // TODO Implement actual cookie jar
            CookieJar: jar,
            Url: url,
        }
        storage = &s
    } else {
        fmt.Println(shit)
    }

    return storage, oops
}

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
