package main

import (
	"fmt"
	"runtime"
	"time"
)

// 百万级并发
/**
1.线程池workerpool
2.单个执行协程worker
执行逻辑：
1.创建num大小的workerpool
1.2 创建num个worker
将worker的jobqueue,插入到workerpool的workerQueue
2.wp队列获取到值job，再在workerqueue去获取worker1,相当于是取worker的jobqueue,然后将job分配给worker1（相当给worker的jobqueue）;
3.worker一直阻塞，直到jobqueue取到值，执行，然后再阻塞

chan 是引用传达
 */
//job 任务
type Job interface {
	Do() // do something
}

// worker
type Worker struct {
	JobQueue chan Job // 任务队列
	Quit chan bool // 停止当前任务
}

// 新建一个worker通道示例
func NewWorker() Worker {
	return Worker{
		JobQueue: make(chan Job), // 初始化化工作队列为nil
		Quit: make(chan bool),
	}
}

/**
整个过程，每个worker都会被运行在一个协程，
在整个workpool中就会有num个可空闲的worker，
当来一条数据的时候，从workerpool中取出一个空闲的worker去执行Job,
当workerpool中无空闲的worker,就会阻塞等待一个空闲的worker
每读到一个通道参数，就运行一个worker
 */
func (w Worker)Run (wq chan chan Job)  {

	// 这是一个独立的协程，循环读取通道内的数据
	// 保证每读到一个通道参数就去做这件事，没有读到就阻塞
	go func() {
		for{
			wq <- w.JobQueue // 注册工作通道到线程池
			select {
			case job := <-w.JobQueue:
				//读到参数
				job.Do()
			case <- w.Quit: // 终止当前任务
				return

			}
		}
	}()
}

type WorkerPool struct {
	workerlen int // 线程池中worker的大小,worker数量
	JobQueue chan Job // 线程池的job 通道
	WorkerQueue chan chan Job
}

func NewWorkerPool(workerlen int) *WorkerPool  {
	return &WorkerPool{
		workerlen: workerlen, // 开始建立workerlen个worker协程
		JobQueue: make(chan Job), // 工作队列通道
		WorkerQueue: make(chan chan Job, workerlen), // 最大通道参数设置为最大协程数workerlen, worker数量的最大值
	}
}

// 运行线程池
func (wp *WorkerPool) Run()  {
	// 初始化时按照传入的num,启动num个后台协程，然后循环读取Job通道内的数据，
	// 读到一个数据时，再获取一个可用的worker，并将job对象传到该worker的chan通道
	fmt.Println("初始化worker")
	for i:=0;i<wp.workerlen;i++{
		// 新建workerlen 20万个worker协程,每个协程可处理一个请求
		worker := NewWorker() //运行一个协程，将线程池通道的参数传递到 worker协程的通道中进而处理这个请求
		worker.Run(wp.WorkerQueue)
	}

	// 循环获取的可用worker，往worker中写job
	go func() {
		for {
			select {
			case job := <-wp.JobQueue://读取任务
			// 尝试获取一个可用的worker
			// 这将阻塞，直到一个worker空闲
			worker := <- wp.WorkerQueue
			worker <- job // 将任务分配给改worker，插入job到通道
			}
		}
	}()
}


// 测试用例
type Dosomething struct {
	Num int
}

func (d *Dosomething) Do()  {
	fmt.Println("开启线程数：",d.Num)
	time.Sleep(1*1*time.Second)
}

func main()  {
	// 设置最大线程数,值越大处理的越快
	num := 100000
	//  注册工作池 出入任务
	p := NewWorkerPool(num)
	p.Run() //有任务就去做，没有就阻塞，任务做不过来也阻塞

	datanum := 100 * 100 * 100
	go func() {
		for i:=1; i<= datanum;i++{
			sc := &Dosomething{Num: i}
			p.JobQueue <- sc // 往线程池的通道写参数，每个参数相当于一个请求
		}
	}()

	for{
		fmt.Println("运行的协程数：", runtime.NumGoroutine())
		time.Sleep(2 * time.Second)
	}
}
