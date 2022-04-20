package v1

import (
	"context"
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

func SetupRoute(ctx context.Context, r *gin.RouterGroup, db *gorm.DB, c *redis.Client, logger *log.Logger, conf *config.Config) {

	teacher := r.Group("/v1/:teacher_domain", mw.CheckTeacherDomain())
	rInfo := teacher.Group("/info")
	repoDbInfo := inforepo.NewInfoRepository(db, conf.DB)
	repoCacheInfo := inforepo.NewCacheRepository(c, conf.Redis)
	infoUsecase := infousecase.NewInfoUsecase(repoDbInfo, repoCacheInfo, logger)
	infodelivery.NewInfoHandler(ctx, rInfo, infoUsecase, conf)
	// rCourse := teacher.Group("/course")
}
