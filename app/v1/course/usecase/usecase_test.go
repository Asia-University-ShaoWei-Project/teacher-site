package usecase

import (
	"teacher-site/mock"
	mockRepo "teacher-site/mock/course/repository"
)

var (
	ctx       = mock.Ctx
	dbRepo    = mockRepo.NewDbRepository()
	cacheRepo = mockRepo.NewCacheRepository()
	usecase   = NewUsecase(dbRepo, cacheRepo, mock.Conf, mock.Log)
)
