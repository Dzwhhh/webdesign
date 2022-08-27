package app

import (
	"github.com/demian/webdesign/framework"
	"github.com/demian/webdesign/framework/contract"
)

// 服务提供者
type WdAppProvider struct {
	framework.ServiceProvider
	BaseFolder string
}

func (wd *WdAppProvider) Name() string {
	return contract.AppKey
}

func (wd *WdAppProvider) Boot(c framework.Container) error {
	return nil
}

func (wd *WdAppProvider) IsDefer() bool {
	return false
}

func (wd *WdAppProvider) Params(c framework.Container) []interface{} {
	return []interface{}{c, wd.BaseFolder}
}

func (wd *WdAppProvider) Register(c framework.Container) framework.NewInstance {
	return NewAppSevice
}
