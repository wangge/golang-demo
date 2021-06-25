package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

var once sync.Once
var over chan bool = make(chan bool)
var over2 chan bool = make(chan bool)
var sigChan chan os.Signal = make(chan os.Signal)
func main()  {
	go onceDemo()
	go GetSingletonObj()
	//o := <-over
	//if o{
	//	fmt.Println("over")
	//}
	select {
	case <-over:
		fmt.Println("over")
	case <-over2:
		fmt.Println("singleton over")
	}
	signal.Notify(sigChan,syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for{
		s:=<- sigChan
		switch s {
		case syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			fmt.Println("quit")
		default:
			continue
		}
	}

}

func onceDemo()  {
	onceBody := func(){
		fmt.Println("only once")
	}
	done := make(chan bool)
	for i:=0;i<10;i++{
		go func() {
			once.Do(onceBody)
			done <- true
		}()
	}

	for i := 0; i<10;i++{
		<-done
	}

	over <- true
}

// 通过sync.Once 创建单例
type Singleton struct {

}

var singleton *Singleton

func GetSingletonObj() *Singleton {
	once.Do(func() {
		fmt.Println("create obj")
		singleton = new(Singleton)
	})
	return  singleton
}

func TestSingleton()  {
	var wg sync.WaitGroup
	for i:=0; i<5;i++{
		go func() {
			wg.Add(1)
			obj := GetSingletonObj()
			fmt.Sprintf("%p\n",obj)
			wg.Done()
		}()
	}
	wg.Wait()
	over2<-true
}