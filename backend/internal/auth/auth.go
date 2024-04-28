package auth

import (
	"encoding/json"
	"github.com/wlcmtunknwndth/hackBPA/internal/lib/httpResponse"
	"github.com/wlcmtunknwndth/hackBPA/internal/lib/slogResponse"
	"log/slog"
	"net/http"
	"time"
)

const (
	badRequest          = "Bad request"
	internalServerError = "Internal server error"
	unauthorized        = "Unauthorized"
	authorized          = "Authorized"
	noEnoughPrivileges  = "Not enough privileges"
)

type User struct {
	isAdmin  bool
	Username string `json:"username"`
	Password string `json:"password"`
}

//go:generate mockery --name Storage

type Storage interface {
	GetPassword(string) (string, error)
	RegisterUser(*User) error
	IsAdmin(string) bool
	DeleteUser(string) error
}

type Auth struct {
	Db Storage
}

func (a *Auth) Register(w http.ResponseWriter, r *http.Request) {
	const op = "auth.auth.Register"
	var usr User

	err := json.NewDecoder(r.Body).Decode(&usr)
	if err != nil {
		slog.Error("couldn't decode request: ", slogResponse.SlogOp(op), slogResponse.SlogErr(err))
		httpResponse.Write(w, http.StatusBadRequest, badRequest)
		return
	}

	if err = a.Db.RegisterUser(&usr); err != nil {
		slog.Error("could't register user: ", slogResponse.SlogOp(op), slogResponse.SlogErr(err))
		httpResponse.Write(w, http.StatusInternalServerError, internalServerError)
	}

	WriteNewToken(w, usr)

	return
}

func (a *Auth) LogIn(w http.ResponseWriter, r *http.Request) {
	const op = "auth.auth.Login"
	var usr User

	err := json.NewDecoder(r.Body).Decode(&usr)
	if err != nil {
		httpResponse.Write(w, http.StatusBadRequest, badRequest)
		slog.Error("couldn't decode user: ", slogResponse.SlogOp(op), slogResponse.SlogErr(err))
		return
	}

	if len(usr.Password) < 4 || len(usr.Username) < 4 {
		httpResponse.Write(w, http.StatusUnauthorized, unauthorized)
		return
	}

	pass, err := a.Db.GetPassword(usr.Username)
	if err != nil || pass != usr.Password {
		slog.Error("couldn't get password from storage: ", slogResponse.SlogOp(op), slogResponse.SlogErr(err))
		httpResponse.Write(w, http.StatusUnauthorized, unauthorized)
		return
	}

	usr.isAdmin = a.Db.IsAdmin(usr.Username)

	WriteNewToken(w, usr)

	return
}

func (a *Auth) LogOut(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:    access,
		Expires: time.Now(),
	})

	return
}

func (a *Auth) DeleteUser(w http.ResponseWriter, r *http.Request) {
	const op = "auth.auth.DeleteUser"

	if ok, err := IsAdmin(r); !ok {
		if err != nil {
			httpResponse.Write(w, http.StatusInternalServerError, internalServerError)
			slog.Error("couldn't determine if user is admin: ", slogResponse.SlogOp(op), slogResponse.SlogErr(err))
			return
		}
		httpResponse.Write(w, http.StatusForbidden, noEnoughPrivileges)
		return
	}

	type query struct {
		Username string `json:"username"`
	}

	var qry query
	err := json.NewDecoder(r.Body).Decode(&qry)
	if err != nil {
		httpResponse.Write(w, http.StatusBadRequest, badRequest)
		slog.Error("couldn't decode user: ", slogResponse.SlogOp(op), slogResponse.SlogErr(err))
		return
	}

	err = a.Db.DeleteUser(qry.Username)
	if err != nil {
		httpResponse.Write(w, http.StatusInternalServerError, internalServerError)
		slog.Error("couldn't delete user: ", slogResponse.SlogOp(op), slogResponse.SlogErr(err))
		return
	}

	return
}
