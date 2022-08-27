package contract

// demo服务的凭证
const DemoKey = "wd:demo"

// demo服务结构
type DemoIService interface {
	GetGroot() Groot
}

// demo服务定义的数据结构
type Groot struct {
	Name string
}
