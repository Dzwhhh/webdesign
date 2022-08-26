package demo

import (
	"fmt"

	"github.com/demian/webdesign/framework"
)

type DemoServiceProvider struct {
	framework.ServiceProvider
}

func (d *DemoServiceProvider) Name() string {
	return Key
}

func (d *DemoServiceProvider) Register(c framework.Container) framework.NewInstance {
	return NewDemoService
}

func (d *DemoServiceProvider) Boot(c framework.Container) error {
	fmt.Println("demo service boot")
	return nil
}

func (d *DemoServiceProvider) Params(c framework.Container) []interface{} {
	return []interface{}{c}
}

func (d *DemoServiceProvider) IsDefer() bool {
	return true
}
