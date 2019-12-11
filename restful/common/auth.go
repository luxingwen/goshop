package common

import (
	jwt "github.com/dgrijalva/jwt-go"

	"github.com/gin-gonic/gin"

	"time"

	"net/http"
)

var jwtSecret = []byte("secret-key")

type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

func GenerateToken(username, password string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour)

	claims := Claims{
		username,
		password,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "lxw-goshop",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}

		token := c.Query("token")
		if token == "" {
			token = c.GetHeader("X-Token")
			if token == "" {
				code = 999
			}

		} else {
			claims, err := ParseToken(token)
			if err != nil {
				code = 1
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = 2
			}
		}

		if code != 0 {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  "无效的token",
				"data": data,
			})

			c.Abort()
			return
		}

		c.Next()
	}
}
