package go_frame

import (
	"net/http"
)

type Core struct {
	CoreConfig
	router map[string] func(w http.ResponseWriter, r *http.Request)
	middlewares []func(w http.ResponseWriter, r *http.Request)
	index int
}

type CoreConfig struct {
	Addr string
}

func NewCore(conf CoreConfig) *Core{
	return &Core{
		conf,
		make(map[string] func(w http.ResponseWriter, r *http.Request)),
		make([]func(w http.ResponseWriter, r *http.Request),0),
		0,
	}
}

func (c *Core) Use(handlers ...func (w http.ResponseWriter, r *http.Request)) {
	c.middlewares = append(c.middlewares, handlers...)
}

func (c *Core) Get(router string, handler func(w http.ResponseWriter, r *http.Request)) {
	r := "GET" + "-" + router
	c.router[r] = handler
}

func (c *Core) Post(router string, handler func(w http.ResponseWriter, r *http.Request)) {
	r := "POST" + router
	c.router[r] = handler
}

func (c *Core) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	// 处理中间件
	for c.index<len(c.middlewares) {
		c.middlewares[c.index](w, r)
		c.index++
	}

	uri := r.RequestURI
	method := r.Method
	path := method + "-" + uri
	if handler, ok := c.router[path]; !ok {
		w.Write([]byte("not found"))
	} else {
		handler(w, r)
	}
}

func (c *Core) Run() {
	http.ListenAndServe(c.Addr, c)
}