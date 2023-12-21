package token

import (
//	"fmt"
//	"github.com/golang-jwt/jwt/v4"
	"time"
	"github.com/go-chi/jwtauth/v5"
	"github.com/lestrrat-go/jwx/v2/jwt"
)

var TokenAuthorization = jwtauth.New("HS256", []byte("secret"), nil, jwt.WithAcceptableSkew(30*time.Second))



/*
const (
	AccessToken = iota
	RefreshToken
)

type TokenManager interface {
	CreateToken(userID, role, groups string, ttl time.Duration, kind int) (string, error)
	ParseToken(inputToken string, kind int) (UserClaims, error)
}

type TokenJWT struct {
	AccessSecret  []byte
	RefreshSecret []byte
}

func NewTokenJWT() TokenManager {
	return &TokenJWT{AccessSecret: []byte("secret"), RefreshSecret: []byte("secret")}
}

// UserClaims include custom claims on jwt.
type UserClaims struct {
	ID     string `json:"uid"`
	Role   string `json:"role"`
	Groups string `json:"groups"`
	Layers string `json:"layers"`
	jwt.RegisteredClaims
}

type UserFromClaims struct {
	ID     int
	Role   int
	Groups []int
	Layers []int
}




// CreateToken create new token with parameters.
func (o *TokenJWT) CreateToken(userID, role, groups string, ttl time.Duration, kind int) (string, error) {
	claims := UserClaims{
		ID:               userID,
		Role:             role,
		Groups:           groups,
		RegisteredClaims: jwt.RegisteredClaims{ExpiresAt: jwt.NewNumericDate(time.Now().Add(ttl))},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	var secret []byte
	switch kind {
	case AccessToken:
		secret = o.AccessSecret
	case RefreshToken:
		secret = o.RefreshSecret
	default:
		return "", fmt.Errorf("error secret token")
	}

	return token.SignedString(secret)
}

// ParseToken parsing input token, and return email and role from token.
func (o *TokenJWT) ParseToken(inputToken string, kind int) (UserClaims, error) {
	token, err := jwt.Parse(inputToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		var secret []byte
		switch kind {
		case AccessToken:
			secret = o.AccessSecret
		case RefreshToken:
			secret = o.RefreshSecret
		default:
			return "", fmt.Errorf("error secret token")
		}
		_ = secret

		return secret, nil
	})

	if err != nil {
		return UserClaims{}, err
	}

	if !token.Valid {
		return UserClaims{}, fmt.Errorf("not valid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return UserClaims{}, fmt.Errorf("error get user claims from token")
	}

	return UserClaims{
		ID:               claims["uid"].(string),
		Role:             claims["role"].(string),
		Groups:           claims["groups"].(string),
		RegisteredClaims: jwt.RegisteredClaims{},
	}, nil
}
*/