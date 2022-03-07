package main

import (
	"fmt"
	"sync"
)

// LFU算法

// 最近最少使用频率

type LFU struct {
	Cap           int
	KeyToValue    map[int]int // 存储值
	KeyToFreq     map[int]int // 存储频率
	FreqToKeysMap map[int][]int
	MinFreq       int
}

var G_LFU *LFU
var lock *sync.Mutex = new(sync.Mutex)

func NewLFU(cap int) *LFU {
	if G_LFU == nil {
		lock.Lock()
		defer lock.Unlock()
		if G_LFU == nil {
			return &LFU{
				Cap:           cap,
				KeyToValue:    make(map[int]int),
				KeyToFreq:     make(map[int]int),
				FreqToKeysMap: make(map[int][]int),
			}
		}
	}
	return G_LFU
}

/*
todo : 从valueMap 中查找相应的键值对 如果存在，也需要更新下key对应的频率 （使用Map + 栈的结构） 推荐使用map + 链表的结构
*/

func (self *LFU) get(key int) int {
	// 1. 从value map中获取
	v, ok := self.KeyToValue[key]
	if !ok {
		return -1
	}

	// 需要更新freqmap
	self.increaseFreq(key)
	return v
}

func (self *LFU) put(key, val int) int {
	if self.Cap <= 0 {
		return -1
	}
	// 分为两种情况 key 不存在和存在
	// key存在 需要更新freqMap
	if _, exist := self.KeyToValue[key]; exist {
		// 更新最新的值
		self.KeyToValue[key] = val
		// 更新频率
		self.increaseFreq(key)
		return 0
	}

	// todo: key不存在，需要判断容量是不是满了
	// 如果容量满了，需要根据最小的minFreq，找到对应的key(列表) 删除里面的key
	if len(self.KeyToValue) >= self.Cap {
		// full
		self.removeMinFreq(self.MinFreq)
	}
	//如果容量未满 valueMap里面新加一个key,并且加入到freqMap key=1 对应的列表中
	//同时更新最小的minFreq = 1
	//更新 KV
	self.KeyToValue[key] = val

	// 更新KF
	self.KeyToFreq[key] = 1
	self.MinFreq = 1
	// 使用map + 栈来实现
	self.FreqToKeysMap[self.MinFreq] = append(self.FreqToKeysMap[self.MinFreq], key)
	return 0
}

// 增加key对应的频率
func (self *LFU) increaseFreq(key int) {
	// todo: 根据key,获取freq
	freq, _ := self.KeyToFreq[key]
	// 更新KF
	self.KeyToFreq[key] = freq + 1

	// todo:移除freqtokeys 建议使用链表删除
	for k, v := range self.FreqToKeysMap[freq] {
		if key == v {
			self.FreqToKeysMap[freq] = append(self.FreqToKeysMap[freq][:k], self.FreqToKeysMap[freq][k+1:]...) // 利用切片左闭右开的特性删除第k个元素
		}
	}
	//todo:将key加入到freq+1对应的列表中
	self.FreqToKeysMap[freq+1] = append(self.FreqToKeysMap[freq+1], key)

	// 需要判断freqToKeys的列表，如果空了，需要移除这个freq
	if len(self.FreqToKeysMap[freq]) == 0 {
		delete(self.FreqToKeysMap, freq)
		// 如果这个freq是minfreq 需要更新一下minFreq
		if freq == self.MinFreq {
			self.MinFreq++
		}
	}

}

// todo: 从最小频率对应的数据中，删除最初加入的key
func (self *LFU) removeMinFreq(minFreq int) {
	// 获取要删除的key
	var delKey int
	if len(self.FreqToKeysMap) > 0 {
		// 删除栈顶
		delKey = self.FreqToKeysMap[minFreq][0]
		self.FreqToKeysMap[minFreq] = self.FreqToKeysMap[minFreq][:len(self.FreqToKeysMap[minFreq])-1]
	}
	// 删除KV表
	delete(self.KeyToValue, delKey)

	// 更新KF表
	delete(self.KeyToFreq, delKey)

}

func main() {
	cache := NewLFU(2)
	cache.put(1, 10)
	cache.put(2, 20)

	fmt.Printf("key = 1;value = %d\n", cache.get(1))
	fmt.Printf("key = 1;value = %d\n", cache.get(2))

	cache.put(3, 30)
	fmt.Printf("key = 1; value = %d\n", cache.get(1))
}
