你好，我是小小酥，在上一篇文章中，我们介绍了go语言中的数组，切片和map，也介绍了面试中经常问到的make和new的区别，今天我们继续介绍go语言中的函数，结构体和接口，坐好了，准备发车~

## 函数

### 定义

```go
func add($a int, $b int) int {
	
} 
```

如上面的代码，go语言的函数由五部分组成

其中我们通过**func**关键字来声明一个函数，**func**后跟着的是**函数名**，**参数列表**，以及**函数返回值。**

我们将这几部分称为**函数签名**（func + 函数名 + 参数列表 + 函数返回值），在函数签名后则是由一对大括号{}包装的函数体。

**PHPER请注意**：和php不同的是，在go语言中，函数的参数和返回值都需要明确参数类型。

### 多返回值

```go
func swap($a int, $b int) (int, int) {
	return $b, $a
}
```

在go语言中，函数是可以拥有多个返回值的，如上面的代码，在交换两个数的场景中，不借助其他函数的情况下，一行代码即可实现。（在php中，我们通常会利用中间变量或者list函数来协助实现）

多返回值，在go语言内置的函数中随处可见，也是我们在go工程中经常使用到的。

### 一等公民

go语言中的函数，**可以存储在变量中，也可以作为参数传递给另一个函数或者作为返回值从另一个函数返回**，我们将拥有这样特性的元素称为一等公民。

**函数作为参数**

```go
package main

var add = func(a, b int) int{ // 定义一个加法函数并存入变量add
	return a+b
}

var sub = func(a, b int) int { // 定义一个减法函数并存入变量sub
	return a-b
}

func calculator(exp func(a, b int) int, a, b int) { //定义一个计算器函数
	exp(a,b)
}

func main() {
	calculator(add, 1, 2) //将加法函数和需要计算的值一起传入计算器
	calculator(sub, 1, 2) //将减法函数和需要计算的值一起传入计算器
}
```

**函数作为返回值**

```go
package main

import "fmt"

func funcMaker(exp string) func(int, int) int { // 声明一个函数生成器，通过传入的字符串来返回不同的函数
	switch exp { 
	case "add":
		return func(a int, b int) int { // 当exp=="add"时，返回一个加法函数
			return a+b
		}
	case "sub":
		return func(a int,b int) int { // 当exp=="sub"时，返回一个减法函数
			return a-b
		}
	default:
		return nil
	}
}

func main() {
	exp := funcMaker("add")
	
	fmt.Println(exp(1,2))
}
```

其实通过上面的两个例子我们可以看到，无论是将函数赋值给一个变量，还是将函数作为返回值，我们的func关键字后面都没有具体的函数名，因此我们也将其称为**匿名函数**。

### 函数的可见性

```go
package user

func getUserInfo(id int) (string,uint8){
	return "xiaoming",19 //返回name,age
}

func GetUserInfo(id) {
	return "xiaoming"
}
```

go语言中，我们通过函数**首字母的大小写**来控制函数的可见性。

对于首字母大写的函数，既为可见性为public的函数，可以被其他包引用。

对于首字母小写的函数，既为可见性为private的函数，只能在当前的包中使用。

### 参数的传递

最后我们再聊一下go语言中函数参数的传递，需要记住的是，在go语言中，函数参数的传递都是**值传递**。

> 值传递
>
> 指在调用函数时，将实际参数复制一份传递到函数中，因此在函数中对参数进行修改，不会影响到实际参数的值。
>
> 引用传递
>
> 指在函数调用时，将实际参数的地址传递到函数中，那么在函数中对参数进行修改时，会影响到实际参数的值。


我们通过两个实例来证实go语言中函数参数的传递都是值传递这一说法。

```go
package main

import "fmt"

func modify(j int) {
	fmt.Printf("变量j的内存地址:%p\n", &j)
	j = 100
}

func main() {
	i := 10
	fmt.Printf("变量i的内存地址：%p\n", &i)
	fmt.Printf("变量i的值:%d\n", i)
	modify(i)
	fmt.Printf("变量i的值:%d\n", i)
}

/*
 * 变量i的内存地址：0xc00001c0b0
 * 变量i的值:10
 * 变量j的内存地址:0xc00001c0b8
 * 变量i的值:10
 */
```

通过上面这段代码的执行，我们可以看到函数内外对变量地址的打印，值是不同的，因此证明函数内部是对参数进行了值拷贝。

那么如果我们参数传递的是指针呢？

```go
package main

import "fmt"
 
func modify(j *int) { // 接受一个指针
	fmt.Printf("变量j的内存地址:%p\n", &j)
	*j = 100
}

func main() {
	i := 10
	fmt.Printf("变量i的内存地址：%p\n", &i)
	fmt.Printf("变量i的值:%d\n", i)
	modify(&i)
	fmt.Printf("变量i的值:%d\n", i)
}

/*
 * 变量i的内存地址：0xc00012a008
 * 变量i的值:10
 * 变量j的内存地址:0xc000124020
 * 变量i的值:100
 */
```

这个例子中，我们对代码进行了改动

modify函数参数的类型是一个指针，通过打印函数内外变量i的内存地址，我们可以看到依然是两个地址，这是因为即使传递的是一个指针，go语言依然会进行一次值拷贝将指针变量放到一个新的内存地址上。

但需要注意的是，最终我们变量的值被修改成了100，这是因为指针变量j的内存地址0xc000124020实际指向的也是变量i所在的内存地址0xc00012a008

## 结构体

在php中，我们通过类(class)来对一类事物进行抽象，但是在go语言中没有类的概念，而是通过结构体(struct)来进行对一类事物的抽象。

### 定义

```go
type User struct {

} 
```

如上面的代码，我们通过**type xxx struct{}**关键字即可定义一个结构体


和class不一样的是，在go语言中，结构体自身也是存在可见性的区分，我们通过结构体名称首字母的大小写来控制其可见性，public的结构体可以被别的包访问，private的结构体只能在本包中访问。


### 属性

```go
type User struct {
	Name string
	age uint8
}
```

这里我们为User结构体注入了name和age两个属性。

可以看到，结构体中属性也可以区分可见性，对于public的属性可以在别的包中被访问，对于private的属性则只能在本包中被访问。

### 方法

go语言中，一个方法必须要归属到某个类型（**不一定是结构体**），我们通过

**func (t T) methodName(参数列表) 返回值列表**

来定义一个方法，其中T，被称为方法的接收者，对应了一个类型，t是对应类型参数的变量名，我们看个demo

```go
type User struct {
	Name string
	age uint8
}

func (u User) SetName(name string) {
	u.Name = name
}
func (u User) SetAge(age uint8) {
	u.age = age
}
```

在这个demo中，我们定义了User的结构体，并为其创建了两个方法，接收者为User，用于设置这个接收者的两个属性。

**T和\*T**

方法的接收者除了是某个类型也可以是某个类型的指针，那么不同的选择有什么区别呢？我们通过实际的例子来学习。

```go
package main

import "fmt"

type User struct {
	Name string
	age uint8
}

func (u User) SetName(name string) {
	u.Name = name
}
func (u User) SetAge(age uint8) {
	u.age = age
}

func main() {
	u := User{}
	fmt.Printf("before name:%#v\n",u.Name)
	u.SetName("xiaoming")
	fmt.Printf("after name:%#v\n",u.Name)
}

/*
 * before name:""
 * after name:""
 */
```

可以看到，对于非指针的接收者，方法中是无法修改原实例的属性的，原因相信你也能猜到一二了，和函数一样，方法中的参数也是值传递，我们修改的只是对于实例的值拷贝。

我们再尝试一下指针类型的接收者。

```go
package main

import "fmt"

type User struct {
	Name string
	age uint8
}

func (u *User) SetName(name string) {
	u.Name = name
}
func (u *User) SetAge(age uint8) {
	u.age = age
}

func main() {
	u := &User{}
	fmt.Printf("name:%#v\n",u.Name)
	u.SetName("xiaoming")
	fmt.Printf("name:%#v\n",u.Name)
}

/*
 * before name:""
 * after name:"xiaoming"
 */
```

通过指针类型的接收者，我们成功的修改了实例的属性。

**所以，如果你需要修改实例的属性则需要选择指针类型作为接收者。**

### 组合优于继承

go语言在设计的时候，没有支持继承，而是提倡**组合优于继承**的理念。

那么什么是组合呢，和字面意思一样，我们将多个结构体嵌套在一起的方式就叫组合，我们来看一个demo

```go
package main

import "fmt"

type Worker struct {

}

func (w Worker) work() {
	fmt.Println("working")
}

type User struct {
	Name string
	age uint8
	Worker
}

func main() {
	u := User{}
	u.work()
}
```

在这个例子中，我们定义了两个结构体，worker和user，并且将worker注入了user的属性中，这样结构体嵌套的方式就完成了结构体的组合，我们的user成功的变成了打工人，可以执行work方法了，是不是很简单。

## 接口

接口是**一组方法签名的集合**，在实际开发的过程中，我们通常会将一些对象的行为抽象成接口的形式以实现依赖注入，工厂模式等场景的应用。

其实这段对接口的诠释不止适用于go，也适用于php，java等编程语言。

那么php的接口和go语言的接口在我们使用上又有什么不同呢？我们接着往下看。

### 定义

```go
type p6 interface {
		golang() string
		mysql() string
		redis() string
}
```

举个例子，在这段代码中，我们定义了一个名为p6的接口，内部包含了三个方法的签名。

通过**type 接口名 interface{}**，我们就定义了一个接口。这和php的接口定义除了语法上的区别，没有太大的不同。

### 实现

```go
type p6 interface {
		golang() string
		mysql() string
		redis() string
}

type seniorProgrammer struct {

}

func (s seniorProgrammer) golang() string {

}

func (s seniorProgrammer) mysql() string {

}

func (s seniorProgrammer) redis() string {

}
```

在这段代码中，我们沿用了前文p6的接口，同时我们定义了一个名为seniorProgrammer的结构体，并为这个结构体创建了三个方法，可以看到这三个方法和p6接口中的方法签名是对应的。

到这里，seniorProgrammer的结构体就已经完成了对p6接口的实现。

这和php不一样了是吧，在php中我们需要手动通过implements关键字来显示的实现一个接口，如下

```php
seniorProgrammer implements p6
```

但是在go语言中，只要类型T的方法包含了接口I签名的方法，那么这个类型T就完成了对接口I的实现。用go语言官方的定义则是，**如果一个类型 T 的方法集合是某接口类型 I 的方法集合的等价集合或超集，我们就说类型 T 实现了接口类型 I**

### 应用场景

学会了接口的定义和实现，我们再介绍一下go语言中接口具体的应用场景

**依赖注入**

```go
package main

type database interface { // 定义一个database接口
	insert(data string)
}

type mysql struct{}
func (m *mysql) insert(data string){}

type oracle struct{}
func (m *oracle) insert(data string){}

func addData(db database, data string) { // db参数为database接口类型，以实现不同数据库的注入
	db.insert(data)
}
```

在这个例子中，我们定义了addData的方法，方法包中含一个接口类型的参数，通过这个参数，我们就可以为方法注入不同类型的数据库以实现对数据的写入了。

我们在go内置的包中，其实可以看到很多这种的用法，比如我们第一篇文章中展示过的http.ListenAndServe这个启动一个http服务器的方法，第二个参数就是一个接口类型，我们可以通过实现这个接口就可以接管http请求了，实际上常用的一些go语言的web框架，也是基于这个原理去实现的。

```go
net/http/server.go
type Handler interface { // handler接口的定义
	ServeHTTP(ResponseWriter, *Request)
}

net/http/server.go
func ListenAndServe(addr string, handler Handler) error { //ListenAndServe第二个参数handler就是一个接口类型
	server := &Server{Addr: addr, Handler: handler}
	return server.ListenAndServe()
}

```

看个具体的demo

```go
package main

import "fmt"
import "net/http"

type myHttp struct { // 定义一个myHttp结构体

}

func (m *myHttp) ServeHTTP(w http.ResponseWriter,r *http.Request) { // myHttp实现ServeHTTP方法
    fmt.Println("通过实现ServeHTTP来接管http请求")
}

func main() {
    http.HandleFunc("/hello",func(writer http.ResponseWriter, r *http.Request){
        fmt.Println("hello")
    })
    http.ListenAndServe(":8080",&myHttp{}) // 通过myHttp去接管http请求
}

```

**工厂模式**

go语言中，工厂模式也是接口应用的一个典型场景。

go内置的error，就是一个接口类型，因此我们可以很容易的去实现自己的error类型

```go
builtin/builtin.go

type error interface { //error接口的定义
	Error() string
}
```

看个具体的demo

```go
package main

import "fmt"

type myError struct {
	level string
	info string
}

func newMyError(level string,info string) *myError{
	return &myError{level,info}
}

func (m *myError) Error() string {
	return fmt.Sprintf("level:%v, info:%v", m.level, m.info)
}

func getData(key string) (string, error) {
	return "", newMyError("warning", "not found")
}

func main() {
	if _, err := getData("test"); err != nil {
		fmt.Printf("%v", err)
	}
}
```

### 空接口

在php中，在没有开启强类型模式的情况下，对于一个方法的调用是不需要太多关注参数类型的

```php
<?php
	function add($a,$b) {
		return $a+$b
	}
?>
```

就如上面的代码，$a和$b是可以传任意类型的参数，比如字符串，整形，甚至是数组。

在go语言中，因为空接口本身是没有任何方法签名的，所以任意的数据类型都可以认为是实现了空接口，因此，如果将方法的参数类型设置为interface{}，那么传参的时候就可以指定任意类型的数据了。

```go
package main

import "fmt"

func hello(data interface{}) {
	fmt.Printf("%v", data)
}

func main() {
	hello("xiaoming")
	hello(123)
}
```

如上，我们就通过空接口的方式，让参数具备了“泛型”的能力。

## 餐后甜点

文章的最后我们再分享一下go语言中一个常用的包，fmt包的应用。

> fmt包是go语言中内置的包，通常用于实现数据的格式化输入输出。

### 函数

##### Println

> fmt.Println用于将数据直接输出到控制台，并自动换行
> 

```go
package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello world") //控制台中输出 hello world
}
```

#### Printf

> fmt.Printf用于将数据格式化后在输出到控制台
> 

```go
package main

import (
	"fmt"
)

func main() {
	name := "xiaoming"
	age  := 99

	fmt.Printf("%s is %d years old.", name, age) //控制台中输出 xiaoming is 99 years old
}
```

#### Sprintf

> fmt.Sprintf用于将数据格式化后并返回格式化后的字符串，我们可以将其赋值给某个变量
> 

```go
package main

import (
	"fmt"
)

func main() {
	name := "xiaoming"
	age  := 99
	s := fmt.Sprintf("%s is %d years old.", name, age) //变量s被赋值为“xiaoming is 99 years old”
}
```

#### Scan

> fmt.Scan用于从控制台中读取数据，并将其赋值给指定的变量中
> 

```go
package main

import (
	"fmt"
)

func main() {
	var name string
	fmt.Scan(&name) // 从控制台接受一个字符串，并赋值给name变量

	fmt.Println(name) // 输出name变量的值
}
```

#### Errorf

> fmt.Errorf用于将数据格式化后，并返回一个go语言中内置的error
> 

```go
package main

import (
	"fmt"
)

func main() {
	username := "123"

	if err := addUser(username); err != nil { //调用addUser函数，如果有错误返回则输出该错误
		fmt.Printf("%v", err)
	}
}

func addUser(username string) error{
	if username == "123" {
		return fmt.Errorf("the username is illegal", username) // 判断如果username=="123"，则通过fmt.Errorf返回一个标准错误
	}

	return nil
}
```

### 占位符

| 占位符 | 说明 | 输入 | 输出 |
| --- | --- | --- | --- |
| %v | 输出默认格式的值 | fmt.Printf(”%v”,user) | {zhangsan 99} |
| %+v | 输出结构体时，会包含属性字段的名称 | fmt.Printf(”%+v”,user) | {name:zhangsan age:99} |
| %#v | 输出带go语法的内容 | fmt.Printf(”%#v”,user) | main.User{name:"zhangsan", age:99} |
| %T | 输出数据对应的类型 | fmt.Printf(”%T”,user) | main.User |
| %b | 以二进制的形式输出数据 | fmt.Printf(”%v”,100000000000) | 1011101001000011101101110100000000000 |
| %d | 以十进制的形式输出数据 | fmt.Printf(”%d”,100000000000) | 100000000000 |
| %x | 以十六进制的形式输出数据 | fmt.Printf(”%x”,100000000000) | 174876e800 |
| %p | 以十六进制的形式输出数据的指针地址 | fmt.Printf(”%p”,&user) | 0xc00000c030 |

## 下期预告

今天我们学习了go语言中的函数，结构体和接口，相信聪明的你已经掌握他们的用法了。

下一次我们会继续对go语言的**并发编程**的内容进行讲解，敬请期待~

如果你喜欢我的文章，欢迎关注我的公众号，万分感谢！

![qrcode_for_gh_83255ce34399_344.jpg](https://p6-juejin.byteimg.com/tos-cn-i-k3u1fbpfcp/8cd160bda9c64741b28c65460aa1a92d~tplv-k3u1fbpfcp-watermark.image?)