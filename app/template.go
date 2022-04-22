package app

import (
	"context"
	"fmt"
	"net/http"
	"teacher-site/config"
	"teacher-site/pkg/util"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

const homeHtml = "home.html"
const teacherHtml = "index.html"
const loginHtml = "login.html"

type Template struct{}

func NewTemplate() *Template {
	return &Template{}

}

func (t *Template) Home(ctx context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, homeHtml, gin.H{})
	}
}
func (t *Template) Login(ctx context.Context, conf *config.Jwt) gin.HandlerFunc {
	return func(c *gin.Context) {
		// todo: check session error, can't get the token!!
		s := sessions.Default(c)
		token := util.GetSessionToken(s)
		if token != nil {
			// todo: error here!
			fmt.Println("not have cookie token")
			c.HTML(http.StatusOK, loginHtml, gin.H{})
			return
		}
		if _, err := util.ParseJwt(ctx, token.(string), conf.Secret); err != nil {
			c.HTML(http.StatusOK, loginHtml, gin.H{})
			return
		}
		c.HTML(http.StatusOK, teacherHtml, gin.H{})

	}
}
func (t *Template) TeacherSite(ctx context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		// todo: setup the cookie and authorization header
		s := sessions.Default(c)
		token := util.GetSessionToken(s)
		if token != nil {
			// todo: check expiration and certify the token with db
			util.AddBearerHeader(c, token.(string))
		}
		c.HTML(http.StatusOK, teacherHtml, gin.H{})
	}
}
