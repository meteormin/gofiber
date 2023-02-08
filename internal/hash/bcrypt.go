package hash

import "golang.org/x/crypto/bcrypt"

var Bcrypt bcryptWrapper

type bcryptWrapper struct {
	cost int
}

func init() {
	Bcrypt = bcryptWrapper{
		cost: 14,
	}
}

func (b bcryptWrapper) HashPassword(password string) (string, error) {
	fromPassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}

	return string(fromPassword), err
}

func (b bcryptWrapper) HashCheck(hashPass string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPass), []byte(password))
	return err == nil
}
