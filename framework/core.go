package framework

import (
	"log"
	"net/http"
	"strings"
)

type Core struct {
	router map[string]*Trie
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
func (c *Core) Get(uri string, handler ControllerHandler) error {
	if err := c.router["GET"].AddRouter(uri, handler); err != nil {
		log.Fatal("add router error", err)
	}
	return nil
}

// 路由注册 Post方法
func (c *Core) Post(uri string, handler ControllerHandler) error {
	if err := c.router["POST"].AddRouter(uri, handler); err != nil {
		log.Fatal("add router error", err)
	}
	return nil
}

// 路由注册 Put方法
func (c *Core) Put(uri string, handler ControllerHandler) error {
	if err := c.router["PUT"].AddRouter(uri, handler); err != nil {
		log.Fatal("add router error", err)
	}
	return nil
}

// 路由注册 Delete方法
func (c *Core) Delete(uri string, handler ControllerHandler) error {
	if err := c.router["DELETE"].AddRouter(uri, handler); err != nil {
		log.Fatal("add router error", err)
	}
	return nil
}

// 路由注册 公共前缀
func (c *Core) Group(prefix string) IGroup {
	return newGroup(prefix, c)
}

func (c *Core) FindRouteByRequest(r *http.Request) ControllerHandler {
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
	// durationCtx, cancel := context.WithTimeout(ctx, 100*time.Second)
	// defer cancel()

	// 寻找路由
	handler := c.FindRouteByRequest(r)

	if handler == nil {
		ctx.Json(404, "Not Found")
		return
	}

	// 调用路由函数
	handler(ctx)
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
