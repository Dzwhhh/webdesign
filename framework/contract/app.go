package contract

const AppKey = "wd:app"

type App interface {
	Version() string          // 框架版本
	BaseFolder() string       // 项目根目录
	ConfigFolder() string     // 配置路径
	LogFolder() string        // 日志路径
	ProviderFolder() string   // 服务提供者路径
	MiddlewareFolder() string // 中间件路径
	CommandFolder() string    // 存储命令行命令
	RuntimeFolder() string    // 存储运行时信息
	TestFolder() string       // 测试用例路径
}
