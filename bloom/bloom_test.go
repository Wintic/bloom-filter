package bloom

import (
	"encoding/binary"
	"fmt"
	"testing"

	"github.com/Wintic/bloom-filter/bloom/bit"
)

// TestBloomFilter 测试
func TestBloomFilter(t *testing.T) {
	bloom := NewBloomFilter(6, bit.NewByteArray(95851))

	var count int32
	// 添加数据
	for i := 0; i < 10000; i++ {
		buf := make([]byte, 8)
		binary.BigEndian.PutUint64(buf, uint64(i))
		if bloom.Verify(buf) {
			fmt.Println("----->", i)
			count++
			continue
		}
		bloom.AddItem(buf)
	}

	fmt.Println("==========================", count)
}
