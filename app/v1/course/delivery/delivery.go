package delivery

import (
	"context"
	"fmt"
	"net/http"
	"teacher-site/config"
	"teacher-site/domain"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Usecase domain.CourseUsecase
	conf    *config.Config
}

func NewHandler(ctx context.Context, r *gin.RouterGroup, usecase domain.CourseUsecase, conf *config.Config) {
	handler := &Handler{
		Usecase: usecase,
		conf:    conf,
	}

	r.GET("", handler.Get(ctx))
	r.POST("", handler.Create(ctx))
	course := r.Group("/:courseId")
	{
		course.GET("", handler.GetContent(ctx))

		bulletin := course.Group("/bulletin")
		{
			bulletin.POST("", handler.CreateBulletin(ctx))
			bulletin.PUT("/:bulletinId", handler.UpdateBulletin(ctx))
			bulletin.DELETE("/:bulletinId", handler.DeleteBulletin(ctx))
		}

		slide := course.Group("/slide")
		{
			slide.POST("", handler.CreateSlide(ctx))
			slide.PUT("/:slideId", handler.UpdateSlide(ctx))
			slide.DELETE("/:slideId", handler.DeleteSlide(ctx))
		}

		homework := course.Group("/homework")
		{
			homework.POST("", handler.CreateHomework(ctx))
			homework.PUT("/:homeworkId", handler.UpdateHomework(ctx))
			homework.DELETE("/:homeworkId", handler.DeleteHomework(ctx))
		}

	}
}
func (h *Handler) Create(ctx context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req domain.CreateCourseRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		res, err := h.Usecase.Create(ctx, &req)
		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": res})
	}
}
func (h *Handler) CreateBulletin(ctx context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req domain.CreateCourseBulletinRequest
		if err := c.ShouldBindUri(&req); err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			// todo: test
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		if (req.CourseId == 0) || (len(req.Content) == 0) {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		res, err := h.Usecase.CreateBulletin(ctx, &req)
		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": res})
	}
}
func (h *Handler) CreateSlide(ctx context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req domain.CreateCourseSlideRequest
		if err := c.ShouldBindUri(&req); err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			// todo: test
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		if (req.CourseId == 0) || (len(req.Chapter) == 0) || (len(req.FileTitle) == 0) {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		res, err := h.Usecase.CreateSlide(ctx, &req)
		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": res})
	}
}
func (h *Handler) CreateHomework(ctx context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req domain.CreateCourseHomeworkRequest
		if err := c.ShouldBindUri(&req); err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			// todo: test
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		if (req.CourseId == 0) || (len(req.Number) == 0) || (len(req.FileTitle) == 0) {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		res, err := h.Usecase.CreateHomework(ctx, &req)
		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": res})
	}
}

func (h *Handler) Get(ctx context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		// todo: test
		var req domain.GetCourseRequest
		if err := c.ShouldBindUri(&req); err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		res, err := h.Usecase.Get(ctx, &req)
		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": res})

	}
}
func (h *Handler) GetContent(ctx context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req domain.GetCourseContentRequest
		if err := c.ShouldBindUri(&req); err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		if err := c.ShouldBindQuery(&req); err != nil {
			// todo: test
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		res, err := h.Usecase.GetContent(ctx, &req)
		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": res})

	}
}

func (h *Handler) UpdateBulletin(ctx context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req domain.UpdateCourseBulletinRequest
		if err := c.ShouldBindUri(&req); err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			// todo: test
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		if (req.CourseId == 0) || (req.BulletinId == 0) || (len(req.Content) == 0) {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		res, err := h.Usecase.UpdateBulletin(ctx, &req)
		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": res})
	}
}
func (h *Handler) UpdateSlide(ctx context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req domain.UpdateCourseSlideRequest
		if err := c.ShouldBindUri(&req); err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			// todo: test
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		if (req.CourseId == 0) || (req.SlideId == 0) || (len(req.Chapter) == 0) || (len(req.FileTitle) == 0) {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		res, err := h.Usecase.UpdateSlide(ctx, &req)
		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": res})
	}
}
func (h *Handler) UpdateHomework(ctx context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req domain.UpdateCourseHomeworkRequest
		if err := c.BindUri(&req); err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			// todo: test
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		if (req.CourseId == 0) || (req.HomeworkId == 0) || (len(req.Number) == 0) || (len(req.FileTitle) == 0) {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		res, err := h.Usecase.UpdateHomework(ctx, &req)
		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": res})
	}
}
func (h *Handler) DeleteBulletin(ctx context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req domain.DeleteCourseBulletinRequest
		if err := c.ShouldBindUri(&req); err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		res, err := h.Usecase.DeleteBulletin(ctx, &req)
		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": res})
	}
}
func (h *Handler) DeleteSlide(ctx context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req domain.DeleteCourseSlideRequest
		if err := c.ShouldBindUri(&req); err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		res, err := h.Usecase.DeleteSlide(ctx, &req)
		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": res})
	}
}
func (h *Handler) DeleteHomework(ctx context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req domain.DeleteCourseHomeworkRequest
		if err := c.ShouldBindUri(&req); err != nil {
			fmt.Println(err)

			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		res, err := h.Usecase.DeleteHomework(ctx, &req)
		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": res})
	}
}
