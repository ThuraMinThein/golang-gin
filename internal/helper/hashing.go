package helper

import "golang.org/x/crypto/bcrypt"

func Hash(str string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(str), 12)
    return string(bytes), err
}

func VerifyHashed(str, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(str))
    return err == nil
}