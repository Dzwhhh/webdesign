package console

import (
	"github.com/demian/webdesign/framework"
	"github.com/demian/webdesign/framework/cobra"
	"github.com/demian/webdesign/framework/command"
)

func RunCommand(container framework.Container) error {
	// 根 command
	var rootCmd = &cobra.Command{
		// 根命令关键字
		Use:   "wd",
		Short: "wd 命令",
		Long:  "框架提供的命令行工具，使用这个命令行工具能很方便执行框架自带命令，也能很方便编写业务命令",
		// 根命令执行函数
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.InitDefaultHelpFlag()
			return cmd.Help()
		},
		CompletionOptions: cobra.CompletionOptions{DisableDefaultCmd: true},
	}
	// 为根command设置服务容器
	rootCmd.SetContainer(container)
	// 绑定框架的命令
	command.AddKernelCommand(rootCmd)
	// 绑定业务的命令
	AddAppCommand(rootCmd)
	// 执行RootCommand
	return rootCmd.Execute()
}

func AddAppCommand(cmd *cobra.Command) error {
	return nil
}
