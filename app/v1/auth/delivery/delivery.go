package delivery

import (
	"context"
	"net/http"
	"teacher-site/config"
	"teacher-site/domain"
	mw "teacher-site/middleware"
	"teacher-site/pkg/util"

	"log"

	"github.com/gin-contrib/sessions"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	Usecase domain.AuthUsecase
	conf    *config.Config
}

func NewHandler(ctx context.Context, r *gin.RouterGroup, usecase domain.AuthUsecase, conf *config.Config) {
	handler := &AuthHandler{
		Usecase: usecase,
		conf:    conf,
	}

	r.POST("/token", handler.GetToken(ctx))
	r.POST("/login", handler.Login(ctx))
	r.POST("/logout", handler.Logout(ctx))
	// auth.POST("/register", handler.Register)
}
func (auth *AuthHandler) GetToken(ctx context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		s := sessions.Default(c)
		token := util.GetSessionToken(s)
		util.AddBearerHeader(c, token.(string))
		c.Status(http.StatusOK)
	}
}

// todo
func (auth *AuthHandler) Login(ctx context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			req domain.LoginRequest
			err error
		)
		if mw.IsTeacher(ctx, c, auth.conf.Jwt.Secret) {
			// todo: how to get the domain
			c.AbortWithStatus(http.StatusFound)
			// c.Redirect(http.StatusFound, "/")
			// c.Abort()
			return
		}

		if err = c.ShouldBindJSON(&req); err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		token, err := auth.Usecase.Login(ctx, &req)
		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		s := sessions.Default(c)
		util.SetSessionToken(s, token)
		s.Save()
		util.AddBearerHeader(c, token)
		c.Status(http.StatusNoContent)
	}
}
func (auth *AuthHandler) Logout(ctx context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		bearerToken, err := util.GetBearerToken(ctx, c)
		// ? token not existed
		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		claims, err := util.ParseJwt(ctx, bearerToken, auth.conf.Jwt.Secret)

		if err != nil {
			log.Println(err)
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		id := util.GetJwtUser(claims)
		err = auth.Usecase.Logout(ctx, id)
		if err != nil {
			log.Println(err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		util.RemoveBearerHeader(c)
		c.Status(http.StatusNoContent)
	}
}
