package sandbox

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"testing"
	"time"
)

// 测试插入数据后统计数据总数
func BenchmarkCount(b *testing.B) {
	memmap := make(map[string]bool)

	t0 := time.Now()
	for i := 0; i < b.N; i++ {
		memmap[strconv.Itoa(rand.Int())] = true
	}
	d := time.Now().Sub(t0)

	fmt.Println(len(memmap), "inserted in", d.Seconds(), "seconds")
}

// 测试并发插入数据后统计数据总数
func BenchmarkCountParallel(b *testing.B) {
	memmap := make(map[string]bool)
	var mux sync.Mutex

	t0 := time.Now()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			mux.Lock()
			memmap[strconv.Itoa(rand.Int())] = true
			mux.Unlock()
		}
	})

	d := time.Now().Sub(t0)
	fmt.Println(len(memmap), "inserted in", d.Seconds(), "seconds")
}
