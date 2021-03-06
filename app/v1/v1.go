package v1

import (
	"context"

	pagedelivery "teacher-site/app/v1/page/delivery"
	pagerepo "teacher-site/app/v1/page/repository"
	pageusecase "teacher-site/app/v1/page/usecase"

	infodelivery "teacher-site/app/v1/info/delivery"
	inforepo "teacher-site/app/v1/info/repository"
	infousecase "teacher-site/app/v1/info/usecase"

	coursedelivery "teacher-site/app/v1/course/delivery"
	courserepo "teacher-site/app/v1/course/repository"
	courseusecase "teacher-site/app/v1/course/usecase"

	authdelivery "teacher-site/app/v1/auth/delivery"
	authrepo "teacher-site/app/v1/auth/repository"
	authusecase "teacher-site/app/v1/auth/usecase"

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
	{
		rAuth := v1.Group("/auth")
		authRepoDb := authrepo.NewDbRepository(db, conf.DB)
		authRepoCache := authrepo.NewCacheRepository(c, conf.Redis)
		authUsecase := authusecase.NewUsecase(authRepoDb, authRepoCache, conf, logger)
		authdelivery.NewHandler(ctx, rAuth, authUsecase, conf)

		rTeacher := v1.Group("/:teacherDomain", mw.CheckTeacherDomain())
		{

			rInfo := rTeacher.Group("/info")
			repoDbInfo := inforepo.NewDbRepository(db, conf.DB)
			repoCacheInfo := inforepo.NewCacheRepository(c, conf.Redis)
			UsecaseInfo := infousecase.NewUsecase(repoDbInfo, repoCacheInfo, conf, logger)
			infodelivery.NewHandler(ctx, rInfo, UsecaseInfo, conf)

			rCourse := rTeacher.Group("/course")
			repoDbCourse := courserepo.NewDbRepository(db, conf.DB)
			repoCacheCourse := courserepo.NewCacheRepository(c, conf.Redis)
			UsecaseCourse := courseusecase.NewUsecase(repoDbCourse, repoCacheCourse, conf, logger)
			coursedelivery.NewHandler(ctx, rCourse, UsecaseCourse, conf)
		}
	}
}
