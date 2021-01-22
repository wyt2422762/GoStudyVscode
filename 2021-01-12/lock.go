package goroutine

import (
	"fmt"
	"sync"
	"time"
)

var lock *sync.Mutex //Go语言包中的sync包提供了两种锁类型：sync.Mutex(普通锁)和sync.RWMutex(读写锁)。
var rwLock *sync.RWMutex
var waitgroup *sync.WaitGroup //sync.WaitGroup 相当于java中的CountDownLatch

var count int

func add() {
	lock.Lock()    //加索
	defer func() { //这里相当于java中的finally
		waitgroup.Done() //计数器-1
		lock.Unlock()    //释放锁
	}()
	count++
}

func mutexDemo() {
	fmt.Println("mutexDemo")

	waitgroup = new(sync.WaitGroup)
	lock = new(sync.Mutex)

	waitgroup.Add(10) //计数器+1
	for i := 0; i < 10; i++ {
		go add()
	}

	waitgroup.Wait() //等待计数器减为0

	fmt.Println("所有线程执行完毕")
	fmt.Println("最终结果: ", count)
}

func read() {
	rwLock.RLock() //加读锁
	defer func() {
		rwLock.Unlock()
		waitgroup.Done()
	}()
	fmt.Println("读方法")
	time.Sleep(5 * time.Second)
}

func write() {
	rwLock.Lock() //加写锁
	defer func() {
		rwLock.Unlock()
		waitgroup.Done()
	}()
	fmt.Println("写方法")
	time.Sleep(5 * time.Second)
}

func rwMutexDemo() {
	waitgroup = new(sync.WaitGroup)
	rwLock = new(sync.RWMutex)

	waitgroup.Add(2)

	go read()
	go write()

	waitgroup.Wait()
}
