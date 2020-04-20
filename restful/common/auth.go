package common

import (
	jwt "github.com/dgrijalva/jwt-go"

	"github.com/gin-gonic/gin"

	"time"

	"net/http"
)

var jwtSecret = []byte("secret-key")

type Claims struct {
	Uid    int    `json:"uid"`
	OpenId string `json:"openId"`
	jwt.StandardClaims
}

func GenerateToken(uid int, openId string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour)

	claims := Claims{
		uid,
		openId,
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
			c.Set("uid", claims.Uid)

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

// 不需要强制验证
func JWTNoMust() gin.HandlerFunc {
	return func(c *gin.Context) {

		token := c.Query("token")
		if token == "" {
			token = c.GetHeader("X-Token")
			if token == "" {
				c.Set("uid", 0)
			}
		}
		claims, err := ParseToken(token)
		if err == nil && claims != nil {
			c.Set("uid", claims.Uid)
		}

		c.Next()
	}
}
