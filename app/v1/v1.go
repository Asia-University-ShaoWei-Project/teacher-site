package v1

import (
	"context"

	authdelivery "teacher-site/app/v1/auth/delivery"
	authrepo "teacher-site/app/v1/auth/repository"
	authusecase "teacher-site/app/v1/auth/usecase"
	infodelivery "teacher-site/app/v1/info/delivery"
	inforepo "teacher-site/app/v1/info/repository"
	infousecase "teacher-site/app/v1/info/usecase"
	"teacher-site/config"
	mw "teacher-site/middleware"

	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

func SetupRoute(ctx context.Context, r *gin.RouterGroup, db *gorm.DB, c *redis.Client, conf *config.Config, logger *log.Logger) {
	v1 := r.Group("/v1")
	v1.GET("/test", func(c *gin.Context) {
		isTeacher := mw.IsTeacher(ctx, c, conf.Jwt.Secret)
		if isTeacher {
			log.Info("is user")
			c.AbortWithStatus(200)
			return
		}
		log.Info("not user")
		c.Status(400)
	})

	teacher := v1.Group("/:teacher_domain", mw.CheckTeacherDomain())

	rInfo := teacher.Group("/info")
	repoDbInfo := inforepo.NewInfoRepository(db, conf.DB)
	repoCacheInfo := inforepo.NewCacheRepository(c, conf.Redis)
	UsecaseInfo := infousecase.NewInfoUsecase(repoDbInfo, repoCacheInfo, conf, logger)
	infodelivery.NewInfoHandler(ctx, rInfo, UsecaseInfo, conf)

	rAuth := v1.Group("/auth")
	authRepoDb := authrepo.NewDbRepository(db, conf.DB)
	authRepoCache := authrepo.NewCacheRepository(c, conf.Redis)
	authUsecase := authusecase.NewUsecase(authRepoDb, authRepoCache, conf, logger)
	authdelivery.NewHandler(ctx, rAuth, authUsecase, conf)

	// rCourse := teacher.Group("/course")
}
