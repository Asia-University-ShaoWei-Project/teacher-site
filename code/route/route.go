package route

import (
	"context"
	"net/http"
	mw "teacher-site/middleware"
	v1 "teacher-site/route/v1"
	"teacher-site/service"

	"github.com/gin-gonic/gin"
)

const templateIndex = "index.html"

func SetupRoute(ctx context.Context, r *gin.Engine, srv service.Servicer) {
	r.GET("/:domain", func(c *gin.Context) {
		c.HTML(http.StatusOK, templateIndex, gin.H{})
	})
	api := r.Group("/api")
	{
		// setupV1(ctx, api, srv)
		// }

		// func setupV1(ctx context.Context, api *gin.RouterGroup, srv service.Servicer) {
		route := api.Group("/v1/:domain", mw.SetupServiceDomain(ctx, srv))
		{
			// route.GET("/init", v1.GetInit(ctx, srv))

			// course := route.Group("/course/:course_id/:last_updated")
			// {
			// 	course.GET("/", v1.GetCourse(ctx, srv))
			// }

			auth := route.Group("/auth")
			{
				auth.GET("/test", func(c *gin.Context) { c.String(200, "teset", "") })
				auth.POST("/login", v1.Login(ctx, srv))
				auth.POST("/register", v1.Register(ctx, srv))
				edit := auth.Group("/edit", mw.VerifyJWT(ctx, srv))
				{
					info := edit.Group("/info")
					{
						// todo: resource id
						info.POST("/", v1.CreateInfo(ctx, srv))
						info.PUT("/", v1.UpdateInfo(ctx, srv))
						info.DELETE("/", v1.DeleteInfo(ctx, srv))
					}
				}
				// todo

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

		}
	}
}
