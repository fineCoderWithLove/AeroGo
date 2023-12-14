package aerocache

// 封装LRU的方法，使之支持并发读写
type ByteView struct {
	b []byte
}

func (v ByteView) Len() int {
	return len(v.b)
}
func (v ByteView) ByteSlice() []byte {
	return cloneBytes(v.b)
}
func (v ByteView) String() string {
	return string(v.b)
}

// 缓存只读不能修改
func cloneBytes(b []byte) []byte {
	c := make([]byte, len(b))
	copy(c, b)
	return c
}
