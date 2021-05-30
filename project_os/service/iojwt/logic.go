package iojwt

import (
	"github.com/google/uuid"
	_ "github.com/jinzhu/gorm"
	"net/http"

	"github.com/dgrijalva/jwt-go"

	"time"

	log "github.com/sirupsen/logrus"
)


func EnCodeHS(request credentials) (result parseCode, msgEncode messageResponse, err error) {
	var jwtKey = []byte("my_secret_key")

	var detail roleLogin
	detail, err = getLogin(request)
	if err != nil {
		msgEncode = messageResponse{
			Status:             http.StatusBadRequest,
			MessageDescription: "notFound"}
		return
	}
	log.Infof("Detail : %+v", detail)

	// expect Password && Username
	if detail.PassWord != request.PassWord || request.UserName != detail.UserName {
		msgEncode = messageResponse{
			Status:             http.StatusUnauthorized,
			MessageDescription: "unauthorized"}
		//return err
		return
	}

	expirationTime := time.Now().Add(time.Duration(12) * time.Hour)

	claims := &claims{
		LoginUuid: detail.LoginUuid,
		UserName:  detail.UserName,
		Name:detail.Name,
		Email:detail.Email,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		msgEncode = messageResponse{
			Status:             http.StatusInternalServerError,
			MessageDescription: "InternalServerError"}
		// If there is an error in creating the JWT return an internal server error
		return
	}

	result = parseCode{
		Value: tokenString,
	}

	return
}

func DeCodeHS(token []string) (result DecodeClaims, msgDecode messageResponse, err error) {
	claims := &claims{}

	tkn, err := jwt.ParseWithClaims(token[0], claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("my_secret_key"), nil
	})

	if !tkn.Valid {
		msgDecode = messageResponse{
			Status:             http.StatusUnauthorized,
			MessageDescription: "unauthorized"}
		return
	}

	result = DecodeClaims{
		LoginUuid: claims.LoginUuid,
		UserName:  claims.UserName,
		Name: claims.Name,
		Email:claims.Email,
	}

	return
}

func CreateUser(request inputRegister) (result parseCode, msgRegister messageResponse, err error) {
	log.Info("[registerUser]")
	request.LoginUuid = uuid.New().String()

	_, err = createUser(request)
	if err != nil  {
		log.Error("not found username in system")
		msgRegister = messageResponse{
			Status:             http.StatusBadRequest,
			MessageDescription: "not found username in system"}

		return
	}

	result = parseCode{
		Value: "create user success",
	}
	return
}

