package main

import (
	"fmt"
	frame "go-frame"
	"net/http"
)

func main() {
	conf := frame.CoreConfig{
		Addr: ":8082",
	}
	c := frame.NewCore(conf)

	c.Use(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("middleware")
	})

	c.Get("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello"))
	})

	c.Run()
}
