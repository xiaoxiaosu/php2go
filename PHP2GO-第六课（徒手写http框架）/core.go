package go_frame

import (
	"net/http"
)

type Core struct {
	CoreConfig
	router map[string] func(w http.ResponseWriter, r *http.Request)
}

type CoreConfig struct {
	Addr string
}

func NewCore(conf CoreConfig) *Core{
	return &Core{
		conf,
		make(map[string] func(w http.ResponseWriter, r *http.Request)),
	}
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