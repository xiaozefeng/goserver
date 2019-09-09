package token

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"time"
)

var (
	ErrMissingHeader = errors.New("The length of the `Authorization` header is zero.")
)

type Context struct {
	Id       uint64
	Username string
}

func Sign(c *gin.Context, ctx Context, secret string) (string, error) {
	// load the jwt secret from the config if secret isn't specified.
	if secret == "" {
		secret = viper.GetString("jwt_secret")
	}

	// generate token content
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       ctx.Id,
		"username": ctx.Username,
		"nbf":      time.Now().Unix(),
		"iat":      time.Now().Unix(),
	})

	return token.SignedString([]byte(secret))
}

func ParseRequest(c *gin.Context) (*Context, error) {
	header := c.Request.Header.Get("Authorization")

	// load the jwt secret from config
	secret := viper.GetString("jwt_secret")
	if len(header) == 0 {
		return &Context{}, ErrMissingHeader
	}

	var t string
	_, _ = fmt.Sscanf(header, "Bearer %s", &t)
	return parse(t, secret)

}

func parse(t string, secret string) (*Context, error) {
	res := &Context{}
	token, err := jwt.Parse(t, secureFunc(secret))
	if err != nil {
		return res, err
	} else if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		res.Id = uint64(claims["id"].(float64))
		res.Username = claims["username"].(string)
		return res, nil
	} else {
		return res, err
	}
}

// secureFunc validates the secret format
func secureFunc(secure string) jwt.Keyfunc {
	return func(t *jwt.Token) (interface{}, error) {
		// make sure the `alg` is what we expect
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(secure), nil
	}
}
