package echo

import (
	"fmt"

	"github.com/demian/webdesign/framework"
	"github.com/demian/webdesign/framework/contract"
)

type EchoServiceProvider struct {
	framework.ServiceProvider
}

func (e *EchoServiceProvider) Name() string {
	return contract.EchoKey
}

func (e *EchoServiceProvider) IsDefer() bool {
	return false
}

func (e *EchoServiceProvider) Params(c framework.Container) []interface{} {
	msg := "Hello"
	name := "World"
	return []interface{}{msg, name}
}

func (e *EchoServiceProvider) Boot(c framework.Container) error {
	fmt.Print("echo service boot")
	return nil
}

func (e *EchoServiceProvider) Register(c framework.Container) framework.NewInstance {
	return NewEchoService
}
