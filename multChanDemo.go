package main

import (
	"fmt"
	"runtime"
	"time"
)

func main()  {

	chanJob := make(chan map[string]interface{},3)
	fmt.Println("运行的协程数：", runtime.NumGoroutine())
	start := time.Now()
	go func() {
		a := map[string]interface{}{
			"a":[]string{"are","you","ok"},
		}
		time.Sleep(time.Second*1)
		chanJob <- a
	}()

	go func() {

		time.Sleep(time.Second*4)
		chanJob <- map[string]interface{}{
			"b":time.Now(),
		}
	}()

	go func(a int , b int) {
		time.Sleep(time.Second*3)
		chanJob <- map[string]interface{}{
			"c":a+b,
		}
	}(1,2)
	for j :=0; j<3; j++{
		rs:= <- chanJob  //取信号数据，如果取不到则会阻塞
		fmt.Println(rs)
	}
	close(chanJob)
	cost := time.Since(start)
	fmt.Printf("cost=[%s]",cost)
	fmt.Println("执行同步任务")
	syncTask()
}

func syncTask(){
	start1 := time.Now()
	a := []string{"are","you","ok"}
	time.Sleep(time.Second*1)
	b := time.Now()
	time.Sleep(time.Second*2)
	c := func(a,b int) int {
		return a+b
	}
	time.Sleep(time.Second*3)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c(2,3))
	cost := time.Since(start1)
	fmt.Printf("sync,cost=[%s]",cost)
}
