package middleware

import (
	"github.com/gin-gonic/gin"
	myErr "love_knot/pkg/error"
	"love_knot/pkg/jwt"
	"love_knot/pkg/response"
	"strconv"
	"strings"
)

const JWTSessionConst = "__JWT_SESSION__"

type JSession struct {
	Uid       int    `json_utils:"uid"`
	Token     string `json_utils:"token"`
	ExpiresAt int64  `json_utils:"expires_at"`
}

var (
	ErrorNoLogin = myErr.Unauthorized("", "未登录！")
)

func Auth(guard string, secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := GetHeaderToken(c)

		claims, err := verify(guard, secret, token)
		if err != nil {
			response.ErrResponse(c, err)
			return
		}

		uid, err1 := strconv.Atoi(claims.ID)
		if err1 != nil {
			response.ErrResponse(c, myErr.InternalServerError("", "解析 JWT 失败！"))
			return
		}

		c.Set(JWTSessionConst, JSession{
			Uid:       uid,
			Token:     token,
			ExpiresAt: claims.ExpiresAt.Unix(),
		})

		c.Next()
	}
}

func GetHeaderToken(c *gin.Context) string {
	token := c.GetHeader("Authorization")
	token = strings.TrimSpace(strings.TrimPrefix(token, "Bearer"))

	if token == "" {
		token = c.DefaultQuery("token", "")
	}

	return token
}

func verify(guard string, secret string, token string) (*jwt.AuthClaims, error) {
	if token == "" {
		return nil, ErrorNoLogin
	}

	claims, err := jwt.ParseToken(token, secret)
	if err != nil {
		return nil, err
	}

	if claims.Guard != guard || claims.Valid() != nil {
		return nil, ErrorNoLogin
	}

	return claims, nil
}
