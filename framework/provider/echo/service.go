package echo

import (
	"fmt"

	"github.com/demian/webdesign/framework/contract"
)

type EchoService struct {
	contract.EchoIService
	Msg  string
	Name string
}

func (e *EchoService) Echo() string {
	return fmt.Sprint(e.Msg, " ", e.Name)
}

func NewEchoService(params ...interface{}) (interface{}, error) {
	msg := params[0].(string)
	name := params[1].(string)

	return &EchoService{
		Msg:  msg,
		Name: name,
	}, nil
}
