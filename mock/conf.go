package mock

import (
	"teacher-site/config"
	"teacher-site/pkg/log"
)

var (
	Conf = config.New()
	Log  = log.NewLogrus(Ctx)
)
