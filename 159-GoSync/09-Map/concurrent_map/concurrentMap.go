package concurrent_map

import (
	"sync"
)

var SHARD_COUNT = 32

// 分成SHARD_COUNT个分片的map
type ConcurrentMap[V any] []*ConcurrentMapShared[V]

// 通过RWMutex保护的线程安全的分片，包含一个map
type ConcurrentMapShared[V any] struct {
	items        map[string]V
	sync.RWMutex // Read Write mutex，guards access to internal map
}

// 创建并发map
func New[V any]() ConcurrentMap[V] {
	m := make(ConcurrentMap[V], SHARD_COUNT)
	for i := 0; i < SHARD_COUNT; i++ {
		m[i] = &ConcurrentMapShared[V]{items: make(map[string]V)}
	}
	return m
}

// 根据key计算分片索引
func (m ConcurrentMap[V]) GetShared(key string) *ConcurrentMapShared[V] {
	return m[uint(fnv32(key))%uint(SHARD_COUNT)]
}

func fnv32(key string) uint32 {
	hash := uint32(2166136261)
	const prime32 = uint32(16777619)
	keyLength := len(key)
	for i := 0; i < keyLength; i++ {
		hash *= prime32
		hash ^= uint32(key[i])
	}
	return hash
}

// 增加或者查询的时候，首先根据分片索引得到分片对象，然后对分片对象加锁进行操作:
func (m ConcurrentMap[V]) Set(key string, value V) {
	// 根据key计算出对应的分片
	shard := m.GetShared(key)
	shard.Lock() // 对这个分片加锁，执行业务操作
	shard.items[key] = value
	shard.Unlock()
}

func (m ConcurrentMap[V]) Get(key string) (V, bool) {
	// 根据key计算出对应的分片
	shard := m.GetShared(key)
	shard.RLock()
	// 从这个分片读取key的值
	val, ok := shard.items[key]
	shard.RUnlock()
	return val, ok
}
