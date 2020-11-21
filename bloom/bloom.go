package bloom

import (
	"hash"

	"github.com/Wintic/bloom-filter/bloom/logger"
	"github.com/spaolacci/murmur3"
)

// 设置日志输出
func SetLogger(l logger.Logger) {
	logger.Log = l
}

// 默认日志输出
var log = logger.Log

// Hash 哈希函数
type HashFn hash.Hash64

// BitArray 字节数组
type BitArray interface {
	Set(uint64)
	Verify(uint64) bool
	Size() uint64
}

// BloomFilter 过滤器
type BloomFilter struct {
	hashFn   []HashFn
	bigArray BitArray
	k        uint   // hash次数
	m        uint64 // size 大小
	n        uint64 // 总数量
}

// AddItem 添加元素
func (bf *BloomFilter) AddItem(item []byte) {
	for _, e := range bf.hashFn {
		_, err := e.Write(item)
		if err != nil {
			log.Errorf("添加元素失败", "item", item, "hashFn", e)
			return
		}
		bf.bigArray.Set(e.Sum64() % bf.m)
		e.Reset()
	}
	bf.n++
}

// Verify 校验是否存在
func (bf *BloomFilter) Verify(item []byte) bool {
	for _, e := range bf.hashFn {
		_, err := e.Write(item)
		if err != nil {
			log.Errorf("验证元素失败", "item", item, "hashFn", e)
			// 默认为通过
			return true
		}
		// 位置不为1，则肯定不存在
		if !bf.bigArray.Verify(e.Sum64() % bf.m) {
			return false
		}
		e.Reset()
	}
	return true
}

// NewBloomFilter 创建bloom过滤器
func NewBloomFilter(k uint, array BitArray) *BloomFilter {
	var hashFns []HashFn
	for i := 0; i < int(k); i++ {
		// 生成多个Hash函数
		hashFns = append(hashFns, murmur3.New64WithSeed(uint32(i)))
	}
	return &BloomFilter{
		hashFn:   hashFns,
		bigArray: array,
		k:        k,
		m:        array.Size(),
	}
}

// NewBloomFilterWithHash 创建过滤器（自定义Hash）
func NewBloomFilterWithHash(array BitArray, hashFns ...HashFn) *BloomFilter {
	return &BloomFilter{
		hashFn:   hashFns,
		bigArray: array,
		k:        uint(len(hashFns)),
		m:        array.Size(),
	}
}
