package delivery

import (
	"context"
	"net/http"
	"teacher-site/config"
	"teacher-site/domain"
	"teacher-site/pkg/util"

	"github.com/gin-contrib/sessions"

	"github.com/gin-gonic/gin"
)

const (
	teacherListHtml = "teacher-list.html"
	homeHtml        = "home.html"
	loginHtml       = "login.html"
	serverErrorHtml = "server-error.html"
)

type PageHandler struct {
	Usecase domain.PageUsecase
	conf    *config.Config
}

func NewHandler(ctx context.Context, r *gin.RouterGroup, usecase domain.PageUsecase, conf *config.Config) {
	handler := &PageHandler{
		Usecase: usecase,
		conf:    conf,
	}

	r.GET("/", handler.TeacherListPage(ctx))
	// todo: get teacher list by api
	r.GET("/page/:page_number", handler.TeacherList(ctx))
	r.GET("/:teacher_domain", handler.Home(ctx))
	r.GET("/login", handler.Login(ctx, conf.Jwt))
}
func (p *PageHandler) TeacherListPage(ctx context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, teacherListHtml, gin.H{})
	}
}

func (p *PageHandler) TeacherList(ctx context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req domain.TeacherListRequest
		if err := c.ShouldBindUri(&req); err != nil {
			c.Status(http.StatusBadRequest)
			return
		}
		// issue: A negative digit of the input
		if req.Page < 0 {
			c.Status(http.StatusBadRequest)
			return
		}
		res, err := p.Usecase.TeacherList(ctx, &req)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": res})
	}
}

// todo: get teacher list by api
// func (p *PageHandler) TeacherListByApi(ctx context.Context) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		var req domain.TeacherListRequest
// 		if err := c.ShouldBindUri(&req); err != nil {
// 			c.AbortWithStatus(http.StatusBadRequest)
// 			return
// 		}
// 		res, err := p.Usecase.TeacherList(ctx, &req)
// 		if err != nil {
// 			c.AbortWithStatus(http.StatusBadRequest)
// 			return
// 		}
// 		c.JSON(http.StatusOK, gin.H{"data":res})
// 	}
// }
func (p *PageHandler) Home(ctx context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req domain.HomeRequest

		if err := c.ShouldBindUri(&req); err != nil {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
		res, err := p.Usecase.Home(ctx, &req)
		if err != nil {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
		c.HTML(http.StatusOK, homeHtml, gin.H{"data": res})
	}
}

func (p *PageHandler) Login(ctx context.Context, conf *config.Jwt) gin.HandlerFunc {
	return func(c *gin.Context) {
		// todo: check session error, can't get the token!!
		s := sessions.Default(c)
		token := util.GetSessionToken(s)
		if token == nil {
			c.HTML(http.StatusOK, loginHtml, gin.H{})
			return
		}
		if claims, err := util.ParseJwt(ctx, token.(string), conf.Secret); err == nil {
			userId := util.GetJwtUser(claims)
			if err := p.Usecase.Login(ctx, userId, token.(string)); err == nil {
				c.Redirect(http.StatusFound, util.GetJwtUserDomain(claims))
				return
			}
		}
		c.HTML(http.StatusOK, loginHtml, gin.H{})
	}
}
