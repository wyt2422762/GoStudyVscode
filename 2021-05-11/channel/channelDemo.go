package channel

import (
	"fmt"
	"time"
)

//channel 可以实现多个Goroutines之间传递数据
func channelDemo() {

	//1. 定义1个通道(string类型)
	chann := make(chan string, 1)

	//新启动一个goroutine
	go func() {
		fmt.Println("在一个新的goroutine中向通道写入数据")
		//chann <- "写入的数据"

		//循环向通道中写入数据
		for i := 0; i < 100; i++ {
			chann <- ("写入的数据" + fmt.Sprint(i))
			//写完最后一个数据后关闭通道
			if i == 99 {
				close(chann)
			}
		}
	}()

	//这里从通道中取数据
	//data := <- chann
	// fmt.Println("从通道中取到的数据：", data)

	//这里持续从通道中取数据，读不到数据会阻塞，除非channel关闭
	for data := range chann {
		fmt.Println("从通道中取到的数据：", data)
	}

}

//阻塞demo
func channerDemoSync() {

	//定义通道
	ch := make(chan string)

	go func() {
		for {
			fmt.Println("协程启动")
			str := <-ch
			fmt.Println("从通道中取到的数据：", str)
			fmt.Println("继续做一些事情")
		}
	}()

	time.Sleep(time.Duration(10) * time.Second) //休眠10秒

	ch <- "Hello World"
}

//非阻塞demo
func channerDemoAsync() {

	//定义通道
	ch := make(chan string)

	go func() {
		for {
			fmt.Println("协程启动")
			select {
			case str := <-ch:
				fmt.Println("从通道中取到的数据：", str)
				break
			default:
				fmt.Println("从通道中没有取到的数据")
				break
			}
			fmt.Println("继续做一些事情")
		}
	}()

	time.Sleep(time.Duration(10) * time.Second) //休眠10秒

	ch <- "Hello World"
}

//通道close Demo
func closeDemo() {
	jobs := make(chan string)
	done := make(chan string)

	//工作协程
	go func() {
		for {
			j, ok := <-jobs //ok表示通道是否正常，ok返回false表示通道关闭
			if ok {
				fmt.Println("工作：", j)
			} else {
				fmt.Println("下班啦")
				done <- "下班"
				return
			}
		}
	}()

	for i := 0; i < 10; i++ {
		fmt.Println("指派工作", i)
		jobs <- fmt.Sprint(i)
	}

	close(jobs) //关闭jobs通道
	fmt.Println("没有工作可干了")

	<-done
}

//通道遍历demo
func rangeChannelDemo(){
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2

	close(ch) //非空的通道仍然可以关闭

	//关闭后的通道仍然可以读取和遍历
	for data := range ch {
		fmt.Println(data)
	}
}

//定时器demo
func timerDemo(){
	//创建一个定时器，超时事件为5秒
	tim1 := time.NewTimer(time.Second * 5)
	d1:= <- tim1.C
	fmt.Println("tim1完毕:", d1)

	tim2 := time.NewTimer(time.Second * 5)
	go func() {
        d2:= <- tim2.C
        fmt.Println("tim2完毕:", d2)
    }()

	// time.Sleep(time.Duration(10) * time.Second) //休眠10秒
	stoped := tim2.Stop()
	if(stoped){
		fmt.Println("tim2被关闭：", stoped)
	} else {
		fmt.Println(stoped)
	}

}

//打点器demo
func tickerDemo(){
	//创建一个打点器，间隔事件为0.5秒
	ticker := time.NewTicker(time.Millisecond * 500)
	go func() {
        for t := range ticker.C {
            fmt.Println("Tick at", t)
        }
    }()
	time.Sleep(time.Millisecond * 1600) //休眠一会
    ticker.Stop() //停止打点器
    fmt.Println("Ticker stopped")

}
