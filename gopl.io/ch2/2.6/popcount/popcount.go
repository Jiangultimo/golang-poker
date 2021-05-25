package popcount

// 2. 使用匿名函数处理需要复杂处理的初始化
var pc [256]byte = func () (pc [256]byte) {
    // i 为索引，_：没有用值
    for i, _ := range pc {
        pc[i] = pc[i/2] + byte(i&1)
    }
    return
}()

// 1. 如果有初始化表达式则用表达式初始化，如果有些数据不是通过赋值而是一系列计算过程
// 可以使用一个特殊的`init`初始化函数来简化初始化工作
// 每个文件都可以包含多个`init`初始化函数

//func init() {
//    for i := range pc {
//        pc[i] = pc[i/2] + byte(i&1)
//    }
//}

func PopCount(x uint64) int {
    return int(
        pc[byte(x>>(0*8))] +
            pc[byte(x>>(1*8))] +
            pc[byte(x>>(2*8))] +
            pc[byte(x>>(3*8))] +
            pc[byte(x>>(4*8))] +
            pc[byte(x>>(5*8))] +
            pc[byte(x>>(6*8))] +
            pc[byte(x>>(7*8))])
}
