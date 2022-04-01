// TODO: request get from redis

// TODO: store data to redis

// TODO: request get from db
package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"teacher-site/cache"
	"teacher-site/database"
	"teacher-site/log"
	"teacher-site/model"
	"teacher-site/route"
	"teacher-site/service"
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
	cache := cache.NewCache()
	db := database.NewSqlite(_dbFilePath)
	// logger := log.NewLogzap(ctx)
	// defer logger.Sync()
	logger := log.NewLogrus(ctx)
	srv := service.NewService(db, cache, logger)
	// todo: use os.Getenv()
	cfg := &model.Config{
		JWTSecure:      []byte(`secure`),
		PasswordSecure: `secure`,
		HashCost:       10,
		// jwtSecure: []byte(os.Getenv(`secure`)),
	}
	r := gin.Default()
	r.Use(cors.Default())
	r.Static(_staticRelativePath, _staticRoot)
	r.LoadHTMLGlob(_templatePath)
	route.SetupRoute(ctx, r, srv, cfg)
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
