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
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"

	// heroku
	_ "github.com/heroku/x/hmetrics/onload"
)

func main() {
	ctx := context.Background()
	// todo: use os.Getenv() for config
	conf := config.New()
	redis := database.NewRedis(conf.Redis)
	logger := _log.NewLogrus(ctx)
	// todo: use postgres
	db := database.NewDB("./pkg/database", conf.DB)
	cookieStore := cookie.NewStore(conf.Secure.SessionSecret)

	// todo: release(mode, migrate, config(port))
	// migrate.Setup(db)
	// gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.MaxMultipartMemory = conf.Server.MaxMultipartMemory
	r.Use(sessions.Sessions("session", cookieStore))
	r.Use(cors.Default())
	r.Static(conf.Server.StaticRelativePath, conf.Server.StaticRootPath)
	r.LoadHTMLGlob(conf.Server.TemplatePath)
	app.SetupRoute(ctx, r, db, redis, logger, conf)
	//? Graceful shutdown: https://blog.wu-boy.com/2020/02/what-is-graceful-shutdown-in-golang/
	server := &http.Server{
		Addr:    conf.Server.Addr,
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
