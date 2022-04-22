package app

import (
	"context"
	v1 "teacher-site/app/v1"
	"teacher-site/config"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func SetupRoute(ctx context.Context, r *gin.Engine, db *gorm.DB, c *redis.Client, logger *log.Logger, conf *config.Config) {
	template := NewTemplate()
	r.Any("/", template.Home(ctx))
	r.GET("/:teacher_domain", template.TeacherSite(ctx))
	r.GET("/login", template.Login(ctx, conf.Jwt))

	api := r.Group("/api")
	v1.SetupRoute(ctx, api, db, c, conf, logger)
	// v2.SetupRoute(ctx, api)
}
