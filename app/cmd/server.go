package cmd

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/spf13/cast"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"

	"sample/bootstrap"
	"sample/pkg/config"
	"sample/pkg/console"
	"sample/pkg/logger"
)

func init() {
	rootCmd.AddCommand(ServerCommand)
}

// ServerCommand 运行 web 服务
var ServerCommand = &cobra.Command{
	Use:   "serve",
	Short: "启动web服务",
	Run:   runWeb,
	Args:  cobra.NoArgs,
}

func init() {
	ServerCommand.Flags().StringP("port", "p", "", "服务启动端口，默认 8080")
	ServerCommand.Flags().StringP("shutdown_time", "s", "", "关闭服务器超时时间，默认 5，单位：s")
}

func runWeb(cmd *cobra.Command, args []string) {
	// gin.SetMode(gin.ReleaseMode)
	// gin 实例
	engine := gin.New()

	// 初始化路由绑定
	bootstrap.SetupRoute(engine)

	// 解析端口
	port, _ := cmd.Flags().GetString("port")
	if port == "" {
		port = config.Get("app.port")
	}
	if port == "" {
		port = "8080"
	}
	if !strings.HasPrefix(port, ":") {
		port = ":" + port
	}

	// 解析超时时间
	t, _ := cmd.Flags().GetString("shutdown_time")
	dt := cast.ToInt64(t)
	if dt <= 0 {
		dt = 5
	}

	serve := http.Server{
		Addr:    port,
		Handler: engine,
	}

	go func() {
		err := serve.ListenAndServe()
		if err != nil {
			logger.ErrorString("CMD", "start serve", err.Error())
			console.Exit("启动服务求失败，错误:" + err.Error())
		}
	}()

	quit := make(chan os.Signal, 1)
	// kill 默认会发送 syscall.SIGTERM 信号
	// kill -2 发送 syscall.SIGINT 信号，我们常用的Ctrl+C就是触发系统SIGINT信号
	// kill -9 发送 syscall.SIGKILL 信号，但是不能被捕获，所以不需要添加它
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	go func() {
		select {
		case <-quit:
			logger.ErrorString("CMD", "stop serve", "force")
			console.Exit("强制退出")
		}
	}()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(dt))
	defer cancel()
	err := serve.Shutdown(ctx)
	if err != nil {
		logger.ErrorString("CMD", "stop serve", err.Error())
		console.Exit("关闭服务器失败，错误：" + err.Error())
	}
	console.Warning("服务退出成功")
}
