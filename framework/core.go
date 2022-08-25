package framework

import (
	"log"
	"net/http"
	"strings"
)

type Core struct {
	router      map[string]*Trie
	middlewares []ControllerHandler
}

func NewCore() *Core {
	router := map[string]*Trie{}
	router["GET"] = NewTrie()
	router["POST"] = NewTrie()
	router["PUT"] = NewTrie()
	router["DELETE"] = NewTrie()
	return &Core{
		router: router,
	}
}

// 路由注册 Get方法
func (c *Core) Get(uri string, handlers ...ControllerHandler) error {
	allHandlers := append(c.middlewares, handlers...)
	if err := c.router["GET"].AddRouter(uri, allHandlers); err != nil {
		log.Fatal("add router error", err)
	}
	return nil
}

// 路由注册 Post方法
func (c *Core) Post(uri string, handlers ...ControllerHandler) error {
	allHandlers := append(c.middlewares, handlers...)
	if err := c.router["POST"].AddRouter(uri, allHandlers); err != nil {
		log.Fatal("add router error", err)
	}
	return nil
}

// 路由注册 Put方法
func (c *Core) Put(uri string, handlers ...ControllerHandler) error {
	allHandlers := append(c.middlewares, handlers...)
	if err := c.router["PUT"].AddRouter(uri, allHandlers); err != nil {
		log.Fatal("add router error", err)
	}
	return nil
}

// 路由注册 Delete方法
func (c *Core) Delete(uri string, handlers ...ControllerHandler) error {
	allHandlers := append(c.middlewares, handlers...)
	if err := c.router["DELETE"].AddRouter(uri, allHandlers); err != nil {
		log.Fatal("add router error", err)
	}
	return nil
}

// 路由注册 公共前缀
func (c *Core) Group(prefix string) IGroup {
	return newGroup(prefix, c)
}

// 注册中间件
func (c *Core) Use(middlewares ...ControllerHandler) {
	c.middlewares = append(c.middlewares, middlewares...)
}

func (c *Core) FindRouteByRequest(r *http.Request) []ControllerHandler {
	// uri 和 method 全部转换为大写
	upperUri := strings.ToUpper(r.URL.Path)
	upperMethod := strings.ToUpper(r.Method)

	// 查找第一层map
	if methodHandlers, ok := c.router[upperMethod]; ok {
		return methodHandlers.FindHandler(upperUri)
	}
	return nil
}

func (c *Core) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	ctx := NewContext(r, w)

	// 寻找路由
	handlers := c.FindRouteByRequest(r)

	if handlers == nil {
		ctx.Json(404, "Not Found")
		return
	}

	// 设置context中的handlers字段
	ctx.SetHandlers(handlers)

	// 调用路由函数
	if err := ctx.Next(); err != nil {
		ctx.Json(500, "inner error")
	}
}

func (c *Core) Listen(port string) {
	if port == "" {
		port = "8080"
	}
	if !strings.HasPrefix(port, ":") {
		port = ":" + port
	}
	server := http.Server{
		Addr:    port,
		Handler: c,
	}
	server.ListenAndServe()
}
