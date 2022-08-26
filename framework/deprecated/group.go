package framework

import (
	"errors"
)

type IGroup interface {
	Get(string, ...ControllerHandler) error
	Post(string, ...ControllerHandler) error
	Put(string, ...ControllerHandler) error
	Delete(string, ...ControllerHandler) error

	// 实现嵌套group
	Group(string) IGroup

	// 嵌套中间件
	Use(middlewares ...ControllerHandler)
}

type Group struct {
	prefix      string
	core        *Core
	parent      *Group // 指向上一个Group
	middlewares []ControllerHandler
}

func newGroup(prefix string, core *Core) *Group {
	return &Group{
		prefix:      prefix,
		core:        core,
		parent:      nil,
		middlewares: []ControllerHandler{},
	}
}

func (g *Group) getAbsolutePrefix() string {
	if g.parent == nil {
		return g.prefix
	}
	return g.parent.getAbsolutePrefix() + g.prefix
}

func (g *Group) Get(uri string, handlers ...ControllerHandler) error {
	uri = g.getAbsolutePrefix() + uri
	allHandlers := append(g.getMiddlewares(), handlers...)
	if err := g.core.Get(uri, allHandlers...); err != nil {
		return errors.New("register failed: " + uri)
	}
	return nil
}

func (g *Group) Post(uri string, handlers ...ControllerHandler) error {
	uri = g.getAbsolutePrefix() + uri
	allHandlers := append(g.getMiddlewares(), handlers...)
	if err := g.core.Post(uri, allHandlers...); err != nil {
		return errors.New("register failed: " + uri)
	}
	return nil
}

func (g *Group) Put(uri string, handlers ...ControllerHandler) error {
	uri = g.getAbsolutePrefix() + uri
	allHandlers := append(g.getMiddlewares(), handlers...)
	if err := g.core.Put(uri, allHandlers...); err != nil {
		return errors.New("register failed: " + uri)
	}
	return nil
}

func (g *Group) Delete(uri string, handlers ...ControllerHandler) error {
	uri = g.getAbsolutePrefix() + uri
	allHandlers := append(g.getMiddlewares(), handlers...)
	if err := g.core.Delete(uri, allHandlers...); err != nil {
		return errors.New("register failed: " + uri)
	}
	return nil
}

// 注册中间件
func (g *Group) Use(middlewares ...ControllerHandler) {
	g.middlewares = append(g.middlewares, middlewares...)
}

func (g *Group) getMiddlewares() []ControllerHandler {
	if g.parent == nil {
		return g.middlewares
	}
	return append(g.parent.getMiddlewares(), g.middlewares...)
}

// 实现Group方法
func (g *Group) Group(uri string) IGroup {
	cgroup := newGroup(uri, g.core)
	cgroup.parent = g
	return cgroup
}
