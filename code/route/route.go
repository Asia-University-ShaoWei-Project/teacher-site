package route

import (
	"context"
	"net/http"
	mw "teacher-site/middleware"
	"teacher-site/model"
	"teacher-site/service"

	"github.com/gin-gonic/gin"
)

func SetupRoute(ctx context.Context, r *gin.Engine, srv service.Servicer, cfg *model.Config) {
	v1 := r.Group("/v1/:domain", mw.SetupServiceDomain(ctx, srv))
	{
		r.GET("/:domain", func(c *gin.Context) {
			c.HTML(http.StatusOK, "index.html", gin.H{})
		})
		v1.GET("/init", GetInit(ctx, srv))

		course := v1.Group("/course/:course_id/:last_updated")
		{
			course.GET("/", GetCourse(ctx, srv))
		}
		edit := v1.Group("/edit", mw.VerifyJWT(ctx, cfg.JWTSecure))
		{
			info := edit.Group("/info")
			{
				info.POST("/", CreateInfo(ctx, srv))
				info.PUT("/", UpdateInfo(ctx, srv))
				info.DELETE("/", DeleteInfo(ctx, srv))
			}
			// 	courseEdit := edit.Group("/course/:course_id")
			// 	{
			// 		bulletinBoard := courseEdit.Group("/bulletin_board")
			// 		{
			// 			bulletinBoard.POST("/", createCourse)
			// 			bulletinBoard.PUT("/", updateCourse)
			// 			bulletinBoard.DELETE("/", deleteCourse)
			// 		}
			// 		slide := courseEdit.Group("/bulletin_board")
			// 		{
			// 			slide.POST("/", createSlide)
			// 			slide.PUT("/", updateSlide)
			// 			slide.DELETE("/", deleteSlide)
			// 		}
			// 		homework := courseEdit.Group("/bulletin_board")
			// 		{
			// 			homework.POST("/", createHomework)
			// 			homework.PUT("/", updateHomework)
			// 			homework.DELETE("/", deleteHomework)
			// 		}
			// 	}
		}
		auth := v1.Group("/auth")
		{
			auth.POST("/login", Login(ctx, srv, cfg))
			auth.POST("/register", Register(ctx, srv, cfg))
		}
	}
}
