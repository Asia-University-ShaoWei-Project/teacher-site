package v1

import (
	"context"

	pagedelivery "teacher-site/app/v1/page/delivery"
	pagerepo "teacher-site/app/v1/page/repository"
	pageusecase "teacher-site/app/v1/page/usecase"

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

func SetupRoute(ctx context.Context, r *gin.Engine, db *gorm.DB, c *redis.Client, conf *config.Config, logger *log.Logger) {
	rPage := r.Group("/")
	repoDbPage := pagerepo.NewDbRepository(db, conf.DB)
	repoCachePage := pagerepo.NewCacheRepository(c, conf.Redis)
	usecasePage := pageusecase.NewUsecase(repoDbPage, repoCachePage, conf, logger)
	pagedelivery.NewHandler(ctx, rPage, usecasePage, conf)

	v1 := r.Group("/api/v1")
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

	rTeacher := v1.Group("/:teacher_domain", mw.CheckTeacherDomain())

	rInfo := rTeacher.Group("/info")
	repoDbInfo := inforepo.NewDbRepository(db, conf.DB)
	repoCacheInfo := inforepo.NewCacheRepository(c, conf.Redis)
	UsecaseInfo := infousecase.NewUsecase(repoDbInfo, repoCacheInfo, conf, logger)
	infodelivery.NewHandler(ctx, rInfo, UsecaseInfo, conf)

	rAuth := v1.Group("/auth")
	authRepoDb := authrepo.NewDbRepository(db, conf.DB)
	authRepoCache := authrepo.NewCacheRepository(c, conf.Redis)
	authUsecase := authusecase.NewUsecase(authRepoDb, authRepoCache, conf, logger)
	authdelivery.NewHandler(ctx, rAuth, authUsecase, conf)

	// rCourse := teacher.Group("/course")
}
