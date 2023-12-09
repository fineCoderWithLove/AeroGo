package aerogo

import (
	"net/http"
)

// 定义一个接收req和resp的函数类型
type HandlerFunc func(ctx *Context)

// 路由表来存储路由信息
type Engine struct {
	router *router
}

// 生成一个路由表的信息
func New() *Engine {
	return &Engine{router: newRouter()}
}

// 添加路由信息
func (engine *Engine) addRoute(method string, PathName string, handler HandlerFunc) {
	engine.router.addRoute(method, PathName, handler)
}

// GET 请求路径
func (engine *Engine) GET(PathName string, handler HandlerFunc) {
	engine.addRoute("GET", PathName, handler)
}

// POST 请求方法
func (engine *Engine) POST(PathName string, handler HandlerFunc) {
	engine.addRoute("POST", PathName, handler)
}

// 封装原生的http.listenAndServe，会自动把路由表注入
func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}

// 用来判断是否存在路由信息。
func (engine *Engine) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	//TODO implement me
	c := newContext(writer, request)
	engine.router.handle(c)
}
