package session

import (
	"github.com/gorilla/sessions"
	"github.com/hem-nav-gateway/config"
	"net/http"
)

var Store = sessions.NewCookieStore(config.SessionKey)

var handler *sessions.Session

func SetSession(r *http.Request, companyName string) {
	handler, _ = Store.New(r, "main-session")
	handler.Values["company"] = companyName
}

func GetSession() *sessions.Session {
	return handler
}

func CompanyName() interface{} {
	s := GetSession()
	return s.Values["company"]
}
