package cookie

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

const (
	tokenKey = "token"
)

// todo
func SetToken(c *gin.Context, token string, maxAge int) {
	fmt.Println("set the token")

	c.SetCookie(tokenKey, token, maxAge, "/", "location", false, true)

}

// todo
func Token(c *gin.Context) (string, error) {
	fmt.Println(c.Writer.Header().Get(tokenKey))

	return c.Cookie(tokenKey)
}
