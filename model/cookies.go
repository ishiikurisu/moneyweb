package model

import "net/http/cookiejar"

type ServerStorage struct {
    CookieJar *cookiejar.Jar
}

func NewServerStorage() (*ServerStorage, error) {
    var storage *ServerStorage = nil
    jar, oops := cookiejar.New(nil)

    if oops == nil {
        s := ServerStorage{
            // TODO Implement actual cookie jar
            CookieJar: jar,
        }
        storage = &s
    }

    return storage, oops
}
