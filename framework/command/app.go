package command

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/demian/webdesign/framework/cobra"
	"github.com/demian/webdesign/framework/contract"
)

var appCommand = &cobra.Command{
	Use:   "app",
	Short: "业务应用控制命令",
	RunE: func(c *cobra.Command, args []string) error {
		c.Help()
		return nil
	},
}

func initAppCommand() *cobra.Command {
	appCommand.AddCommand(appStartCommand)
	return appCommand
}

var appStartCommand = &cobra.Command{
	Use:   "start",
	Short: "服务启动命令",
	RunE: func(c *cobra.Command, args []string) error {
		// 从command中获取容器
		container := c.GetContainer()
		// 从服务容器中获取kernel服务实例
		kernelService := container.MustMake(contract.KernelKey).(contract.Kernel)
		// 从kernel服务实例中获取引擎
		core := kernelService.HttpEngine()

		// 创建http服务
		server := &http.Server{
			Addr:    ":8080",
			Handler: core,
		}

		go func() {
			server.ListenAndServe()
		}()

		// 平滑退出
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
		<-quit // 阻塞 等待退出信号

		if err := server.Shutdown(context.Background()); err != nil {
			log.Fatal("Server Shutdown Error:", err)
		}
		return nil
	},
}
