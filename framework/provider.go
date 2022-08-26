package framework

// 创建服务实例方法
type NewInstance func(...interface{}) (interface{}, error)

type ServiceProvider interface {
	// 服务提供者的凭证
	Name() string

	// 服务容器中注册实例化服务的方法
	Register(Container) NewInstance

	// 调用实例化服务的准备工作
	Boot(Container) error

	// 定义传递给NewInstance的参数
	Params(Container) []interface{}

	// 是否需要延迟实例化
	IsDefer() bool
}
