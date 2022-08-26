package demo

import "github.com/demian/webdesign/framework"

type DemoService struct {
	// 实现接口
	Service

	// 参数
	c framework.Container
}

func (d *DemoService) GetGroot() Groot {
	return Groot{
		Name: "I am Groot",
	}
}

func NewDemoService(params ...interface{}) (interface{}, error) {
	// 将参数展开
	c := params[0].(framework.Container)
	return &DemoService{c: c}, nil
}
