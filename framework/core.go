package framework

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
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

func (c *Core) FindRouteNodeByRequest(r *http.Request) *node {
	//uri 和 method 全部转换为大写
	upperUri := strings.ToUpper(r.URL.Path)
	upperMethod := strings.ToUpper(r.Method)

	// 查找第一层map
	if methodHandlers, ok := c.router[upperMethod]; ok {
		return methodHandlers.root.matchNode(upperUri)
	}
	return nil
}

func (c *Core) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	ctx := NewContext(r, w)

	// 寻找路由
	node := c.FindRouteNodeByRequest(r)
	if node == nil {
		ctx.SetStatus(404).Json("Not Found")
		return
	}

	// 设置context中的handlers字段
	ctx.SetHandlers(node.handlers)

	// 设置路由参数
	params := node.parseParamsFromEndNode(r.URL.Path)
	ctx.SetParams(params)

	// 调用路由函数
	if err := ctx.Next(); err != nil {
		ctx.SetStatus(500).Json("Internal Error")
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
	go func() {
		server.ListenAndServe()
	}()

	// 关闭连接
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	// 阻塞当前Goroutine等待信息
	<-quit

	if err := server.Shutdown(context.Background()); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
}
