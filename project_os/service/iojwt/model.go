package iojwt

import (
	"github.com/dgrijalva/jwt-go"
)

type roleLogin struct {
	LoginUuid string `json:"login_uuid"`
	UserName  string `json:"user_name"`
	PassWord  string `json:"pass_word"`
	Name      string `json:"name"`
	Email     string `json:"email"`
}

type credentials struct {
	UserName string `json:"user_name"`
	PassWord string `json:"pass_word"`
}

type claims struct {
	LoginUuid string `json:"login_uuid"`
	UserName  string `json:"user_name"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	jwt.StandardClaims
}

type DecodeClaims struct {
	LoginUuid string `json:"login_uuid"`
	UserName  string `json:"user_name"`
	Name      string `json:"name"`
	Email     string `json:"email"`
}

type parseCode struct {
	Value string `json:"value"`
}

type messageResponse struct {
	Status             int    `json:"status"`
	MessageDescription string `json:"message_description"`
}

type inputRegister struct {
	LoginUuid string `json:"login_uuid"`
	UserName  string `json:"user_name"`
	PassWord  string `json:"pass_word"`
	Name      string `json:"name"`
	Email     string `json:"email"`
}