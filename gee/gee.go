package gee

import (
	"log"
	"net/http"
)

// HandlerFunc 为初始化类型
type HandlerFunc func(ctx *Context)

// Engine 定义实例 router是路由
type Engine struct {
	router *router
}

//把路由添加到实例上去
func (e *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	e.router.addRoute(method, pattern, handler)
}

//GET 请求添加到路由
func (e *Engine) GET(pattern string, handler HandlerFunc) {
	e.addRoute("GET", pattern, handler)
}

//POST 请求添加到路由
func (e *Engine) POST(pattern string, handler HandlerFunc) {
	e.addRoute("POST", pattern, handler)
}

// New 创建一个新gee.Engin
func New() *Engine {
	return &Engine{router: newRouter()}

}

// Run 启动监听和http服务
func (e *Engine) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, e))
}

//继承handler
func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := newContext(w, req)
	e.router.handler(c)
}
