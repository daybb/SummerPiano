package cache

//使用 sync.Mutex 封装 LRU 的几个方法，使之支持并发的读写。
type ByteView struct {
	b []byte
}

// Len returns the view's length
func (v ByteView) Len() int {
	return len(v.b)
}

//ByteSlice returns a copy of the data as a byte slice
func (v ByteView) ByteSlice() []byte {
	return cloneBytes(v.b)
}

//String returns the data as a string, making a copy if necessary
func (v ByteView) String() string {
	return string(v.b)
}

//b 是只读的，使用 ByteSlice() 方法返回一个拷贝，防止缓存值被外部程序修改
func cloneBytes(b []byte) []byte {
	c := make([]byte, len(b))
	copy(c, b)
	return c
}
