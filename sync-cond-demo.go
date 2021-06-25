package main

import (
	"fmt"
	"sync"
	"time"
)

var locker sync.Mutex
var cond = sync.NewCond(&locker)
func main()  {
	for i:= 0; i< 10;i++{
		go func(x int) {
			fmt.Println("获取锁前",x)
			cond.L.Lock() // 获取锁
			defer cond.L.Unlock() // 释放锁
			fmt.Println("获取锁后",x)
			cond.Wait() // 等待通知，阻塞当前goroutine
			// 通知到来的时候，cond.wait()就会结束阻塞
			fmt.Println("被唤醒",x)
		}(i)
	}
	time.Sleep(time.Second ) // 睡眠1秒,等待所有goroutine 进入wait阻塞状态
	fmt.Println("signal...")

	cond.Signal() // 1秒后下发一个通知给已经获取到锁的goroutine
	time.Sleep(time.Second )
	fmt.Println("signal...")

	cond.Signal()
	time.Sleep(time.Second )
	// 1秒后广播
	cond.Broadcast()
	fmt.Println("Broadcast...")
	for{

	}

}
