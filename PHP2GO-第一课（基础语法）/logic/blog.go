package logic

import (
	"log"
	"github.com/xiaoxiaosu/php2go/01basic/model"
)

func AddBlog(title, content string) (bool, error){
	blog := model.NewBlog(title, content)
	res, err := blog.Add()

	if err != nil {
		log.Printf("%v", err)
		return res, err
	}

	return res,nil
}

func ListBlog() ([]*model.Blog, error){
	blogs, err := model.ListBlog()
	if err != nil {
		log.Printf("%v", err)
		return nil, err
	}

	return blogs, nil
}


