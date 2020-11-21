# bloom-filter
## 简介
布隆过滤器简单实现-单数组

默认提供两种array实现方式

    1. 使用redis的bitMap实现
    2. 使用内存数组实现

>可拓展，实现 BitArray 接口即可

## Hash 
默认使用 murmur3 散列 
>可以使用自定义Hash

## 使用方法
```
bloom := NewBloomFilter(6, bit.NewByteArray(95851))

bloom.AddItem(offset) 添加
bloom.Verify(offset) 校验是否存在
```