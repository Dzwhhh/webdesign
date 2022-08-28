package command

import (
	"fmt"

	"github.com/demian/webdesign/framework/cobra"
	"github.com/demian/webdesign/framework/contract"
)

var DirCommand = &cobra.Command{
	Use:   "dir",
	Short: "展示framework的Base路径",
	RunE: func(c *cobra.Command, args []string) error {
		container := c.GetContainer()
		appService := container.MustMake(contract.AppKey).(contract.App)
		fmt.Println("Base Dir: ", appService.BaseFolder())
		return nil
	},
}
