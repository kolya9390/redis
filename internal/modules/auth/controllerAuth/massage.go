package controllerauth

//go:generate easytags $GOFILE

type RegisterRequest struct {
	Email          string `json:"email"`
	Password       string `json:"password"`
	RetypePassword string `json:"retype_password"`
}

type RegisterResponse struct {
	Success   bool `json:"success"`
	ErrorCode int  `json:"error_code"`
	Data      Data `json:"data"`
}

type Data struct {
	Message string `json:"message"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RefreshRequest struct {
	UserID int `json:"user_id"`
}

type AuthResponse struct {
	Success   bool      `json:"success"`
	ErrorCode int       `json:"error_code"`
	Data      LoginData `json:"data"`
}

type LoginData struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	Message      string `json:"message"`
}
