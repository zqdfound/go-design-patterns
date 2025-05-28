package structure

import (
	"fmt"
	"net/http"
)

//decorator pattern
//装饰器模式是一种结构型设计模式，它允许你通过将对象放入包含行为的特殊封装对象中来为原对象动态添加新的行为。
//装饰器模式通常用于在不修改现有代码的情况下扩展对象的功能。

// HTTP 中间件装饰器
// Handler 定义处理HTTP请求的接口
type Handler interface {
	ServeHTTP(http.ResponseWriter, *http.Request)
}

// BasicHandler 是基础HTTP处理器
type BasicHandler struct{}

func (h *BasicHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "基本处理器响应")
}

// LoggingDecorator 是日志装饰器
type LoggingDecorator struct {
	Handler Handler
}

func (d *LoggingDecorator) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("请求开始: %s %s\n", r.Method, r.URL.Path)
	d.Handler.ServeHTTP(w, r)
	fmt.Println("请求结束")
}

// AuthDecorator 是认证装饰器
type AuthDecorator struct {
	Handler Handler
}

func (d *AuthDecorator) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Authorization") != "valid-token" {
		http.Error(w, "未授权", http.StatusUnauthorized)
		return
	}
	d.Handler.ServeHTTP(w, r)
}
