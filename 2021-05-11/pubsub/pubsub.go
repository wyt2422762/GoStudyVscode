package pubsub

import (
	"sync"
	"time"
)

//简单的发布订阅模型

type (
	subscriber chan interface{}       //订阅者(任意类型的管道)
	topicFunc  func(interface{}) bool //主题过滤器(一个需要自己实现的返回bool的任意参数的方法)
)

//发布者结构体
type publisher struct {
	name        string                   //发布者名称
	timeout     time.Duration            //超时时间
	m           sync.RWMutex             //读写锁
	subscribers map[subscriber]topicFunc //订阅者信息
	buffer      int                      //通道缓冲区的大小
}

//添加订阅者方法
func (p *publisher) addSubscriber(tp topicFunc) chan interface{} {
	ch := make(chan interface{}, p.buffer)
	p.m.Lock() //加锁
	defer func() {
		p.m.Unlock() //解锁
	}()
	p.subscribers[ch] = tp
	return ch
}

//发布消息方法
func (p *publisher) sendMsg(msg interface{}) {
	p.m.Lock() //加锁，这里控制一次只能发布一条消息
	defer func() {
		p.m.Unlock()
	}()

	//发布消息
	var wg sync.WaitGroup //这个相当于java的countdownlanch
	//新创建一个goroutine来发送消息
	for ch, tp := range p.subscribers {
		wg.Add(1)
		go send(ch, tp, msg, p.timeout, &wg)
	}
}

func send(ch chan interface{}, tp topicFunc, msg interface{}, timeout time.Duration, wg *sync.WaitGroup) {
	defer wg.Done()
	//主题过滤
	if tp != nil && !tp(msg) {
		return //没有通过过滤，直接返回
	}

	select {
	case ch <- msg:
	case <-time.After(timeout): //超时处理
	}
}

//发布者构造方法
func newPublisher(name string, timeout time.Duration, buffer int) *publisher {
	return &publisher{
		name:        name,
		timeout:     timeout,
		subscribers: make(map[subscriber]topicFunc),
	}
}
