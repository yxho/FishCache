package consistenthash

import (
	"hash/crc32"
	"sort"
	"strconv"
)

// Hash maps bytes to uint32
type Hash func(data []byte) uint32

// Map constains all hashed keys
type Map struct {
	hash     Hash
	replicas int
	hashRing []int // Sorted
	hashMap  map[int]string
}

// New creates a Map instance
func New(replicas int, fn Hash) *Map {
	m := &Map{
		replicas: replicas,
		hash:     fn,
		hashMap:  make(map[int]string),
	}
	if m.hash == nil {
		m.hash = crc32.ChecksumIEEE
	}
	return m
}

// Add adds some keys to the hash.
func (m *Map) Add(keys ...string) {
	for _, key := range keys {
		for i := 0; i < m.replicas; i++ {
			hash := int(m.hash([]byte(strconv.Itoa(i) + key)))
			m.hashRing = append(m.hashRing, hash)
			m.hashMap[hash] = key
		}
	}
	sort.Ints(m.hashRing)
}

// Get gets the closest item in the hash to the provided key.
func (m *Map) Get(key string) string {
	if len(m.hashRing) == 0 {
		return ""
	}

	hash := int(m.hash([]byte(key)))
	// Binary search for appropriate replica.
	idx := sort.Search(len(m.hashRing), func(i int) bool {
		return m.hashRing[i] >= hash
	})
	// idx%len(m.hashRing)   =>   如果 idx == len(m.keys)，说明应选择 m.keys[0]，因为 m.keys 是一个环状结构，所以用取余数的方式来处理这种情况。
	return m.hashMap[m.hashRing[idx%len(m.hashRing)]]
}
