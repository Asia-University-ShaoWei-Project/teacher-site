package delivery

import (
	"context"
	"net/http"
	"teacher-site/config"
	"teacher-site/domain"
	mw "teacher-site/middleware"
	"teacher-site/pkg/util"

	"github.com/gin-contrib/sessions"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Usecase domain.AuthUsecase
	conf    *config.Config
}

func NewHandler(ctx context.Context, r *gin.RouterGroup, usecase domain.AuthUsecase, conf *config.Config) {
	handler := &Handler{
		Usecase: usecase,
		conf:    conf,
	}

	r.POST("/token", handler.GetToken(ctx))
	r.POST("/login", handler.Login(ctx))
	r.POST("/logout", handler.Logout(ctx))

	// todo
	r.POST("/register", handler.Register(ctx))
}
func (h *Handler) GetToken(ctx context.Context) gin.HandlerFunc {
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
			if _, err := util.ParseJwt(ctx, _token, h.conf.Jwt.Secret); err == nil {
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

func (h *Handler) Login(ctx context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			req domain.LoginRequest
			err error
		)

		if mw.IsTeacher(ctx, c, h.conf.Jwt.Secret) {
			// todo: redirect the teacher domain(response with doamin data)
			c.AbortWithStatus(http.StatusFound)
			return
		}

		if err = c.ShouldBindJSON(&req); err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		res, err := h.Usecase.Login(ctx, &req)
		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		// todo: use httponly cookie instead of the session
		// c.SetSameSite(http.SameSiteDefaultMode)
		// utilcookie.SetToken(c, res.Token, h.conf.Secure.CookieTokenMaxAge)

		s := sessions.Default(c)
		util.SetSessionToken(s, res.Token)
		util.AddBearerHeader(c, res.Token)
		s.Save()
		// todo: redirect the teacher domain(response with doamin data)
		c.JSON(http.StatusFound, gin.H{"domain": res.Domain})
	}
}

func (h *Handler) Logout(ctx context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {

		bearerToken, err := util.GetBearerToken(ctx, c)
		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		claims, err := util.ParseJwt(ctx, bearerToken, h.conf.Jwt.Secret)
		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		id := util.GetJwtUser(claims)

		if err = h.Usecase.Logout(ctx, id); err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		s := sessions.Default(c)
		util.DeleteSessionToken(s)
		s.Save()
		c.Status(http.StatusNoContent)
	}
}

func (h *Handler) Register(ctx context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req *domain.RegisterRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		if err := h.Usecase.Register(ctx, req); err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		c.Status(http.StatusCreated)
	}
}
