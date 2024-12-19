package lib

import (
	"os"
	"time"

	"github.com/go-jose/go-jose/v4"
	"github.com/go-jose/go-jose/v4/jwt"
	"github.com/joho/godotenv"
)

func GenerateToken(payload any) string {
	godotenv.Load()
	var JWT_SECRET []byte = []byte(GetMD5Hash(os.Getenv("JWT_SECRET")))
	sig, _ := jose.NewSigner(jose.SigningKey{Algorithm: jose.HS256, Key: JWT_SECRET}, (&jose.SignerOptions{}).WithType("JWT"))
	baseInfo := jwt.Claims{
		IssuedAt: jwt.NewNumericDate(time.Now()),
	}

	token, _ := jwt.Signed(sig).Claims(baseInfo).Claims(payload).Serialize()

	return token

}
