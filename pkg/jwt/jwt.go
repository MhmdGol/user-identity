package jwt

import (
	"Identity/cmd/config"
	"Identity/internal/model"
	"crypto/rsa"
	"fmt"
	"os"

	"github.com/bwmarrin/snowflake"
	"github.com/dgrijalva/jwt-go"
)

type JwtToken struct {
	SecretKey *rsa.PrivateKey
	PublicKey *rsa.PublicKey
}

func NewJwtHandler(conf config.RSAPair) *JwtToken {
	privateKeyBytes, err := os.ReadFile(conf.SecretKeyPath)
	if err != nil {
		panic(err)
	}
	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(privateKeyBytes)
	if err != nil {
		panic(err)
	}

	publicKeyBytes, err := os.ReadFile(conf.PublicKeyPath)
	if err != nil {
		panic(err)
	}
	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(publicKeyBytes)
	if err != nil {
		panic(err)
	}

	return &JwtToken{
		SecretKey: privateKey,
		PublicKey: publicKey,
	}
}

func (j *JwtToken) MakeToken(c model.TokenClaim) (model.JwtToken, error) {

	claims := jwt.MapClaims{
		"id":       c.ID,
		"username": c.Username,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	tokenString, err := token.SignedString(j.SecretKey)

	return model.JwtToken(tokenString), err
}

func (j *JwtToken) ExtractClaims(t model.JwtToken) (model.TokenClaim, error) {
	token, err := jwt.Parse(string(t), func(token *jwt.Token) (interface{}, error) {
		return j.PublicKey, nil
	})
	if err != nil {
		return model.TokenClaim{}, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		id, ok := claims["id"].(string)
		if !ok {
			return model.TokenClaim{}, fmt.Errorf("invalid token: id not found")
		}
		sfId, _ := snowflake.ParseString(id)

		username, ok := claims["username"].(string)
		if !ok {
			return model.TokenClaim{}, fmt.Errorf("invalid token: username not found")
		}

		return model.TokenClaim{
			ID:       model.ID(sfId),
			Username: username,
		}, nil
	}
	return model.TokenClaim{}, fmt.Errorf("invalid token")
}
