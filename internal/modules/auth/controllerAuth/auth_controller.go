package controllerauth

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"studentgit.kata.academy/Nikolai/historysearch/internal/infrastructure/responder"
	passhesher "studentgit.kata.academy/Nikolai/historysearch/internal/infrastructure/tools/passHesher"
	"studentgit.kata.academy/Nikolai/historysearch/internal/infrastructure/tools/token"
)



type Auther interface {
	Register(http.ResponseWriter, *http.Request)
	Login(http.ResponseWriter, *http.Request)

}

type Auth struct {
	responder.Responder
}

var UserRegister = make(map[string]string)
var UserLoginData = make(map[string]LoginData)


func NewAuth(responder responder.Responder)*Auth{

	return &Auth{Responder: responder}
}

func(a *Auth) Login(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		a.Responder.ErrorBadRequest(w,err)
		return
	}

	if _, OK := UserRegister[req.Email]; !OK {
		a.Responder.ErrorUnauthorized(w,fmt.Errorf("an account with this email  has not been registered"))
	
	}

	if !passhesher.CheckPassword(UserRegister[req.Email], req.Password) {
		a.Responder.ErrorUnauthorized(w,fmt.Errorf("invalid password"))
	}

	_, AccessToken, err := token.TokenAuthorization.Encode(map[string]interface{}{req.Email: req.Password})
	if err != nil {
		a.Responder.ErrorInternal(w,err)
	}
	AccessToken = fmt.Sprintf("BEARER %s", AccessToken)


	a.OutputJSON(w,AuthResponse{
		Success:   true,
		ErrorCode: http.StatusOK,
		Data: LoginData{
			AccessToken: AccessToken,
			Message:     "you logged in successfully",
		}})

}

func(a *Auth) Register(w http.ResponseWriter, r *http.Request) {

	var req RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		a.Responder.ErrorBadRequest(w,err)
		return
	}

	hashPass, err := passhesher.HashPassword(req.Password)

	if err != nil {
		log.Println(err)
	}

	if _, OK := UserRegister[req.Email]; OK {
		http.Error(w, fmt.Sprintf("An account with this Email ( %s ) has already been registered", req.Email), http.StatusOK)
		return
	}

	UserRegister[req.Email] = hashPass
	w.WriteHeader(http.StatusOK)
	a.OutputJSON(w,RegisterResponse{
		Success: true,
		ErrorCode: http.StatusOK,
		Data: Data{
			Message: "Successfully Register",
		},
	
	})

}
