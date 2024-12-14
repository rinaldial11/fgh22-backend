package lib

import (
	"time"

	"github.com/go-jose/go-jose/v4"
	"github.com/go-jose/go-jose/v4/jwt"
)

var JWT_SECRET []byte = []byte(GetMD5Hash("SECRET_KEY"))

func GenerateToken(payload any) string {
	sig, _ := jose.NewSigner(jose.SigningKey{Algorithm: jose.HS256, Key: JWT_SECRET}, (&jose.SignerOptions{}).WithType("JWT"))
	baseInfo := jwt.Claims{
		IssuedAt: jwt.NewNumericDate(time.Now()),
	}

	token, _ := jwt.Signed(sig).Claims(baseInfo).Claims(payload).Serialize()

	return token

}
