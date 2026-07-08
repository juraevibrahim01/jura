package models

import (
	"errors"

	"github.com/golang-jwt/jwt/v5"
)

// ------------- login ---------------------
type Auth_Req_Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Auth_Res_Login struct {
	Status      string `json:"status"`
	Description string `json:"description"`
	Error       string `json:"error"`
	Step        string `json:"step"`
}

// ------------------------------------------

// ---------------------- otp ----------------
type Auth_req_otp struct {
	Email string `json:"email"`
	Code  string `json:"code"`
	Role  string `json:"role"`
}

type Auth_res_otp struct {
	Status       string `json:"status"`
	Description  string `json:"description"`
	Error        string `json:"error"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type Claims struct {
	Email  string `json:"email"`
	Role   string `json:"role"`
	UserID int    `json:"user_id"`
	jwt.RegisteredClaims
}

// -------------------------- Middleware --------------------------
type Auth_middleware_res struct {
	Status      string `json:"status"`
	Description string `json:"description"`
}

type Contextkey string

// -------------------------------------------------------------------

// --------------------- ref_token ----------------------------------

type Auth_req_ref_token struct {
	RefToken string `json:"ref_token"`
	Email    string `json:"email"`
}

type Auth_res_ref_token struct {
	Status   string `json:"status"`
	Error    string `json:"error"`
	Token    string `json:"token"`
	RefToken string `json:"ref_token"`
}

// -----------------------------------------------------------------------

// ------------------------ Errors -----------------------------------------
var ErrTokenInvalid = errors.New("Token is invalid")
var ErrTokenExpired = errors.New("Token is expired") // Истек
// ------------------------------------------------------------------------
