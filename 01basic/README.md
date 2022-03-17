你好，我是小小酥，欢迎来到我们php2go的第一课，go的基础语法。

## 开胃小菜

在开始讲解go的基础语法之前，我想先请你耐心看以下以下代码（github地址：https://github.com/xiaoxiaosu/php2go/tree/main/01basic），相信作为有php经验的你，一定能看懂。

main.go

``` go
package main

import (
	"fmt"
	"net/http"
	"github.com/xiaoxiaosu/php2go/01basic/controller"
)

func main() {
	http.HandleFunc("/blog/add", controller.AddBlog)
	http.HandleFunc("/blog/list", controller.ListBlog)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("%v", err)
	}
}

```

controller/blog.go

```go
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
```

logic/blog.go

``` go
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
```

model/blog.go

``` go
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
```

看完代码，你应该能发现这就是一个最简单的基于golang的博客系统的模拟实现，为了简化phper学习，我将目录结构按照php框架习惯的controller,logic,model做了划分。

## 正餐

食用完开胃小菜，你应该发现了，按照php的思路，虽然是能看懂项目的功能，但是很多语法和php还是有区别的，所以接下来我们会对这些内容和php对比起来一一进行学习。

### 入口文件

php的入口是没有强制的要求的，在web项目中，我们往往会在nginx中将index.php配置为我们的入口文件。

go的入口一定是main包(package main)下的main函数

``` go
package main

func main() {
  
  
}
```

这里我们引出了package的概念，这是php中没有的，package是go语言中基本的管理单元，遵循以下原则

1. 一个目录下的同级文件归属一个包
2. 包名可以与其目录不同名
3. 包名为 main 的包为应用程序的入口包，编译源码没有 main 包时，将无法编译输出可执行的文件

为了便于你的理解，这里你可以先将其类比为php的命名空间，虽然这两者之间有着本质上的不同，但是随着学习的深入，你自然就能理解其中的区别。

### 变量

因为php是动态的弱类型语言，在运行时确定变量类型且变量类型确定后能自动转换的语言，所以我们在日常使用的过程中，通常是'$'走天下，并不会在声明时去指定变量的类型。

go是静态的强类型语言，在编译时确定变量的类型并且变量类型确定后就不能自动转换的语言，对于这类型的语言我们怎么定义变量呢？以下我会介绍几种go语言定义变量的方式。

``` go
package main

func main() {
  var a int //声明变量名，变量类型并赋值
  
  var b int = 10 //声明变量名，变量类型并赋值
  
  var c = "hello" //声明变量并赋值
  
  d := "world"  //短变量声明
}
```

通过这几种定义变量的方式，我们可以看到虽然go语言是静态的强类型语言，但是go自身是支持了变量类型推导的，所以大多数情况下和我们平常写php一样，是无需主动在声明时指定变量类型的，因此上面的第四种，短变量声明的方式是go官方推荐且go工程中使用最多的方式。

但是和php不同的是，因为go本身是强类型语言，并不支持在确定变量类型后再对变量类型进行修改，所以一下代码在go编译时会提示错误。

``` go
package main

func main() {
    a := 10 //定义变量a=10
    a = "hello" //将a重新赋值为"hello"
}

# command-line-arguments //编译时报错
./main.go:5:7: cannot use "hello" (type untyped string) as type int in assignment
```

前文我们有提到，大多数情况下，我们在go工程中定义变量是通过短变量声明的方式去定义的，那么什么情况下我们不能这样做呢，这里就需要引出包变量和局部变量的概念了，请看如下代码

``` go
package main

var b int

func main() {
  a := 10
}
```

在php中，函数内部的我们称为局部变量，函数外的我们称为全局变量，函数内需要使用全局变量的话通过$GLOBALS即可访问

在go中，函数内部的我们称为局部变量，函数外部的我们称为包变量，相同包的函数内可以直接使用到包变量



### 流程控制

php和golang的流程控制的写法上，可以说几乎没有区别，关键词也是'if','else','else if,'switch'，所以需要关注两点即可：

1. go中if的条件不需要加括号
2. go中switch不需要主动加break来跳出流程控制，因为go中默认执行完一个流程后会自动break掉，如果需要继续执行下一个流程在该流程语句中主动加上fallthrough即可

``` go
package main

func main() {
  
  if a>b { //if
    
  }
  
  if a>b { //if else
    
  } else {
    
  }
  
  if a>b { //if else if
    
  } else if a>c {
    
  } else {
    
  }
  
  switch exp{
    case a:
    case b:
    case c:
    default:
  }
  
}
```

### 循环

php中支持的循环方式包含了for，foreach，while，do while

go中支持的循环方式包含了for，for range，go本身并没有提供大多数语言提供的while和do while，因为在go中通过for的各种写法即可实现while和do while的作用

``` go
package main

func main() {
  for i:=0; i<n; i++ {
    
  }
  
  for ;i<n; i++{ 
    
  }
  
  for ; ; i++ {
    
  }
  
  for {
    
  }
}
```

go中for range在使用上也和php的foreach类似，都是做变量的键值循环用

``` go
package main

func main() {
  arr := [5]int{1,2,3,4,5}
  
  for k,v := range arr {
    
  }
}
```



在今天的正餐中，我们通过和php对比学习的方式介绍了go语言中变量的定义方式，流程控制以及循环，可以发现这几部分go的语法和php是非常相似的。

在结束本次的正餐前，我想再提一个日常写php和golang时不同的地方，相信你已经发现了，那就是在go语言的项目中，我们是不需要在每行代码结束时自己写';'的。这是因为golang的编译器会自动帮助我们在行末补齐分号，这也就映射上了go语言崇尚简单的哲学。

## 饭后甜品

在本次的饭后甜品中，我们承接上面提到的go语言的设计哲学之一简单，进行一个介绍

Go 开发者Dave Cheney曾说过：“大多数编程语言创建伊始都致力于成为一门简单的语言，但最终都只是满足于做一个强大的编程语言”。而 Go 语言是一个例外。Go 语言的设计者们在语言设计之初，就拒绝了走语言特性融合的道路，选择了“做减法”并致力于打造一门简单的编程语言。

那么具体简单在哪里呢，我大致做一些罗列

* 简洁、常规的语法，它仅有 25 个关键字
* 内置垃圾收集，降低开发人员内存管理的心智负担
* 首字母大小写决定可见性，无需通过额外关键字修饰
* 内置并发支持，简化并发程序设计
* 变量初始为类型零值，避免以随机值作为初值的问题
* 内置接口类型，为组合的设计哲学奠定基础
* 原生提供完善的工具链，开箱即用
* ... ...

可能作为phper的你会有一些疑惑，觉得垃圾回收，接口类型这些内容不是php中也有么，为什么到golang就上升到简单的设计哲学这个高度了，我的理解是因为php自身其实也是一类简单的语言，甚至在某种程度上比golang更简单，但是golang其实是从c语言演进来的，因此相对c,c++这类语言，golang着实是配得上简单的哲学，同时对比php这类简单的语言，golang又有着编译型语言执行速度快，性能更高以及内置并发支持，原生提供完善的工具链等优势，所以确实值得我们学习和工程上的使用。

## 下期预告

下期我们会继续对go语言的基本语法进行和php对比式的讲解，包含了数组，切片，map，结构体，敬请期待，

如果你喜欢我的文章，欢迎转发给你的同事和朋友，万分感谢！



