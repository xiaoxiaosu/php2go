package controller

import (
	"net/http"
	"github.com/xiaoxiaosu/php2go/01basic/logic"
)

func AddBlog(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	title := r.FormValue("title")
	content := r.FormValue("content")

	if title == "" {
		w.Write([]byte("标题不能为空"))
		return
	}

	if content == "" {
		w.Write([]byte("内容不能为空"))
		return
	}
	res, _ := logic.AddBlog(title, content)

	if !res {
		w.Write([]byte("err"))
		return
	}

	w.Write([]byte("ok"))
}

func ListBlog(w http.ResponseWriter, r *http.Request) {
	res ,err := logic.ListBlog()
	if err != nil {
		w.Write([]byte("err"))
		return
	}

	for i:=0; i<len(res); i++ {
		w.Write([]byte("title:" + res[i].Title + " "))
		w.Write([]byte("content:" + res[i].Content))
		w.Write([]byte("\n"))
	}
}

