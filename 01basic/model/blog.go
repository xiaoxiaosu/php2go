package model

type Blog struct {
	Title string
	Content string
}

func NewBlog(title, content string) *Blog {
	return &Blog{Title: title, Content: content}
}

func ListBlog() ([]*Blog, error){
	var blogs []*Blog // 声明一个blog
	blogs = append(blogs, &Blog{"标题1", "内容1"})
	blogs = append(blogs, &Blog{"标题2", "内容2"})
	blogs = append(blogs, &Blog{"标题3", "内容3"})

	return blogs, nil
}

func (b *Blog) Add() (bool,error){
	return true,nil
}

