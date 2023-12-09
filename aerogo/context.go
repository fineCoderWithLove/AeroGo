package aerogo

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// 定义别名
type H map[string]interface{}

type Context struct {
	// origin objects
	Writer http.ResponseWriter
	Req    *http.Request
	// request info
	Path   string
	Method string
	Params map[string]string
	// response info
	StatusCode int
}

func (c *Context) Param(key string) string {
	value, _ := c.Params[key]
	return value
}

// 类似于面向对象
func newContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Writer: w,
		Req:    req,
		Path:   req.URL.Path,
		Method: req.Method,
	}
}

// POST 请求传递的参数
func (c *Context) PostForm(key string) string {
	log.Println(key)
	log.Println(c.Req.FormValue(key))
	return c.Req.FormValue(key)
}

// GET 请求传递参数
func (c *Context) Query(key string) string {
	return c.Req.URL.Query().Get(key)
}

// 封装响应Code
func (c *Context) Status(code int) {
	c.StatusCode = code
	c.Writer.WriteHeader(code)
}

// 封装响应头
func (c *Context) SetHeader(key string, value string) {
	c.Writer.Header().Set(key, value)
}

// 快速构造String的响应方法
func (c *Context) String(code int, format string, values ...interface{}) {
	c.SetHeader("content-Type", "text/plain")
	c.Status(code)
	c.Writer.Write([]byte(fmt.Sprintf(format, values...)))
}

//封装JSON响应格式
/*
	设置响应头
	设置状态码
*/
func (c *Context) JSON(code int, obj interface{}) {
	c.SetHeader("content-Type", "application/json")
	c.Status(code)
	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.Writer, err.Error(), 500)
	}
}

// 快速构造Data
func (c *Context) Data(code int, data []byte) {
	c.Status(code)
	c.Writer.Write(data)
}

// 返回HTML结构
func (c *Context) HTML(code int, html string) {
	c.SetHeader("content-type", "text/html")
	c.Status(code)
	c.Writer.Write([]byte(html))
}
