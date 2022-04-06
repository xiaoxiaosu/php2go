package main

import (
	frame "go-frame"
	"net/http"
)

func main() {
	conf := frame.CoreConfig{
		Addr: ":8082",
	}
	c := frame.NewCore(conf)

	c.Get("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello"))
	})

	c.Run()
}
