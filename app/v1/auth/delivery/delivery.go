package delivery

import (
	"context"
	"net/http"
	"teacher-site/config"
	"teacher-site/domain"
	mw "teacher-site/middleware"

	"log"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	Usecase domain.AuthUsecase
	conf    *config.Jwt
}

func NewAuthHandler(ctx context.Context, r *gin.RouterGroup, usecase domain.AuthUsecase, conf *config.Jwt) {
	handler := &AuthHandler{
		Usecase: usecase,
		conf:    conf,
	}

	r.POST("/login", handler.Login(ctx))
	// auth.POST("/register", handler.Register)
}

// todo
func (auth *AuthHandler) Login(ctx context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			req domain.LoginRequest
			err error
		)
		if mw.IsTeacher(ctx, c, auth.conf.Secure) {
			log.Print("is teacher -> redirect index")
			// todo: how to get the domain
			c.Redirect(http.StatusFound, "/")
			c.Abort()
			return
		}

		if err = c.ShouldBindJSON(&req); err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			log.Print(err)
			return
		}
		token, err := auth.Usecase.Login(ctx, &req)
		if err != nil {
			log.Print(err)
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		addBearerHeader(c, token)
		// todo: redirect 302 index
		c.Status(http.StatusOK)
	}
}
func addBearerHeader(c *gin.Context, token string) {
	c.Request.Header.Add("Authorization", `Bearer `+token)
}

// func (auth *AuthHandler) Get(ctx context.Context) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		var bind domain.ReqGetAuth
// 		if err := c.ShouldBindUri(&bind); err != nil {
// 			c.AbortWithStatus(http.StatusBadRequest)
// 			return
// 		}
// 		res, err := i.Usecase.Get(ctx, &bind)
// 		if err != nil {
// 			c.AbortWithStatus(http.StatusNoContent)
// 			return
// 		}
// 		c.JSON(http.StatusOK, &gin.H{
// 			"data": res,
// 		})
// 	}
// }

// func (auth *AuthHandler) Update(ctx context.Context) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		var bind domain.ReqUpdateInfoBulletin
// 		if err := c.ShouldBindUri(&bind); err != nil {
// 			c.AbortWithStatus(http.StatusBadRequest)
// 			return
// 		}
// 		if err := c.ShouldBindJSON(&bind); err != nil {
// 			c.AbortWithStatus(http.StatusBadRequest)
// 			return
// 		}
// 		res, err := i.Usecase.Update(ctx, &bind)
// 		if err != nil {
// 			c.AbortWithStatus(http.StatusBadRequest)
// 			return
// 		}
// 		c.JSON(http.StatusOK, gin.H{
// 			"data": res,
// 		})
// 		c.Status(http.StatusOK)
// 	}
// }

// func (auth *AuthHandler) Delete(ctx context.Context) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		var bind domain.ReqDeleteInfo
// 		if err := c.ShouldBindUri(&bind); err != nil {
// 			c.AbortWithStatus(http.StatusBadRequest)
// 			return
// 		}
// 		if err := i.Usecase.Delete(ctx, &bind); err != nil {
// 			c.AbortWithStatus(http.StatusBadRequest)
// 			return
// 		}
// 		c.Status(http.StatusNoContent)
// 	}
// }
