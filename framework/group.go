package framework

import "errors"

type IGroup interface {
	Get(string, ControllerHandler) error
	Post(string, ControllerHandler) error
	Put(string, ControllerHandler) error
	Delete(string, ControllerHandler) error
}

type Group struct {
	prefix string
	core   *Core
}

func newGroup(prefix string, core *Core) *Group {
	return &Group{
		prefix: prefix,
		core:   core,
	}
}

func (g *Group) Get(uri string, handler ControllerHandler) error {
	uri = g.prefix + uri
	if err := g.core.Get(uri, handler); err != nil {
		return errors.New("register failed: " + uri)
	}
	return nil
}

func (g *Group) Post(uri string, handler ControllerHandler) error {
	uri = g.prefix + uri
	if err := g.core.Post(uri, handler); err != nil {
		return errors.New("register failed: " + uri)
	}
	return nil
}

func (g *Group) Put(uri string, handler ControllerHandler) error {
	uri = g.prefix + uri
	if err := g.core.Put(uri, handler); err != nil {
		return errors.New("register failed: " + uri)
	}
	return nil
}

func (g *Group) Delete(uri string, handler ControllerHandler) error {
	uri = g.prefix + uri
	if err := g.core.Delete(uri, handler); err != nil {
		return errors.New("register failed: " + uri)
	}
	return nil
}
