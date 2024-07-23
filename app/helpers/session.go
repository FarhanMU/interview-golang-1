package helpers

import (
	"net/http"
	"time"
)

const (
	UserSession  = "user_session"
	RefreshToken = "refresh_token"
)

func SetCookie(w http.ResponseWriter, r *http.Request, sessionType string, value string) {
	cookie := http.Cookie{
		Name:     sessionType,
		Value:    value,
		Expires:  time.Now().Add(24 * time.Hour),
		Path:     "/",
		HttpOnly: true,
	}

	http.SetCookie(w, &cookie)
}

func GetCookie(w http.ResponseWriter, r *http.Request, sessionType string) (*http.Cookie, error) {
	cookie, err := r.Cookie(sessionType)
	if err != nil {
		if err == http.ErrNoCookie {
			w.Write([]byte("no cookie found"))
			return nil, err
		}
		w.WriteHeader(http.StatusBadRequest)
		return nil, err
	}
	return cookie, nil
}

func DeleteCookieHandler(w http.ResponseWriter, r *http.Request, sessionType string) {
	cookie := http.Cookie{
		Name:     sessionType,
		Value:    "",
		Path:     "/",
		Expires:  time.Unix(0, 0),
		MaxAge:   -1,
		HttpOnly: true,
	}
	http.SetCookie(w, &cookie)

	w.Write([]byte("cookie has been deleted"))
}

func SetUserCookie(w http.ResponseWriter, r *http.Request, user string) {
	SetCookie(w, r, UserSession, user)
}

func GetUserCookie(w http.ResponseWriter, r *http.Request) (*http.Cookie, error) {
	return GetCookie(w, r, UserSession)
}

func DeleteUserCookieHandler(w http.ResponseWriter, r *http.Request) {
	DeleteCookieHandler(w, r, UserSession)
}
