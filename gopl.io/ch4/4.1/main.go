package main

import "fmt"

type Currency int

const (
    USD Currency = iota
    EUR
    GBP
    RMB
)

func main () {
    var a [3]int
    fmt.Println(a[0])

    for i, v := range a {
        fmt.Printf("%d %d\n", i ,v)
    }
    // 使用`...`则表示数组的长度是根据初始化值的个数来计算
    q := [...]int{1, 2, 3}
    r := [3]int{1, 2}
    fmt.Println(q)
    fmt.Println(r, r[2])

    // 数组字面值
    symbol := [...]string{USD: "$", EUR: "o", GBP: "e", RMB: "r"}
    fmt.Println(RMB, symbol[RMB])

    //
    b := [...]int{99: -1}
    fmt.Println(b)

}
