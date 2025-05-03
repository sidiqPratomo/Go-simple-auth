package util

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/sidiqPratomo/DJKI-Pengaduan/config"
)

type JwtCustomClaims struct {
	UserId        int64  `json:"user_id"`
	Email         string `json:"email"`
	Role          string `json:"role"`
}

type TokenAuthentication interface {
	CreateAndSign(customClaims JwtCustomClaims, secretKey string) (*string, *string, error)
	ParseAndVerify(signed string, secretKey string) (*JwtCustomClaims, error)
}

type JwtAuthentication struct {
	Config config.Config
	Method *jwt.SigningMethodHMAC
}

func (ja JwtAuthentication) CreateAndSign(customClaims JwtCustomClaims, secretKey string) (*string, *string, error) {
	customClaimsJsonBytes, err := json.Marshal(customClaims)
	if err != nil {
		return nil, nil, err
	}

	now := time.Now()

	seconds := now.Unix()
	nanoseconds := now.UnixNano() % 1e9

	decimalTime := float64(seconds) + float64(nanoseconds)/1e9

	expiredAt := time.Now().Add(time.Duration(ja.Config.TokenDuration) * time.Minute).Unix()

	token := jwt.NewWithClaims(ja.Method, jwt.MapClaims{
		"jti":  generateJTI(),
		"iss":  ja.Config.Issuer,
		"iat":  decimalTime,
		"exp":  expiredAt,
		"data": string(customClaimsJsonBytes),
	})

	expired := time.Now().Add(time.Duration(ja.Config.TokenDuration) * time.Minute).Format("2006-01-02 15:04:05")

	signed, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return nil, nil, err
	}

	return &signed, &expired, nil
}

func generateJTI() string {
	return fmt.Sprintf("%x", time.Now().UnixNano()) // Menghasilkan ID unik berbasis waktu
}

func (ja JwtAuthentication) ParseAndVerify(signed string, secretKey string) (*JwtCustomClaims, error) {
	token, err := jwt.Parse(signed, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	}, jwt.WithValidMethods([]string{ja.Method.Name}),
		jwt.WithIssuer(ja.Config.Issuer),
		jwt.WithExpirationRequired(),
	)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, err
	}

	customClaims := JwtCustomClaims{}
	data := claims["data"].(string)
	if err := json.Unmarshal([]byte(data), &customClaims); err != nil {
		return nil, err
	}

	return &customClaims, nil
}
