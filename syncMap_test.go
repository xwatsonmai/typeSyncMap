package typeSyncMap

import (
	"fmt"
	"sync"
	"testing"
)

func TestNewTypeSyncMap(t *testing.T) {
	data := sync.Map{}
	fmt.Println(data.LoadOrStore("1", "one"))
	fmt.Println(data.CompareAndSwap("1", "one", "two"))
	fmt.Println(data.Load("1"))
	fmt.Println(data.CompareAndDelete("1", "one"))
	fmt.Println(data.Load("1"))
	fmt.Println(data.CompareAndDelete("1", "two"))
	fmt.Println(data.Load("1"))
	fmt.Println("=====================================")
	data2 := NewTypeSyncMap[string, string]()
	fmt.Println(data2.LoadOrStore("1", "one"))
	fmt.Println(data2.CompareAndSwap("1", "one", "two"))
	fmt.Println(data2.Load("1"))
	fmt.Println(data2.CompareAndDelete("1", "one"))
	fmt.Println(data2.Load("1"))
	fmt.Println(data2.CompareAndDelete("1", "two"))
	fmt.Println(data2.Load("1"))
}
