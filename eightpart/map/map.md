# map 的存储方式

## 根据 key 获取 value
1. 计算hash(key)
2. 根据哈希值定位应该在哪个桶
3. 找到桶后需要进一步定位key在桶中哪个槽
4. 当前桶遍历完没有key时，若bmap.overflow不为空，需要继续寻找溢出桶
5. 没有找到key，返回零值和false