package kernel

import (
	"github.com/demian/webdesign/framework"
	"github.com/demian/webdesign/framework/contract"
	"github.com/demian/webdesign/framework/gin"
)

type KernelServiceProvider struct {
	framework.ServiceProvider
	HttpEngine *gin.Engine
}

func (k *KernelServiceProvider) Name() string {
	return contract.KernelKey
}

func (k *KernelServiceProvider) IsDefer() bool {
	return false
}

func (k *KernelServiceProvider) Boot(c framework.Container) error {
	if k.HttpEngine == nil {
		k.HttpEngine = gin.Default()
	}
	k.HttpEngine.SetContainer(c)
	return nil
}

func (k *KernelServiceProvider) Params(c framework.Container) []interface{} {
	return []interface{}{k.HttpEngine}
}

func (k *KernelServiceProvider) Register(c framework.Container) framework.NewInstance {
	return NewKernelService
}
