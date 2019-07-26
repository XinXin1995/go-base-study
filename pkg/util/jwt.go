package util

import (
	"blog/models"
	"blog/pkg/set"
	"crypto/md5"
	"encoding/hex"
	"github.com/dgrijalva/jwt-go"
	"github.com/satori/go.uuid"
	"time"
)

var jwtSecret = []byte(set.JwtSecret)

type Claims struct {
	Uuid     uuid.UUID `json:"id"`
	Username string    `json:"username"`
	Password string    `json:"password"`
	RoleUuid string    `json:"roleUuid"`
	jwt.StandardClaims
}

func GenerateToken(user *models.User) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour)
	claims := Claims{
		user.Uuid,
		user.Name,
		user.Password,
		user.RoleUuid,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "gin-blog",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (i interface{}, e error) {
		return jwtSecret, nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}

func GeneratePwd(password string, salt string) (newPassword string) {
	m5 := md5.New()
	m5.Write([]byte(password))
	m5.Write([]byte(salt))
	st := m5.Sum(nil)
	newPassword = hex.EncodeToString(st)
	return
}
