package command

import "github.com/demian/webdesign/framework/cobra"

func AddKernelCommand(root *cobra.Command) {
	// 绑定start命令
	root.AddCommand(initAppCommand())
	// 绑定dir命令
	root.AddCommand(DirCommand)
}
