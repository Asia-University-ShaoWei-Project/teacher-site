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
		isTeacher := false
		// todo: checkout the domain is existed
		s := sessions.Default(c)
		// The token from the cookie
		token := util.GetSessionToken(s)
		if token != nil {
			_token := token.(string)
			// todo: check expiration and certify the token with db
			// Add token to the authorization header of response
			if _, err := util.ParseJwt(ctx, _token, auth.conf.Jwt.Secret); err == nil {
				isTeacher = true
				util.AddBearerHeader(c, _token)
			} else {
				util.DeleteSessionToken(s)
				s.Save()
			}
		}
		c.JSON(http.StatusOK, gin.H{
			"isTeacher": isTeacher,
		})
	}
}

func (auth *AuthHandler) Login(ctx context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {

		var (
			req domain.LoginRequest
			err error
		)

		if mw.IsTeacher(ctx, c, auth.conf.Jwt.Secret) {
			// todo: redirect the teacher domain(response with doamin data)
			c.AbortWithStatus(http.StatusFound)
			return
		}

		if err = c.ShouldBindJSON(&req); err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		res, err := auth.Usecase.Login(ctx, &req)
		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		s := sessions.Default(c)
		util.SetSessionToken(s, res.Token)
		util.AddBearerHeader(c, res.Token)
		s.Save()
		// todo: redirect the teacher domain(response with doamin data)
		c.JSON(http.StatusFound, gin.H{"domain": res.Domain})
	}
}

func (auth *AuthHandler) Logout(ctx context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {

		bearerToken, err := util.GetBearerToken(ctx, c)
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

		if err = auth.Usecase.Logout(ctx, id); err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		util.RemoveAuthHeader(c)
		c.Status(http.StatusNoContent)
	}
}
