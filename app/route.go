package app

import (
	"context"
	"net/http"
	v1 "teacher-site/app/v1"
	"teacher-site/config"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

const templateIndex = "index.html"

func SetupRoute(ctx context.Context, r *gin.Engine, db *gorm.DB, c *redis.Client, logger *log.Logger, conf *config.Config) {
	r.GET("/:teacher_domain", func(c *gin.Context) {
		c.HTML(http.StatusOK, templateIndex, gin.H{})
	})
	api := r.Group("/api")
	v1.SetupRoute(ctx, api, db, c, logger, conf)
	// v2.SetupRoute(ctx, api)
}