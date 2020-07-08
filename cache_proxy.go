package main

import (
	"fmt"
	"sort"
	"strconv"
)

type CacheClient interface {
	Put(key, value string) error
	Get(key string) (string, error)
	Delete(key string) error
	Init()
	Debug() int
}

// CacheProxy cache代理
// 通过数字来快速获取CacheProxy中的某个实例
type CacheProxy struct {
	slot2CacheClient map[int]CacheClient
	realSlot         int
	virtualSlot      int
	slotMap          map[int]int
	minSlot          int
	slotSet          []int
}

// Init 初始化CacheProxy实例
// 若中间报错, 则返回错误
func (proxy *CacheProxy) Init(realSlot, virtualSlot int) error {
	if realSlot <= 0 || virtualSlot < 0 {
		return fmt.Errorf("slot is invalid, realSlot:<%d>, virtualSlot:<%d>", realSlot, virtualSlot)
	}
	proxy.slot2CacheClient = map[int]CacheClient{}
	proxy.slotMap = map[int]int{}

	proxy.minSlot = 4294967295
	proxy.realSlot = realSlot
	proxy.virtualSlot = virtualSlot
	// 创建虚拟节点与物理节点的映射
	for i := 0; i < proxy.realSlot; i++ {
		for j := 0; j < proxy.virtualSlot; j++ {
			tmpHash := hash()
			proxy.slotMap[tmpHash] = i
			if tmpHash < proxy.minSlot {
				proxy.minSlot = tmpHash
			}
		}
	}
	// 创建物理节点例
	// 采用MapClient
	for i := 0; i < proxy.realSlot; i++ {
		proxy.slot2CacheClient[i] = &MapClient{}
		proxy.slot2CacheClient[i].Init()
	}

	return nil
}

func (proxy *CacheProxy) Put(key, value string) error {
	client := proxy.fetchOneClient(key)
	return client.Put(key, value)
}

func (proxy *CacheProxy) Get(key string) (string, error) {
	client := proxy.fetchOneClient(key)
	return client.Get(key)
}

func (proxy *CacheProxy) Delete(key string) error {
	client := proxy.fetchOneClient(key)
	return client.Delete(key)
}

func (proxy *CacheProxy) fetchOneClient(key string) CacheClient {
	hashCode := hashByString(key)
	keys := []int{}
	for key, _ := range proxy.slotMap {
		keys = append(keys, key)
	}
	sort.Ints(keys)
	for _, val := range keys {
		if val >= hashCode {
			return proxy.slot2CacheClient[proxy.slotMap[val]]
		}
	}
	return proxy.slot2CacheClient[proxy.slotMap[proxy.minSlot]]
}

func (proxy *CacheProxy) Print() {
	str := ""
	for _, val := range proxy.slot2CacheClient {
		str += strconv.Itoa(val.Debug()) + ","
	}
	fmt.Println(str)
}
