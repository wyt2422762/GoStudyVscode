package goroutine

import(
	"fmt"
	"testing"
)

func TestHelloWorld(t *testing.T){
	fmt.Println("TestHelloWorld")
	helloWorldDemo()
}

func TestMutexDemo(t *testing.T){
	mutexDemo()
}

func TestRWMutexDemo(t *testing.T){
	rwMutexDemo()
}
