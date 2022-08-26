package echo

import "fmt"

type EchoService struct {
	Service
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
