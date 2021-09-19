package helper

import "golang.org/x/crypto/bcrypt"

func HassPassword(password string) (string, error) {
	byte, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(byte), err
}

func CheckPasswordHash(password, hash string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return false, err
	}

	return true, nil
}
