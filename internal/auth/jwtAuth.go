package auth

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"os"
	"time"
)

type Info struct {
	Username string `json:"username"`
	IsAdmin  bool   `json:"isAdmin"`
	jwt.RegisteredClaims
}

const (
	access             = "access"
	ttlToken           = 4 * time.Minute
	statusUnauthorized = "Unauthorized"
	statusBadRequest   = "Bad request"
	authKeyName        = "auth_key"
)

func checkRequest(r *http.Request) (*Info, error) {
	const op = "auth.jwtAuth.checkRequest"
	cookie, err := r.Cookie(access)
	if err != nil {
		if errors.Is(err, http.ErrNoCookie) {
			return nil, fmt.Errorf("%s: No cookie: %w", op, err)
		}
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	var info Info

	token, err := jwt.ParseWithClaims(cookie.Value, &info, func(token *jwt.Token) (any, error) {
		key, ok := os.LookupEnv(authKeyName)
		if !ok {
			return nil, fmt.Errorf("no secret key found")
		}
		return []byte(key), nil
	})
	if err != nil {
		if errors.Is(err, jwt.ErrSignatureInvalid) {
			return nil, fmt.Errorf("%s: Invalid jwt signature: %w", op, err)
		}
		return nil, fmt.Errorf("%s:%w", op, err)
	}

	if !token.Valid {
		return nil, fmt.Errorf("%s: Invalid token", op)
	}

	return &info, err
}

func Access(r *http.Request) (*Info, error) {
	const op = "auth.jwtAuth.Access"
	info, err := checkRequest(r)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return info, err
}

func Refresh(w http.ResponseWriter, r *http.Request) error {
	const op = "auth.jwtAuth.Refresh"

	info, err := checkRequest(r)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	expiresAt := time.Now().Add(ttlToken)
	info.ExpiresAt = jwt.NewNumericDate(expiresAt)

	//token := jwt.NewWithClaims(jwt.SigningMethodHS512, info)
	key, ok := os.LookupEnv(authKeyName)
	if !ok {
		return fmt.Errorf("%s: %w", op, err)
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS512, info).SignedString([]byte(key))
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	http.SetCookie(w, &http.Cookie{
		Name:    access,
		Value:   token,
		Expires: expiresAt,
	})
	return nil
}
