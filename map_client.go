package main

import (
	"fmt"
	"sync"
)

type MapClient struct {
	client map[string]string
	mutex  *sync.RWMutex
}

func (client *MapClient) Debug() int {
	keys := []string{}
	for k, _ := range client.client {
		keys = append(keys, k)
	}
	return len(keys)
}

func (client *MapClient) Init() {
	client.client = map[string]string{}
	client.mutex = &sync.RWMutex{}
	return
}

func (client *MapClient) Put(key, value string) error {
	client.mutex.Lock()
	defer client.mutex.Unlock()
	client.client[key] = value
	return nil
}

func (client *MapClient) Get(key string) (string, error) {
	client.mutex.RLock()
	defer client.mutex.RUnlock()
	if value, ok := client.client[key]; ok {
		return value, nil
	}
	return "", fmt.Errorf("key not exist")
}

func (client *MapClient) Delete(key string) error {
	client.mutex.Lock()
	defer client.mutex.Unlock()
	delete(client.client, key)
	return nil
}
