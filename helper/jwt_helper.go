package helper

import (
	"GolangEchoGorm/db"
	"GolangEchoGorm/model"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"net/http"
	"time"
)

type JwtCustomClaims struct {
	Username string `json:"username"`
	Admin    bool   `json:"admin"`
	jwt.StandardClaims
}

// CreateJwtToken create JWT token from username
func CreateJwtToken(user model.User) (interface{}, error) {
	currentDb, err := db.Database()
	if err != nil {
		return "DB_OPEN_ERROR", err
	}
	defer currentDb.Close()

	expireTime := time.Now().Add(time.Hour * 72).Unix()
	claims := &JwtCustomClaims{
		user.Username,
		user.Admin,
		jwt.StandardClaims{
			//ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
			ExpiresAt: expireTime,
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "can not create token", err
	}

	tokenDB := model.Token{
		Username:       user.Username,
		Token:          t,
		DateExpiration: expireTime,
	}
	currentDb.Save(&tokenDB)

	return echo.Map{
		"token":     t,
		"expiresAt": expireTime,
	}, nil
}

func ValidToken(token string) (bool, error) {

	// Get db
	currentDb, err := db.Database()
	if err != nil {
		return false, err
	}
	defer currentDb.Close()

	selectToken := model.Token{Token: token}
	currentDb.Where(&selectToken).First(&selectToken)

	if selectToken.ID < 1 {
		return false, errors.New("TOKEN INVALID")
	}

	return true, nil
}

// GetIdentifierFromJWT gets Identifier object from the Jwt claims
func GetIdentifierFromJWT(user interface{}) (Identifier, bool, error) {
	identifier := Identifier{}
	token, isParsed := user.(*jwt.Token)
	if !isParsed {
		return identifier, false, errors.New("CAN'T PARSE")
	}

	valid, err := ValidToken(token.Raw)

	if err != nil {
		return identifier, false, err
	}
	if !valid {
		return identifier, false, errors.New("IS NOT VALID")
	}

	claims := token.Claims.(*JwtCustomClaims)
	identifier.Username = claims.Username
	identifier.Admin = claims.Admin

	return identifier, token.Valid, nil
}

// RemoveToken gets Identifier object from the Jwt claims
func RemoveToken(user interface{}) (int, error) {
	currentDb, err := db.Database()
	if err != nil {
		return http.StatusInternalServerError, err
	}
	defer currentDb.Close()
	token, isParsed := user.(*jwt.Token)
	if !isParsed {
		return http.StatusInternalServerError, errors.New("CAN'T PARSE")
	}
	claims := token.Claims.(*JwtCustomClaims)
	tokenDB := model.Token{Username: claims.Username}

	dbResponse := currentDb.Where(&tokenDB).Delete(&tokenDB)
	if dbResponse.Error != nil {
		return http.StatusBadRequest, dbResponse.Error
	}
	return http.StatusBadRequest, nil
}
