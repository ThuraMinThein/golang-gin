package helper

import (
	"os"
	"time"

	"github.com/ThuraMinThein/golang-gin/config/enums"
	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte(os.Getenv("JWT_SECRET_KEY"))

func GetTokens(userId uint, role *enums.UserRoleEnum) (string, string, error) {
	accessToken, err := getAccessToken(userId, role)
	if err != nil {
		return "","", err
	}
	refreshToken, err := getRefreshToken(userId)
	if err != nil {
		return "", "", err
	}
	return accessToken, refreshToken, nil
}

func getAccessToken(userId uint, role *enums.UserRoleEnum) (string, error) {

	claims := jwt.MapClaims{
			"exp": jwt.NewNumericDate(time.Now().Add(15 * time.Minute)),
			"userId": userId,
			"role": role,
		}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(secretKey)
	return ss, err
}

func getRefreshToken(userId uint) (string, error) {

	claims := jwt.MapClaims{
			"exp": jwt.NewNumericDate(time.Now().Add(168 * time.Hour)),
			"userId": userId,
		}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(secretKey)
	return ss, err
}

func VerifyToken(tokenString string) error {
   token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
      return secretKey, nil
   })
  
   if err != nil {
      return err
   }
  
   if !token.Valid {
      return err
   }
  
   return nil
}