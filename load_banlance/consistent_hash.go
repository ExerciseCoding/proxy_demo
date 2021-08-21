package load_banlance

import (
	"errors"
	"hash/crc32"
	"sort"
	"strconv"
	"sync"
)

// UInt32Slice /**
type UInt32Slice []uint32



type Hash func(data []byte) uint32
type ConsistentHashBanlance struct {
	mux sync.RWMutex
	hash Hash
	replicas int //负载因子
	keys UInt32Slice // 已排序的节点hash切片
	hashMap map[uint32]string //节点哈希和key的map,键是hash值，值是节点key

}
// UInt32Slice是自定义类型，要使用sort.Sort需要实现Len,Less,Swap三个方法

func(s UInt32Slice) Len() int {
	return len(s)
}

func (s UInt32Slice) Less(i,j int) bool{
	return s[i] < s[j]
}
func (s UInt32Slice) Swap(i, j int) {
	s[i],s[j] = s[j],s[i]
}

// NewConsistantHashBanlance 初始化hash环
func NewConsistantHashBanlance( replicas int, fn Hash)*ConsistentHashBanlance{
	m := &ConsistentHashBanlance{
		hash: fn,
		replicas: replicas,
		hashMap: make(map[uint32]string),
	}
	if m.hash == nil{
		m.hash = crc32.ChecksumIEEE
	}
	return m
}


// Add 方法用来添加缓存节点，参数为节点key,比如使用IP
func (c *ConsistentHashBanlance) Add(params ...string)error{
	if len(params) == 0{
		return errors.New("param len 1 at least")
	}
	addr := params[0]
	c.mux.Lock()
	defer c.mux.Unlock()
	// 结合复制因子计算所以虚拟节点的hash值，并存入m.keys中，同时在m.hashMap中保存哈希值和key的映射
	for i := 0; i < c.replicas; i++{
		hash := c.hash([]byte(strconv.Itoa(i)+addr))
		c.keys = append(c.keys,hash)
		c.hashMap[hash] = addr
	}
	// 对所以虚拟节点的哈希值进行排序，方便后面进行二分查找
	sort.Sort(c.keys)
	return nil
}

// Get 方法根据给定的对象获取最靠近它的那个节点
func (c *ConsistentHashBanlance) Get(key string)(string,error){
	if c.isEmpty(){
		return "", errors.New("node is empty")
	}
	hash := c.hash([]byte(key))

	// 通过二分查找获取最优节点，第一个服务器hash值大于数据hash值的就是最优服务器节点, sortSearch使用二分查找
	idx := sort.Search(len(c.keys),func(i int)bool{return c.keys[i] >= hash})

	// 如果查找结果大于服务器节点哈希数组的最大索引，表示此时该对象哈希值位于最后一个节点之后，那么放入第一个节点中
	if idx == len(c.keys){
		idx = 0
	}
	c.mux.RLock()
	defer c.mux.RUnlock()
	return c.hashMap[c.keys[idx]],nil
}

func (c *ConsistentHashBanlance) isEmpty()bool{
	return len(c.keys) == 0
}