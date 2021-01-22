package db

import(
	"fmt"
	"github.com/garyburd/redigo/redis"
	"time"
)

func redisTest(){
	fmt.Println("go redis 测试")
}

//简单数据类型操作
func redisString(){
	fmt.Println("redis 简单数据类型(这里用的是string)操作")
	conn, err := redis.Dial("tcp", "47.105.71.238:6379")
	checkErr(err)
	v, err := conn.Do("SET", "name", "ted1")
	checkErr(err)
	fmt.Println("SET结果：", v)
	v, err = redis.String(conn.Do("GET", "name"))
	checkErr(err)
	fmt.Println("GET结果：", v)
}

//列表list操作
func redisList(){
	fmt.Println("redis list操作")
	conn, err := redis.Dial("tcp", "47.105.71.238:6379")
	checkErr(err)
	//往列表里rList插入数据，这里采用头插法 aaa - bbb - ccc
	fmt.Println("向rList中插入数据")
	res, err := conn.Do("LPUSH", "rList", "ccc")
	checkErr(err)
	fmt.Println("插入结果：", res)
	res, err = conn.Do("LPUSH", "rList", "bbb")
	checkErr(err)
	fmt.Println("插入结果：", res)
	res, err = conn.Do("LPUSH", "rList", "aaa")
	checkErr(err)
	fmt.Println("插入结果：", res)
	//读取rList中的数据
	fmt.Println("读取rList中的数据")
	vs, err := redis.Values(conn.Do("LRANGE", "rList", 0, -1))
	checkErr(err)
	for _, v := range vs{
		fmt.Println(string(v.([]byte)))
	}
}

//redis连接池
func redisPool(){
	conn := pool.Get()
	defer conn.Close()
	v, err := redis.String(conn.Do("GET", "name"))
	checkErr(err)
	fmt.Println("GET结果：", v)
}

var pool *redis.Pool

func init(){
	fmt.Println("init 方法")
	pool = getPool()
}

//获取连接池
func getPool() *redis.Pool{
	return &redis.Pool{
		MaxIdle: 5,
		MaxActive: 5,
		IdleTimeout: time.Duration(60) * time.Second,
		Dial: func()(redis.Conn, error){
			conn, err := redis.Dial("tcp", "47.105.71.238:6379")
			if err != nil{
				return nil, err
			}
			return conn, err
		},
	}
}







