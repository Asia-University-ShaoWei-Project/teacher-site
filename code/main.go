package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"teacher-site/app"
	"teacher-site/config"
	"teacher-site/pkg/database"
	_log "teacher-site/pkg/log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

const (
	_staticRelativePath = "/static"
	_staticRoot         = "./static"
	_templatePath       = "templates/*"
	_dbFilePath         = "./database"
	_serverPort         = ":80"
)

func main() {
	ctx := context.Background()
	// todo: tmp
	conf := config.New()
	// cacheConf := model.NewMockCacheConfig()
	// cache := cache.NewCache(cacheConf)
	redis := database.NewRedis(conf.Redis)
	logger := _log.NewLogrus(ctx)
	// db := database.NewSqlite(_dbFilePath, logger)
	db := database.NewDB(".", conf.DB)

	// conf := model.NewMockServiceConfig()
	// srv := service.NewService(db, cache, logger, conf)
	// todo: use os.Getenv()

	r := gin.Default()
	r.Use(cors.Default())
	r.Static(_staticRelativePath, _staticRoot)
	r.LoadHTMLGlob(_templatePath)
	// route.SetupRoute(ctx, r, srv)
	app.SetupRoute(ctx, r, db, redis, logger, conf)
	r.Run(_serverPort)
	//? graceful shutdown: https://blog.wu-boy.com/2020/02/what-is-graceful-shutdown-in-golang/
	server := &http.Server{
		Addr:    _serverPort,
		Handler: r,
	}
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatalf("listen: %s\n", err)
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		logger.Fatal("Server Shutdown: ", err)
	}
	logger.Println("Server exiting")
}
