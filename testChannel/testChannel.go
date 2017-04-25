package main

import (
	"fmt"
)

func main() {
	resultChan := make(chan int, 10)
	for i := 1; i <= 10; i++ {
		go func() {
			//收集结果，比如接口调用的返回的结构，存入channal中
			resultChan <- 1
		}()
	}

	//存放resultChan中的值
	var res []int

	//一直等待取够 10 个
	for len(res) < 10 {
		select {
		case tmp := <-resultChan:
			res = append(res, tmp)
		default:
			//你的代码其实是走到了这里，然后退出了
		}
	}
	fmt.Println(res)
}
