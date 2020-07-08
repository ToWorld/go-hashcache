package main

import (
	"hash/crc32"
	"math/rand"
)

func hashByString(key string) int {
	v := int(crc32.ChecksumIEEE([]byte(key)))
	if v >= 0 {
		return v
	}
	return -v
}

func hash() int {
	return rand.Intn(4294967295)
}
