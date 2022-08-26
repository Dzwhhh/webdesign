package demo

// demo服务的凭证
const Key = "wd:demo"

// demo服务结构
type Service interface {
	GetGroot() Groot
}

// demo服务定义的数据结构
type Groot struct {
	Name string
}
