package lru

import "container/list"

//缓存放在内存中，限制缓存大小为N
//strategy
//1：FIFO：淘汰最老的。创建队列新增记录到队尾，淘汰队首
//2：LFU：淘汰频率最低的。维护一个按照访问次数排序的队列。缺点是消耗内存，受历史数据影响大
//3：LRU：淘汰最近最少使用。维护一个队列，某条记录被访问则移动到队尾，队首就是最少访问的数据

//LRU核心数据结构：map和double linked list。map存储kv关系，插入查找复杂度O（1）
//v放在双向链表内，移动v到队尾、删除、队尾新增的复杂度为O（1）

// Cache is an LRU cache. It is not safe for concurrent access.
type Cache struct {
	maxBytes  int64 //允许使用的最大内存
	nbytes    int64 //当前已使用的内存
	ll        *list.List
	cache     map[string]*list.Element      //键是字符串，值是双向链表中对应节点的指针
	OnEvicted func(key string, value Value) //某条记录被移除时的回调函数，可以为 nil
}

//双向链表节点的数据类型
//在链表中仍保存每个值对应的 key 的好处在于，淘汰队首节点时，需要用 key 从字典中删除对应的映射
type entry struct {
	key   string
	value Value
}

// Value use Len to count how many bytes it takes
type Value interface {
	Len() int
}

// New is the Constructor of Cache
func New(maxBytes int64, onEvicted func(string2 string, value Value)) *Cache {
	return &Cache{
		maxBytes:  maxBytes,
		ll:        list.New(),
		cache:     make(map[string]*list.Element),
		OnEvicted: onEvicted,
	}
}

// Get look ups a key's value
func (c *Cache) Get(key string) (value Value, ok bool) {
	//根据k查找v，并移动到队尾（双向链表队首队尾是相对的）
	if ele, ok := c.cache[key]; ok {
		c.ll.MoveToFront(ele)
		kv := ele.Value.(*entry)
		return kv.value, true
	}
	return
}

// RemoveOldest removes the oldest item
func (c *Cache) RemoveOldest() {
	//ele 为链表最后一位
	ele := c.ll.Back()
	if ele != nil {
		c.ll.Remove(ele)
		//断言的语法是 x.( T )，x 指的是类型为 interface{} 的变量，T 是我们断言它可能是的类型
		kv := ele.Value.(*entry)
		//从字典中 c.cache 删除该节点的映射关系
		delete(c.cache, kv.key)
		//更新当前所用的内存
		c.nbytes -= int64(len(kv.key)) + int64(kv.value.Len())
		//回调函数 OnEvicted 不为 nil，则调用回调函数
		if c.OnEvicted != nil {
			c.OnEvicted(kv.key, kv.value)
		}
	}
}

// Add adds a value to the cache.
func (c *Cache) Add(key string, value Value) {
	//如果存在value，则更新对应节点的值，并将该节点移到队尾。
	if ele, ok := c.cache[key]; ok {
		c.ll.MoveToBack(ele)
		kv := ele.Value.(*entry)
		c.nbytes += int64(value.Len()) - int64(kv.value.Len())
		kv.value = value
	} else {
		ele := c.ll.PushFront(&entry{
			key:   key,
			value: value,
		})
		c.cache[key] = ele
		c.nbytes += int64(value.Len()) + int64(len(key))
	}
	for c.maxBytes != 0 && c.maxBytes < c.nbytes {
		c.RemoveOldest()
	}
}

// Len 计算获取了多少条数据
func (c *Cache) Len() int {
	return c.ll.Len()
}
