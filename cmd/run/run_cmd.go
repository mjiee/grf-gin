package run

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var (
	confFile string

	RunCmd = &cobra.Command{
		Use:     "run",
		Short:   "scaffold-gin run",
		Long:    "run子命令用于运行项目",
		Example: "scaffold-gin run -c ./conf/default.yaml",
		RunE: func(cmd *cobra.Command, args []string) error {
			return run()
			// return nil
		},
	}
)

func init() {
	RunCmd.PersistentFlags().StringVarP(&confFile, "config", "c", "./conf/default.yaml", "指定项目配置文件")
}

func run() error {
	app, cleanup, err := initApp(confFile)
	if err != nil {
		app.zlog.Error("初始化app失败...", zap.Error(err))
		return err
	}

	defer cleanup()

	go func() {
		if err := app.httpSrv.ListenAndServe(); err != http.ErrServerClosed {
			app.zlog.Error("启动http服务失败...", zap.Error(err))
		}
	}()

	app.zlog.Info("app服务启动成", zap.String("addr:", app.addr))

	// 等待进程中断信号, 等待5s后优雅关闭服务
	quit := make(chan os.Signal)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	app.zlog.Info("开始关闭服务...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := app.httpSrv.Shutdown(ctx); err != nil {
		app.zlog.Error("关闭服务失败...", zap.Error(err))
	}

	select {
	case <-ctx.Done():
		app.zlog.Error("关闭服务超时...")
	}

	app.zlog.Info("成功关闭服务....")
	return nil
}
