package jwt

import (
	"crypto"
	"crypto/rsa"
	jwtLib "github.com/golang-jwt/jwt/v4"
	"log"
)

// Generator jwt token generator interface
type Generator interface {
	Generate(claims jwtLib.Claims, privateKey *rsa.PrivateKey) (*string, error)
	GetPrivateKey() *rsa.PrivateKey
	GetExp() int
}

// GeneratorStruct jwt token generator struct
type GeneratorStruct struct {
	PrivateKey *rsa.PrivateKey
	PublicKey  crypto.PublicKey
	Exp        int
}

// NewGenerator create new jwt token generator
func NewGenerator(priv *rsa.PrivateKey, pub crypto.PublicKey, exp int) Generator {
	return &GeneratorStruct{
		priv, pub, exp,
	}
}

// Generate create new jwt token
func (g *GeneratorStruct) Generate(claims jwtLib.Claims, privateKey *rsa.PrivateKey) (*string, error) {
	token := jwtLib.NewWithClaims(jwtLib.SigningMethodRS256, claims)
	t, err := token.SignedString(privateKey)
	if err != nil {
		log.Printf("token.SignedString: %v", err)
		return nil, err
	}

	return &t, nil
}

// GetPrivateKey get RSA private key
func (g *GeneratorStruct) GetPrivateKey() *rsa.PrivateKey {
	return g.PrivateKey
}

// GetExp get token's expires time
func (g *GeneratorStruct) GetExp() int {
	return g.Exp
}
