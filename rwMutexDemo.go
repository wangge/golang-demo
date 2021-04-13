package main
// 读写锁

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var counter2 int = 0

func add2(a, b int, lock *sync.RWMutex) {
	c := a + b
	lock.Lock()
	counter2++
	fmt.Printf("%d: %d + %d = %d\n", counter, a, b, c)
	lock.Unlock()
}

func main() {
	start := time.Now()
	lock := &sync.RWMutex{}
	for i := 0; i < 10; i++ {
		go add2(1, i, lock)
	}

	for {
		lock.RLock()
		c := counter2
		lock.RUnlock()
		runtime.Gosched()
		if c >= 10 {
			break
		}
	}
	end := time.Now()
	consume := end.Sub(start).Seconds()
	fmt.Println("程序执行耗时(s)：", consume)
}
