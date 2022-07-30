package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	//创建一个context上下文对象
	//context.Background 创建一个context对象
	//context.TODO  在不确定使用什么context时候使用
	//context.WithDeadline  超时时间
	//context.WithValue  向context添加键值对
	//context.WithCancel 创建一个可取消的context

	baseCtx := context.Background()
	//为context赋值
	ctx := context.WithValue(baseCtx, "a", "b")
	go func(c context.Context) {
		fmt.Println(c.Value("a"))
	}(ctx)

	timeoutCtx, cancel := context.WithTimeout(baseCtx, 2*time.Second)
	defer cancel()
	go func(ctx context.Context) {
		//节拍器
		ticker := time.NewTicker(time.Second)
		for range ticker.C {
			select {
			case <-ctx.Done():
				fmt.Println("child process interrupt...")
				return
			default:
				fmt.Println("enter default")
			}
		}
	}(timeoutCtx)
	time.Sleep(time.Second)
	select {
	case <-timeoutCtx.Done():
		time.Sleep(time.Second)
		fmt.Println("main process exit!")
	}
}
