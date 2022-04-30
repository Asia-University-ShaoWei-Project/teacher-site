package delivery

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"teacher-site/config"
	"teacher-site/domain"
	mw "teacher-site/middleware"
	"teacher-site/pkg/message"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
	r.POST("", handler.Create(ctx), mw.VerifyAuth(ctx, conf.Jwt.Secret))
	course := r.Group("/:courseId")
	{
		course.GET("", handler.GetContent(ctx))

		auth := course.Group("", mw.VerifyAuth(ctx, conf.Jwt.Secret))
		{
			bulletin := auth.Group("/bulletin")
			{
				bulletin.POST("", handler.CreateBulletin(ctx))
				bulletin.PUT("/:bulletinId", handler.UpdateBulletin(ctx))
				bulletin.DELETE("/:bulletinId", handler.DeleteBulletin(ctx))
			}

			slide := auth.Group("/slide")
			{
				slide.POST("", handler.CreateSlide(ctx))
				slide.PUT("/:slideId", handler.UpdateSlide(ctx))
				slide.DELETE("/:slideId", handler.DeleteSlide(ctx))
			}

			homework := auth.Group("/homework")
			{
				homework.POST("", handler.CreateHomework(ctx))
				homework.PUT("/:homeworkId", handler.UpdateHomework(ctx))
				homework.DELETE("/:homeworkId", handler.DeleteHomework(ctx))
			}
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
		c.JSON(http.StatusCreated, gin.H{"data": res})
	}
}

func (h *Handler) CreateBulletin(ctx context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req domain.CreateCourseBulletinRequest
		if err := c.ShouldBindUri(&req); err != nil {
			fmt.Println(err)

			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			fmt.Println(err)
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		// todo: binding required
		if (req.CourseId == 0) || (len(req.Content) == 0) {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		res, err := h.Usecase.CreateBulletin(ctx, &req)
		if err != nil {
			fmt.Println(err)
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		c.JSON(http.StatusCreated, gin.H{"data": res})
	}
}
func (h *Handler) CreateSlide(ctx context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			uri domain.CreateSlideUriRequest
			req domain.CreateCourseSlideRequest
		)
		if err := c.ShouldBindUri(&uri); err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		if err := c.ShouldBind(&req); err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		req.SetupUri(&uri)
		if req.File != nil && req.File.Size != 0 {
			filename := uuid.New().String()
			dst := fmt.Sprintf(h.conf.Server.SlidePathFormat, req.TeacherDomain, filename)
			if err := c.SaveUploadedFile(req.File, dst); err != nil {
				fmt.Println(err)
				c.AbortWithStatus(http.StatusBadRequest)
				return
			}
			req.SetFilename(filename)
		}

		res, err := h.Usecase.CreateSlide(ctx, &req)
		if err != nil {
			fmt.Println(err)
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		c.JSON(http.StatusCreated, gin.H{"data": res})
	}
}

func (h *Handler) CreateHomework(ctx context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			uri domain.CreateHomeworkUriRequest
			req domain.CreateCourseHomeworkRequest
		)
		if err := c.ShouldBindUri(&uri); err != nil {
			fmt.Println(err)
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		if err := c.ShouldBind(&req); err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		req.SetupUri(&uri)
		if req.File != nil && req.File.Size != 0 {
			filename := uuid.New().String()
			dst := fmt.Sprintf(h.conf.Server.HomeworkPathFormat, req.TeacherDomain, filename)
			if err := c.SaveUploadedFile(req.File, dst); err != nil {
				fmt.Println(err)
				c.AbortWithStatus(http.StatusBadRequest)
				return
			}
			req.SetFilename(filename)
		}

		res, err := h.Usecase.CreateHomework(ctx, &req)
		if err != nil {
			fmt.Println(err)
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		c.JSON(http.StatusCreated, gin.H{"data": res})
	}
}
func (h *Handler) Get(ctx context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
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
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		res, err := h.Usecase.GetContent(ctx, &req)
		if err == message.ErrUnnecessaryUpdate {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": res})
	}
}

func (h *Handler) UpdateBulletin(ctx context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			uri domain.CourseBulletinUriRequest
			req domain.UpdateCourseBulletinRequest
		)
		if err := c.ShouldBindUri(&uri); err != nil {
			fmt.Println(err)
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			fmt.Println(err)
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		req.SetupUri(uri)
		fmt.Println(req)

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
		var (
			uri domain.CourseSlideUriRequest
			req domain.UpdateCourseSlideRequest
			dst string
		)
		if err := c.ShouldBindUri(&uri); err != nil {
			fmt.Println(err)
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		if err := c.ShouldBind(&req); err != nil {
			fmt.Println(err)
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		req.SetupUri(&uri)

		if req.File != nil && req.File.Size != 0 {
			filename := uuid.New().String()
			dst = fmt.Sprintf(h.conf.Server.SlidePathFormat, req.TeacherDomain, filename)
			if err := c.SaveUploadedFile(req.File, dst); err != nil {
				// todo: log error
				fmt.Println(err)
				c.AbortWithStatus(http.StatusBadRequest)
				return
			}
			req.SetFilename(filename)
		}
		res, err := h.Usecase.UpdateSlide(ctx, &req)
		if err != nil {
			// Remove the file cause some errors
			if req.File != nil && req.File.Size != 0 {
				os.Remove(dst)
			}
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": res})
	}
}

func (h *Handler) UpdateHomework(ctx context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			uri domain.CourseHomeworkUriRequest
			req domain.UpdateCourseHomeworkRequest
			dst string
		)
		if err := c.ShouldBindUri(&uri); err != nil {
			fmt.Println(err)
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		if err := c.ShouldBind(&req); err != nil {
			fmt.Println(err)
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		req.SetupUri(&uri)

		if req.File != nil && req.File.Size != 0 {
			filename := uuid.New().String()
			dst = fmt.Sprintf(h.conf.Server.HomeworkPathFormat, req.TeacherDomain, filename)
			if err := c.SaveUploadedFile(req.File, dst); err != nil {
				// todo: log error
				fmt.Println(err)
				c.AbortWithStatus(http.StatusInternalServerError)
				return
			}
			req.SetFilename(filename)
		}
		res, err := h.Usecase.UpdateHomework(ctx, &req)
		if err != nil {
			// Remove the file cause some errors
			if req.File != nil && req.File.Size != 0 {
				os.Remove(dst)
			}
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
			fmt.Println(err)
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
			fmt.Println(err)
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
