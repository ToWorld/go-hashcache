package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	fmt.Println("hello world")
	cacheProxy := &CacheProxy{}
	cacheProxy.Init(10, 300)
	cacheProxy.Put("123", "234")
	fmt.Println("put <123, 234>")
	fmt.Println(cacheProxy.Get("123"))

	startTime := time.Now().Unix()
	for i := 0; i < 1000000; i++ {
		key := strconv.Itoa(i)
		val := strconv.Itoa(i)

		if err := cacheProxy.Put(key, val); err != nil {
			fmt.Println("err", "i", i)
			panic(err)
		}
	}
	endTime := time.Now().Unix()
	fmt.Println("cost:", endTime-startTime)
	cacheProxy.Print()
}
