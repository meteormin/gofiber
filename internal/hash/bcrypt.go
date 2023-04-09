package hash

import "golang.org/x/crypto/bcrypt"

var Bcrypt BcryptWrapper

type BcryptWrapper struct {
	cost int
}

func init() {
	Bcrypt = BcryptWrapper{
		cost: 14,
	}
}

func (b BcryptWrapper) HashPassword(password string) (string, error) {
	fromPassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}

	return string(fromPassword), err
}

func (b BcryptWrapper) HashCheck(hashPass string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPass), []byte(password))
	return err == nil
}
