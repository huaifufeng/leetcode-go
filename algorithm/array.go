package algorithm

type Array struct {
	data   []int //存储的元素列表
	length int   //数组的长度
	count  int   //数组中的元素个数
}

//初始化一个数组的元素，这使用指针，可以修改数组的内容
func NewArray(length int) *Array {
	array := &Array{
		data:   make([]int, length),
		length: length,
		count:  0,
	}
	return array
}

//查找指定位置的元素
func (arr *Array) find(pos int) int {
	if pos < 0 || pos > arr.length {
		return -1
	}

	return arr.data[pos]
}

//在数组指定位置插入一个元素, 头部插入，尾部插入
func (arr *Array) insert(index, item int) bool {
	if arr.count == arr.length {
		return false
	}

	if index < 0 || index > arr.count {
		return false
	}

	//只能够插入到0 - count之间的位置
	for i := arr.count; i > index; i-- {
		arr.data[i] = arr.data[i-1]
	}

	arr.data[index] = item
	arr.count++

	return true
}

//删除数组中的元素
func (arr *Array) delete(index int) bool {
	if index < 0 || index >= arr.count {
		return false
	}

	for i := index; i < arr.count; i++ {
		arr.data[i] = arr.data[i+1]
	}

	arr.count--

	return true
}
