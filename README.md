Golang 征途

- [Program Pactice](https://github.com/Neras/golang-poker/tree/program-pactice)
- [Golang Basic](./basic)
- [Golang Professional](./professional)
- [Go语言圣经-章节练习代码](./gopl.io/)

**NOTE**
- 变量声明
    ```go
    s := "" // 短变量声明，最简洁，但是只能用在函数内部，而不能用于包变量
    var s = string // 形式以来于字符串的默认初始化零值机制，被初始化为""
    var s = "" // 用得很少，除非同时声明多个变量
    var s string = "" // 显式地标明变量的类型，当变量类型与初值类型相同时，类型冗余，但如果两者类型不同，变量类型就必须了
    ```
- 指针：取地址符`&`，放在一个变量前使用就会返回响应变量的内存地址。这个地址可以存储在一个叫做指针的特殊数据类型中。对于任何一个变量var，如下表达式都是正确的：`var == *(&var)`
- 在Go中，函数是不允许被重载的
- 递归--计算斐波那契数列
- 闭包：匿名函数同样被称为闭包，他们被允许调用定义在其他环境下的变量。闭包可使得某个函数捕捉到一些外部状态。另一种表示方式为：一个闭包继承了函数所声明时的作用域。这种状态（作用域内的变量）都被共享到闭包的环境中，因此这些变量可以在闭包中被操作，直到被销毁。闭包经常被用作包装函数：他们会预先定义好一个或多个参数以用于包装。另一个不错的应用就是使用闭包来完成更加简洁的错误检查。
  1. 闭包函数保存并累积其中的变量的值，不管外部函数退出与否，它都能继续操作外部函数中的局部变量。
- 切片（slice）是对数组一个看许片段的引用
  1. new(T) 为每个新的类型T分配一片内存，初始化为 0 并且返回类型为*T的内存地址：这种方法 返回一个指向类型为 T，值为 0 的地址的指针，它适用于值类型如数组和结构体；它相当于`&T{}`。 
  2. make(T) 返回一个类型为 T 的初始值，它只适用于3种内建的引用类型：切片、map 和 channel。
  3. 改变切片长度的过程称之为切片重组**reslicing**
- map是一种特殊的数据结构：一种元素对（pair）的无序集合，pair的一个元素是key，对应的另一个元素的是value，所以这个结构也称为关联数组或字典。map这种数据结构在其他编程语言中也称为字典、hash和HashTable等（对比`JSON`）
  1. map是引用类型，内存用make方法来分配
  2. map初始化`var map1 = make(map[keytype]valyetype)`，或简写为`map1 := make(map[keytype]valuetype)
  3. **不要使用new，永远用make来构造map**
- struct
  1. 试图`make()`一个结构体变量，会引发一个编译错误；`new()`一个映射并试图使用数据填充它，将会引发运行时错误。
- defer：defer的函数会被压入一个栈中，当外层函数返回时，被推迟的函数会按照后进先出的顺序调用。
- 方法：Go方法是作用在接收者（`receiver`）上的一个函数，接收者是某种类型的变量。因此方法是一种特殊类型的函数。
  ```go
  func (recv receiver_type) methodName(parameter_list) (return_value_list) {...}
  ```
  `recv`就像面向对象语言中的`this`或`self`
  1. 如果类型定义了`String()`方法，它会被用在`fmt.Printf()`中生成默认的输出：等同于使用格式化描述符`%V`产生的输出。`fmt.Print()`和`fmt.Println()`也会自动使出`String()`方法
  2. 指针接收者
- GC:通过调用`runtime.GC()`函数可以显示的触发GC。
- 接口：Go语言中的接口痘痕间断，通常它们会包含0个，最多3个方法。
  1. 添加[**第一个例子：使用`Sorter`接口排序**](./interface/sortmain.go)
- 空接口或者最小接口：不包含任何方法，它对实现不做任何要求
  ```go
  type Any interface{}
  ```
- 断言选择
  ```go
  switch v := i.(type) {
  case T:
      // v 的类型为 T
  case S:
      // v 的类型为 S
  default:
      // 没有匹配，v 与 i 的类型相同
  }
  ```
  
**PROFESSIONAL**
- `OpenFile`函数有三个参数：文件名、一个或多个标志（使用逻辑运算符"|"连接），使用文件的权限。通常会用到一下标志：
  1. `os.O_RDONLY`: 只读
  2. `os.O_WRONLY`: 只写
  3. `os.O_CREATE`: 创建：如果指定文件不存在，就创建该文件。
  4. `os.O_TRUNC`: 截断：如果指定文件已存在，就将该文件的长度截为0。
  
  在读文件的时候，文件的权限是被忽略的。
- 切片提供了Go中处理I/O缓冲的标准方式，[在一个切片缓冲内使用无线for循环读取文件](./professional/RWData/cat2.go)
- `fmt.Fprintf()`函数签名为：
  ```go
  func FprintF(w io.Writer, format string, a ...interface{}) (n int, err error)
  ```
- JSON数据格式：
  1. 序列化：是在内存中巴数据转换成指定格式`(data -> string)`，反之亦然`(string -> data structure)`
- Gob是Go自己的以二进制形式序列化和反序列化程序数据的格式；可以在`encoding`包中找到。这种格式的数据简称为Gob（即Go binary的指定）。
- 任何时候需要一个新的错误类型，都可以用`error`包的`errors.New`函数接收合适的错误信息来创建
  ```go
  err := errors.New("math - square root of negative number")
  ```
- 当发生像数组下标越界或者类型断言失败这样的运行错误时，Go运行时会触发运行时panic，伴随程序崩溃抛出一个`runtime.Error`接口类型的值。
- `recover`只能在defer修饰的函数中使用：用于取得panic调用中传递过来的错误值，如果正常执行，调用`recover`会返回nil，且没有其他效果
  ```go
  func protect(g func()) {
    defer func() {
  	    log.Println("done")
	    if err := recover(); err != nil {
   	        log.Printf("run time panic: %v", err)
   	    }
    }()
    log.Println("start")
    g() // posiible runtime-error
  }
  ```
- 所有自定义包实现者应该遵守的最佳实践：
  1. 在包内部，总是应该从panic中recover：不允许显示的超出包范围的panic()
  2. 像包的调用者返回错误值（而不是panic）
- os包含有一个`StartProcess`函数可以调用或启动外部系统命令和二进制可执行文件；它的第一个参数是要运行的进程，第二个三处用来传递选项或参数，第三个参数是含有系统环境基本信息的结构体。这个函数返回被启动进程的id（pid）
- 时间和内存消耗可以用以下便捷脚本来测量：
  ```bash
  /usr/bin/time -f '%Uu %Ss %er %MkB %C' "%@"
  ```
- **不要通过共享内存来通信，而通过通信来共享内存。**
  1. 当`main()`函数返回的时候，程序退出；它不会等待任何其他非main协程的结束。
  2. 默认情况下，通信是同步且无缓冲的：在有接受者接收数据之前，发送不会结束。
  3. 带缓冲的通道
      ```go
      buf := 100
      ch1 := make(chan string, buf)
      ```
     在缓冲满载之前，给一个带缓冲的通道发送数据是不会阻塞的，而从通道读取数据也不会阻塞，直到缓冲空了。
  4. `for`循环的`range`语句可以用在通道上`ch`上，就可以从通道中取值：
      ```go
      for v := range ch {
 	      fmt.Printf("The value is %v\n", v)
      }
      ```
  5. 通道类型可以用注解来表示它只发送或者只接收：
      ```go
      var send_only chan <- int // channel can only receive data
      var recv_only <-chan int // channel can only send data
      ```
- 使用select切换协程
- 使用锁的情景：
  1. 访问共享数据结构中的缓存信息
  2. 保存应用程序上下文和状态信息数据
- 使用通道的情景：
  1. 与异步操作的结果进行交互
  2. 分发任务
  3. 传递数据所有权
- 惰性求值：只在需要时进行求值，同时保留相关变量资源（内存和CPU）
- Futures：在使用某一个值前需要先对其计算，这种情况下，你就可以在另一个处理器上进行该值的计算，到使用时，该值已经计算完毕了。