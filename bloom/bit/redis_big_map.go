package bit

import (
	"time"

	"github.com/go-redis/redis"
)

// redis使用位图实现
// BitMapArray 位图数组
type BitMapArray struct {
	rc   *redis.Client
	key  string // key
	size uint64
}

// Set 设置值
func (bm *BitMapArray) Set(offset uint64) {
	_, err := bm.rc.SetBit(bm.key, int64(offset), value).Result()
	if err != nil {
		log.Errorf("设置值失败", "offset", offset)
		return
	}
}

// Get 获取值
func (bm *BitMapArray) Verify(offset uint64) bool {
	result, err := bm.rc.GetBit(bm.key, int64(offset)).Result()
	if err != nil {
		log.Errorf("获取值失败", "offset", offset)
		return true
	}
	return result == value
}

// Size 获取大小
func (bm *BitMapArray) Size() uint64 {
	return bm.size
}

// NewBigMapArray redis位图实现数组
func NewBigMapArray(rc *redis.Client, key string, expired time.Duration, size uint64) (*BitMapArray, error) {
	_, err := rc.Set(key, "", expired).Result()
	if err != nil {
		log.Errorf("设置值失败", "err", err, "key", key, "expired", expired)
		return nil, err
	}
	return &BitMapArray{
		rc:   rc,
		key:  key,
		size: size,
	}, nil
}
