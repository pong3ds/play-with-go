package main

import (
	"io/ioutil"

	"github.com/labstack/echo"
)

// Engine is the microservices engine
type Engine struct {
	ech *echo.Echo
}

// NewEngine return new engine
func NewEngine() *Engine {
	return &Engine{
		ech: echo.New(),
	}
}

type IContext interface {
	Logger() ILogger
	Requester() IRequester
	ReadInput() string
	JSONPretty(code int, i interface{}, indent string) error
}

type Context struct {
	ctx echo.Context
}

// Logger return logger
func (ctx *Context) Logger() ILogger {
	return NewLogger()
}

// Requester return requester
func (ctx *Context) Requester() IRequester {
	return NewRequester()
}

func (ctx *Context) ReadInput() string {
	body, error := ioutil.ReadAll(ctx.ctx.Request().Body)
	if error != nil {
		return ""
	}
	return string(body)
}

func (ctx *Context) JSONPretty(code int, i interface{}, indent string) error {
	return ctx.ctx.JSONPretty(code, i, indent)
}

func NewContext(ctx echo.Context) *Context {
	return &Context{
		ctx: ctx,
	}
}

type HandlerFunc func(IContext) error

func (e *Engine) GET(endpoint string, handler HandlerFunc) {
	e.ech.GET(endpoint, func(c echo.Context) error {
		ctx := NewContext(c)
		return handler(ctx)
	})
}

func (e *Engine) Start(port string) {
	e.ech.Start(port)
}
