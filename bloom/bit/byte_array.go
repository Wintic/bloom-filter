package bit

type ByteArray []byte

// Set 设置值
func (ba ByteArray) Set(offset uint64) {
	ba[offset] = value
}

// Get 获取值
func (ba ByteArray) Verify(offset uint64) bool {
	return ba[offset] == value
}

func (ba ByteArray) Size() uint64 {
	return uint64(len(ba))
}

// NewByteArray 新数组
func NewByteArray(size uint64) ByteArray {
	return make([]byte, size)
}
