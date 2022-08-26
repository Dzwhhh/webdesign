package framework

import (
	"errors"
	"fmt"
	"sync"
)

type Container interface {
	// 绑定服务提供者
	Bind(provider ServiceProvider) error

	// 服务是否已绑定
	IsBind(key string) bool

	// 根据关键字凭证获取服务实例
	Make(key string) (interface{}, error)

	// 确定服务已绑定的情况下获取服务实例
	MustMake(key string) interface{}

	// 根据不同参数获取不同服务实例
	MakeNew(key string, params []interface{}) (interface{}, error)
}

type WdContainer struct {
	Container                            // 强制要求实现container接口
	providers map[string]ServiceProvider // 存储注册的服务提供者
	instances map[string]interface{}     // 存储服务实例
	lock      *sync.RWMutex
}

func NewWdContainer() *WdContainer {
	return &WdContainer{
		providers: map[string]ServiceProvider{},
		instances: map[string]interface{}{},
		lock:      &sync.RWMutex{},
	}
}

func (wd *WdContainer) PrintProviders() []string {
	if len(wd.providers) == 0 {
		return nil
	}
	ret := []string{}
	for _, provider := range wd.providers {
		name := provider.Name()
		line := fmt.Sprint(name)
		ret = append(ret, line)
	}
	return ret
}

func (wd *WdContainer) Bind(provider ServiceProvider) error {
	wd.lock.Lock() // 写锁定
	defer wd.lock.Unlock()

	key := provider.Name()
	wd.providers[key] = provider

	if !provider.IsDefer() {
		if err := provider.Boot(wd); err != nil {
			return err
		}
		// 实例化方法
		params := provider.Params(wd)
		method := provider.Register(wd)
		instance, err := method(params...)
		if err != nil {
			return errors.New(err.Error())
		}
		wd.instances[key] = instance
	}
	return nil
}

func (wd *WdContainer) IsBind(key string) bool {
	return wd.findServiceProvider(key) == nil
}

func (wd *WdContainer) findServiceProvider(key string) ServiceProvider {
	wd.lock.RLock()
	defer wd.lock.RUnlock()

	if provider, ok := wd.providers[key]; ok {
		return provider
	}
	return nil

}

func (wd *WdContainer) Make(key string) (interface{}, error) {
	return wd.make(key, nil, false)
}

func (wd *WdContainer) MakeNew(key string, params []interface{}) (interface{}, error) {
	return wd.make(key, params, true)
}

func (wd *WdContainer) MustMake(key string) interface{} {
	instance, err := wd.Make(key)
	if err != nil {
		panic(err)
	}
	return instance
}

func (wd *WdContainer) make(key string, params []interface{}, forceNew bool) (interface{}, error) {
	wd.lock.RLock()
	defer wd.lock.RUnlock()

	provider := wd.findServiceProvider(key)
	if provider == nil {
		return nil, errors.New("contract " + key + " have not register")
	}
	// 需要强制重新实例化
	if forceNew {
		return wd.newInstance(provider, params)
	}

	// 实例已存在
	if ins, ok := wd.instances[key]; ok {
		return ins, nil
	}

	// 实例不存在
	inst, err := wd.newInstance(provider, nil)
	if err != nil {
		return nil, err
	}
	wd.instances[key] = inst
	return inst, nil
}

func (wd *WdContainer) newInstance(provider ServiceProvider, params []interface{}) (interface{}, error) {
	if err := provider.Boot(wd); err != nil {
		return nil, err
	}
	if params == nil {
		params = provider.Params(wd)
	}
	// 实例化方法
	method := provider.Register(wd)
	ins, err := method(params...)
	if err != nil {
		return nil, err
	}
	return ins, nil
}
