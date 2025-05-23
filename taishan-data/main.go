package main

import (
	"context"
	"data/config"
	"data/internal"
	"data/internal/biz/log"
	"data/logic"
	"data/router"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"os/signal"
	"time"
)

var configFile string

func main() {

	flag.StringVar(&configFile, "f", "./conf.yml", "use config file")
	flag.Parse()

	config.InitConfig(configFile)

	internal.InitProjects()
	logic.InitJob()

	// 创建 Gin 引擎实例
	engine := gin.Default()
	router.RegisterRouter(engine)

	// 创建 HTTP 服务器
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", config.Conf.Http.Port),
		Handler: engine,
	}

	// 启动 HTTP 服务器
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Logger.Fatalf("服务启动失败: %v", err)
		}
	}()

	// 优雅退出
	gracefulExit(server)
}

// 释放连接池资源，优雅退出
func gracefulExit(server *http.Server) {
	// 等待中断信号，然后优雅地关闭服务器
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Logger.Info("服务关闭中...")

	// 创建一个 5 秒的上下文，用于等待请求处理完成
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 关闭服务器，等待现有的请求处理完成
	if err := server.Shutdown(ctx); err != nil {
		log.Logger.Fatalf("服务关闭失败: %v", err)
	}

	log.Logger.Info("服务已经优雅的关闭~~~")
}
