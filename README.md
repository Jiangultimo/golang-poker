Golang 征途

- [Program Pactice](https://github.com/Neras/golang-poker/tree/program-pactice)
- [Go语言圣经-章节练习代码](./gopl.io/)

**NOTE**
1. 变量声明
    ```go
    s := "" // 短变量声明，最简洁，但是只能用在函数内部，而不能用于包变量
    var s = string // 形式以来于字符串的默认初始化零值机制，被初始化为""
    var s = "" // 用得很少，除非同时声明多个变量
    var s string = "" // 显式地标明变量的类型，当变量类型与初值类型相同时，类型冗余，但如果两者类型不同，变量类型就必须了
    ```

