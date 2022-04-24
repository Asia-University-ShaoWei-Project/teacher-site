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
	v1.SetupRoute(ctx, r, db, c, conf, logger)
	// v2.SetupRoute(ctx, api)
}
