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
	notFoundHtml    = "not-found.html"
)

type Handler struct {
	Usecase domain.PageUsecase
	conf    *config.Config
}

func NewHandler(ctx context.Context, r *gin.RouterGroup, usecase domain.PageUsecase, conf *config.Config) {
	handler := &Handler{
		Usecase: usecase,
		conf:    conf,
	}
	r.GET("", handler.TeacherList(ctx))
	// todo: get teacher list by api
	r.GET("/page/:page_number", handler.TeacherList(ctx))
	r.GET("/:teacherDomain", handler.Home(ctx))
	r.GET("/login", handler.Login(ctx, conf.Jwt))
}
func (h *Handler) TeacherList(ctx context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req domain.TeacherListRequest
		if err := c.ShouldBindUri(&req); err != nil {
			c.HTML(http.StatusNotFound, notFoundHtml, gin.H{})
			return
		}
		if req.Page == 0 {
			req.SetToFirstPage()
		}
		res, err := h.Usecase.TeacherList(ctx, &req)
		if err != nil {
			c.HTML(http.StatusInternalServerError, serverErrorHtml, gin.H{})
			return
		}
		c.HTML(http.StatusOK, teacherListHtml, gin.H{"teachers": res})
	}
}
func (h *Handler) Home(ctx context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req domain.HomeRequest

		if err := c.ShouldBindUri(&req); err != nil {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
		res, err := h.Usecase.Home(ctx, &req)
		if err != nil {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
		c.HTML(http.StatusOK, homeHtml, gin.H{"data": res})
	}
}

func (h *Handler) Login(ctx context.Context, conf *config.Jwt) gin.HandlerFunc {
	return func(c *gin.Context) {
		s := sessions.Default(c)
		token := util.GetSessionToken(s)
		if token == nil {
			c.HTML(http.StatusOK, loginHtml, gin.H{})
			return
		}
		if claims, err := util.ParseJwt(ctx, token.(string), conf.Secret); err == nil {
			userId := util.GetJwtUser(claims)
			if err := h.Usecase.Login(ctx, userId, token.(string)); err == nil {
				c.Redirect(http.StatusFound, util.GetJwtUserDomain(claims))
				return
			}
		}
		c.HTML(http.StatusOK, loginHtml, gin.H{})
	}
}
