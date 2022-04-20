package delivery

import (
	"context"
	"net/http"
	"teacher-site/config"
	"teacher-site/domain"
	mw "teacher-site/middleware"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
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
	auth := r.Group("/auth")
	{
		auth.POST("/login", handler.Login)
		auth.GET("/", handler.Get(ctx))
		_auth := auth.Group("/", mw.VerifyAuth(ctx, conf))
		{
			// _auth.POST("/", handler.Create(ctx))
			_auth.PUT("/:id", handler.Update(ctx))
			_auth.DELETE("/:id", handler.Delete(ctx))
		}
	}
}

// func (auth *AuthHandler) Create(ctx context.Context) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		var bind domain.ReqCreateAuth
// 		if err := c.ShouldBindJSON(&bind); err != nil {
// 			c.AbortWithStatus(http.StatusBadRequest)
// 			return
// 		}
// 		bulletin, err := i.Usecase.Create(ctx, &bind)
// 		if err != nil {
// 			c.AbortWithStatus(http.StatusBadRequest)
// 			return
// 		}
// 		res := domain.ResCreateInfo{
// 			ID:   bulletin.AutoModel.ID,
// 			Date: bulletin.AutoModel.CreatedAT.Format(domain.BulletinDateFormat),
// 		}
// 		c.JSON(http.StatusOK, gin.H{
// 			"data": res,
// 		})
// 	}
// }
// todo
func (auth *AuthHandler) Login(ctx context.Context, req *domain.ReqLoginAuth) gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			token string
			err   error
		)
		if err = c.ShouldBindJSON(&req); err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		if err = srv.Login(ctx, &bind); err != nil {
			srv.Error(err)
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		token, err = srv.NewJwtToken(ctx, &bind)
		if err != nil {
			srv.Error(err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		if err = srv.UpdateJwtToken(ctx, token, &bind); err != nil {
			srv.Error(err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		addBearerHeader(c, token)
		c.Status(http.StatusOK)

		auth := &model.Auths{UserID: bind.UserID}
		if err = srv.db.GetAuth(ctx, auth); err != nil {
			// todo: handle error
			return err
		}
		saltPassword := []byte(bind.UserPassword + auth.Salt)
		if err = bcrypt.CompareHashAndPassword([]byte(auth.UserPassword), saltPassword); err != nil {
			// todo: handle error
			return err
		}
		return nil
	}
}
func (auth *AuthHandler) Get(ctx context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		var bind domain.ReqGetAuth
		if err := c.ShouldBindUri(&bind); err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		res, err := i.Usecase.Get(ctx, &bind)
		if err != nil {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.JSON(http.StatusOK, &gin.H{
			"data": res,
		})
	}
}

func (auth *AuthHandler) Update(ctx context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		var bind domain.ReqUpdateInfoBulletin
		if err := c.ShouldBindUri(&bind); err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		if err := c.ShouldBindJSON(&bind); err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		res, err := i.Usecase.Update(ctx, &bind)
		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"data": res,
		})
		c.Status(http.StatusOK)
	}
}

func (auth *AuthHandler) Delete(ctx context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		var bind domain.ReqDeleteInfo
		if err := c.ShouldBindUri(&bind); err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		if err := i.Usecase.Delete(ctx, &bind); err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		c.Status(http.StatusNoContent)
	}
}
