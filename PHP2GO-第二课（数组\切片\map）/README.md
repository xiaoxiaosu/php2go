# PHP2GO第二课

你好，我是小小酥，在上一篇文章中，我们通过横向对比PHP的方式学习了go语言的变量，流程控制，循环，同时介绍了go语言简单的设计哲学，我们会发现这些内容其实在使用上，和php对比起来并没有太大的不同，那么今天我们要讲解的内容，相对php来说，就有些许不同了，请你坐稳，我们准备发车~

## 正餐

### 数组

go语言中数组的声明方式为 [length]T，其中length代表数组的长度，T代表数组的数据类型。

#### 初始化

``` go
package main

func main() {
  arr1 := [5]int{}
  arr2 := [5]int{1,2,3,4,5}
  arr3 := [...]int{1,2,3}
}
```

通过上面的代码片段可以看到，我们初始化了三个数组变量，那么他们有什么不同呢

arr1指定了数组的长度并未指定数组每个元素的值，元素的值会初始化为"零"值（注意：不同的数据类型有不同的零值）

arr2既指定了数组的长度同时指定了数组每个元素的值，

arr3指定了数组每个元素的值，但是长度的地方用...占位，这代表我们不自己指定数组的长度，而是期望编译器通过值的内容自动推断出数组的长度。

#### 赋值/读取

``` go
package main

import "fmt"

func main() {
  arr1 := [5]int{1,2,3,4,5}
  
  arr1[1] = 100 //对数组元素进行赋值
  
  fmt.Println(arr[1]) // 读取数组元素
  
  for i:=0; i<len(arr1); i++ {
    fmt.Println(arr1[i]) // 遍历数组元素
  }
}
```

从上面的代码中我们可以看到和php一样，我们可以通过数组的下标完成对数组的赋值和读取操作，同时也可以通过for循环对数组进行遍历。

我们再看看下面的代码，你认为会发生什么呢？

``` go
package main

func main() {
  var arr = [...]int{1,2,3}
  arr[3] = 4
}
```

执行以后，你会发现这段代码在编译时报错啦

./main.go:5:6: invalid array index 3 (out of bounds for 3-element array)

这里我们不是指定了由编译器自动帮我们计算出数组的长度么，为什么错误还提示我们越界了，那是因为在go语言中，数组的长度在其被声明时就确定好了，虽然通过[...]arr{}的方式可以自动计算数组的长度，但并不代表我们可以在定义好数组以后再去修改数组的长度。

作为phper的你或许就会有疑惑了，在php里我一个[]就能走遍所有场景，不需要关注长度也不需要关注数据类型，咋到go语言里，限制数据类型就算了，每次声明数组还要先确定好长度？这可没法玩儿了，对吧。

别着急，接下来我们介绍的切片类型能让你在使用上和php一样的丝滑！

### 切片

切片(slice)是go语言中一种比较特殊的数据结构，是围绕动态数组的概念构建的，可以按需自动扩容和缩容，这是不是就和我们使用php的数组一样了呢？让我们看下实际的使用。

#### 初始化一个切片

``` go
package main

func main() {
  s := []int{1,2,3,4,5}
}
```

我们看到与初始化一个数组相比，切片的声明仅仅是少了一个长度属性，虽然不需要在初始化时指定长度，但是切片仍然具备长度属性，其长度既为切片的元素个数，我们可以通过len()来获取一个切片的长度。

``` go
package main
import "fmt"

func main() {
  s := []int{1,2,3,4,5}
  fmt.Printf("%d",len(s)) //打印切片的长度：5
}
```

#### 通过make初始化一个切片

除了上面介绍到的的方式来初始化一个切片，通常在实际开发中我们还经常使用make函数来初始化一个切片

``` go
package main

func main() {
  s := make([]int, 5)
  s[0] = 1
  s[1] = 2
  s[2] = 3
  s[3] = 4
  s[4] = 5
}
```

make函数的第一个参数指定了切片的类型，第二个参数指定了切片的长度

#### 基于数组的切片

上面我们讲了两种初始化切片的方式，可以看到切片和数组非常的像，实际上切片的底层结构就是一个数组，我们可以把切片理解为数组指定范围的一个窗口，因此我们也可以基于现有的数组来得到一个切片

``` go
package main

import "fmt"

func main() {
  arr := []int{1,2,3,4,5}
  
  s := arr[2:4]
  
  fmt.Printf("%v", s) //[3,4]
  
}
```

可以看到，我们通过array[strat:end]的方式对数组进行了切片(动词)，从而获得了一个切片(名词)

#### 切片扩容

那么当我们需要对切片进行动态扩容应该怎么做呢？

``` go
package main
import "fmt"

func main() {
  var s []int
  s = append(s,1)
  s = append(s,2)
  s = append(s,3) 
  
  fmt.Printf("%v",s) //[1, 2, 3]
}
```

通过append函数，即可动态的对切片进行自定义的扩容，因为切片的自动扩容我们无需太过关注切片的长度，这和php的索引数组使用起来也就一样丝滑了，是不是感觉我又可以了！

<img src="/Users/yuhao.su/Library/Containers/com.tencent.WeWorkMac/Data/Documents/Profiles/12DD3D24FE8D199A01074CCC2E835488/Caches/Emotions/2022-03/968a08541d8975474c6a3bfe0931576d/968a08541d8975474c6a3bfe0931576d.gif" alt="968a08541d8975474c6a3bfe0931576d" style="zoom:50%;" />

到这里你已经了解了数组和切片，你是不是觉得还差点啥？php中我们经常使用到的关联数组，在go语言中我们怎么没有介绍呢？

因为go语言中并没有关联数组，但是别着急，我们接下来要介绍的map可以在大多数场景下替代php中的关联数组。

### map

map是go语言提供的一种抽象的数据类型，标识一组无序的键值对

#### 初始化map

``` go
package main

func main() {
  m := map[string]int{}
  
  m["xiaoming"] = 90
}
```

通过上面的代码，我们显示的初始化了一个map，并对其进行了赋值操作

#### make map

``` go
package main

func main() {
  m := make(map[string]int,10)
  
  m["xiaoming"] = 90
}
```

在切片中我们引入了可以通过make来初始化一个切片的方式，同样的make也可以用来初始化map，如上make的第一个参数是需要创建的map类型，第二个参数是可选的，代表map的长度，虽然这里可以指定长度，但是map也可以自动扩容而不会局限于这个长度。

#### 获取map的值

``` go
package main

import "fmt"

func main() {
  m := make(map[string]int)
  
  m["xiaoming"] = 90
  
  fmt.Println(m["xiaoming"]) // 90
}
```

到这里，我们是不是发现go语言中map的使用方式和php的关联数组还是很类似的？

#### 遍历map

``` go
package main

import "fmt"

func main() {
  m := make(map[string]int)
  
  m["xiaoming"] = 90
  m["xiaozhang"] = 95
  m["xiaoxiaosu"] = 100
  
  for k,v := range m {
    fmt.Println(k,v) 
  }
  
  /*
  输出结果：
  xiaoxiaosu 100
	xiaoming 90
	xiaozhang 95
  */
}
```

我们可以通过上一篇文章中讲过的for range来遍历map，细心的你一定已经产生疑惑了，我们输出结果的顺序是xiaoxiaosu、xiaoming、xiaozhang，而并非我们存储时的顺序，这是因为我们在介绍map定义时就提到的，map内部存储的是一组无序的键值对，这一点一定要记住，忽略的话很容易产生线上事故。

## 餐后甜点

今天的正餐到这里就结束了，我们学习了go语言中的数组，切片和map，这三种有着相似结构的数据类型在php中我们统称为数组，但是在go语言中却有着不同的底层结构，希望通过今天的文章你能区分他们的区别。

今天的餐后甜点，我们聊一下make和new，这是go语言的高频面试题，所以你一定要掌握它。

make函数我们在今天的文章中已经使用过了，我们通过make可以对切片和map进行初始化并分配内存，实际上new函数也是用于为变量分配内存的函数，那么make和new有什么区别呢？我们来看一组代码

### demo

``` go
package main

import "fmt"

func main() {
    s1 := new([]int)
    s2 := make([]int,5)

    fmt.Println(s1) // &[]
    fmt.Println(s2) // [0 0 0 0 0]
}
```

这里我们通过new和make分别为s1和s2分配了内存

从调用函数本身可以看到，new函数无法指定数据的长度，make函数可以指定数据的长度

从返回值可以看到，new(T)函数返回的是对应T的指针，make(T,args)函数返回的是一个有初始值的T类型，

最后我们做一个总结

**make(T,args)**

* make可以指定类型的长度
* make会为变量分配内存，同时初始化变量
* make返回的是引用类型本身
* make只能对切片、map以及我们后面会讲到的channel分配内存和初始化

**new(T)**

* new不能指定类型的长度
* new会为变量分配内存但是并不会初始化变量
* new返回的是一个指针
* new可以为任意的数据类型分配内存



## 下期预告

下期我们会继续对go语言的核心内容进行讲解，包含了go语言的函数，结构体和接口，敬请期待~

如果你喜欢我的文章，欢迎转发给你的同事和朋友，万分感谢！
