package LRU_structure

import "container/list"

// 定义LRU数据结构
type keyLru struct {
	limit    int                      // 缓存数量
	evicts   *list.List               // 双向链表用于淘汰数据
	elements map[string]*list.Element // 记录缓存数据
	onEvict  func(key string)
}

// LRU初始化
func NewKeyLru(limit int, onEvict func(key string)) *keyLru {
	return &keyLru{
		limit:    limit,
		evicts:   list.New(),
		elements: make(map[string]*list.Element),
		onEvict:  onEvict,
	}
}

// 添加元素到缓存
func (klru *keyLru) add(key string) {
	// 判断添加的值是否存在于缓存中
	if elem, ok := klru.elements[key]; ok {
		klru.evicts.MoveToFront(elem)
		return
	}

	// 添加新节点
	elem := klru.evicts.PushFront(key)
	klru.elements[key] = elem

	// 检查是否超出容量，如果超出就淘汰末尾节点数据
	if klru.evicts.Len() > klru.limit {
		klru.removeOldest()
	}

}

// 淘汰末尾节点
func (klru *keyLru) removeOldest() {
	elem := klru.evicts.Back() // 获取链表末尾节点
	if elem != nil {
		klru.removeElement(elem)
	}
}

// 删除节点操作
func (klru *keyLru) removeElement(e *list.Element) {
	klru.evicts.Remove(e)
	key := e.Value.(string)
	delete(klru.elements, key)
	klru.onEvict(key)
}
