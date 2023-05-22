package xjwt

import (
	"com.gientech/selection/pkg/config"
	"encoding/json"
	"github.com/cristalhq/jwt/v5"
)

type Key string

type JwtHelper struct {
	key []byte
}

func NewJwtHelper(config *config.Config) *JwtHelper {
	return &JwtHelper{
		key: []byte(config.Viper.GetString("jwt.key")),
	}
}

type UserClaims struct {
	jwt.RegisteredClaims
	Role string // 角色
}

// CreateToken 生成 token
func (a *JwtHelper) CreateToken(claims UserClaims) (*jwt.Token, error) {
	signer, _ := jwt.NewSignerHS(jwt.HS256, a.key)
	// create a Builder
	builder := jwt.NewBuilder(signer)
	// and build a Token
	return builder.Build(claims)
}

// VerifyToken 解析并校验 token
func (a *JwtHelper) VerifyToken(token string) (*jwt.Token, error) {
	verifier, _ := jwt.NewVerifierHS(jwt.HS256, a.key)
	return jwt.Parse([]byte(token), verifier)
}

func (a *JwtHelper) ParseToken(token *jwt.Token) (UserClaims, error) {
	var newClaims UserClaims
	err := json.Unmarshal(token.Claims(), &newClaims)
	return newClaims, err
}
