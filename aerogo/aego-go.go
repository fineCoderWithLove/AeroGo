package aerogo

import (
	"log"
	"net/http"
	"strings"
)

// 定义一个接收req和resp的函数类型
type HandlerFunc func(ctx *Context)

// 路由表来存储路由信息
type Engine struct {
	router *router
	*RouterGroup
	groups []*RouterGroup
}

// 路由分组信息
type RouterGroup struct {
	prefix      string
	middlewares []HandlerFunc
	parent      *RouterGroup
	engine      *Engine
	/*
		1.包含了可以直接注册路由的router
		2.这个RouterGroup中还可以进行分组
		3.可以存储多个子路由分组
	*/
}

// 生成一个路由表的信息
func New() *Engine {
	engine := &Engine{router: newRouter()}
	engine.RouterGroup = &RouterGroup{engine: engine}
	engine.groups = []*RouterGroup{engine.RouterGroup}
	return engine
}

// 将中间件应用于Group
func (group *RouterGroup) Use(middlewares ...HandlerFunc) {
	group.middlewares = append(group.middlewares, middlewares...)
}

// 路由分组函数
func (group *RouterGroup) Group(prefix string) *RouterGroup {
	engine := group.engine
	//update路由的前缀
	newGroup := &RouterGroup{
		prefix: group.prefix + prefix,
		parent: group,
		engine: engine,
	}
	engine.groups = append(engine.groups, newGroup)
	return newGroup
}

func (group *RouterGroup) addRoute(method string, comp string, handler HandlerFunc) {
	pattern := group.prefix + comp
	log.Printf("Route %4s - %s", method, pattern)
	group.engine.router.addRoute(method, pattern, handler)
}

// GET 路由方法直接添加路由信息不进行分组
func (group *RouterGroup) GET(pattern string, handler HandlerFunc) {
	group.addRoute("GET", pattern, handler)
}

// POST 路由方法
func (group *RouterGroup) POST(pattern string, handler HandlerFunc) {
	group.addRoute("POST", pattern, handler)
}

// 封装原生的http.listenAndServe，会自动把路由表注入
func (engine *Engine) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, engine))
}

// 用来判断是否存在路由信息。
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	var middlewares []HandlerFunc
	//遍历分组判断路由前缀
	for _, group := range engine.groups {
		//如果路径有前缀开头，就添加中间件
		if strings.HasPrefix(req.URL.Path, group.prefix) {
			middlewares = append(middlewares, group.middlewares...)
		}
	}
	c := newContext(w, req)
	c.handlers = middlewares
	engine.router.handle(c)
}
