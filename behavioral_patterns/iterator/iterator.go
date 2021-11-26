package iterator

// Iterator 迭代器接口
type Iterator interface {
	HasNext() bool
	Next()
	// 获取当前元素
	CurrentItem() interface{}
}

// ArrayInt 数组
type ArrayInt []int

// Iterator 返回迭代器
func (a ArrayInt) Iterator() Iterator {
	return &ArrayIntIterator{
		arrayInt: a,
		index:    0,
	}
}

// ArrayIntIterator 数组迭代器
type ArrayIntIterator struct {
	arrayInt ArrayInt
	index    int
}

// HasNext
func (iter *ArrayIntIterator) HasNext() bool {
	return iter.index < len(iter.arrayInt)-1
}

// Next
func (iter *ArrayIntIterator) Next() {
	iter.index++
}

func (iter *ArrayIntIterator) CurrentItem() interface{} {
	return iter.arrayInt[iter.index]
}
