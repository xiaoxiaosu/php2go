package main

import (
	"fmt"
	"net/http"
	"github.com/xiaoxiaosu/php2go/01basic/controller"
)

func main() {
	http.HandleFunc("/blog/add", controller.AddBlog)
	http.HandleFunc("/blog/list", controller.ListBlog)

	err := http.ListenAndServe(":8082", nil)
	if err != nil {
		fmt.Printf("%v", err)
	}
}


