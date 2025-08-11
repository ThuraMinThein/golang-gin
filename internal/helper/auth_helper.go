package helper

import (
	"os"
	"time"

	"github.com/ThuraMinThein/golang-gin/config/enums"
	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte(os.Getenv("JWT_SECRET_KEY"))

type UserClaims struct {
		jwt.RegisteredClaims
		UserID   uint   `json:"user_id"`
		Role string `json:"role"`
	}

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
			"user_id": userId,
			"role": role,
		}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(secretKey)
	return ss, err
}

func getRefreshToken(userId uint) (string, error) {

	claims := jwt.MapClaims{
			"exp": jwt.NewNumericDate(time.Now().Add(168 * time.Hour)),
			"user_id": userId,
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

func ParseToken(tokenString string) (*UserClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(token *jwt.Token) (any, error) {
		return secretKey, nil
	})

	claims, ok := token.Claims.(*UserClaims);
	if ok {
		return claims, nil
	}
	return nil, err
}