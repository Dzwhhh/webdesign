package kernel

import (
	"net/http"

	"github.com/demian/webdesign/framework/contract"
	"github.com/demian/webdesign/framework/gin"
)

type KernelService struct {
	contract.Kernel
	engine *gin.Engine
}

func NewKernelService(params ...interface{}) (interface{}, error) {
	httpEngine := params[0].(*gin.Engine)
	return &KernelService{
		engine: httpEngine,
	}, nil
}

func (k *KernelService) HttpEngine() http.Handler {
	return k.engine
}
