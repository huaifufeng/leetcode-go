package algorithm

type LruArray struct {
	data   []int
	length int
	count  int
}

func NewLruArray(len int) *LruArray {
	lru := &LruArray{
		data:   make([]int, len),
		length: len,
		count:  0,
	}

	return lru
}

func (la *LruArray) getItem(val int) int {
	index := -1
	for i := 0; i < la.count; i++ {
		if val == la.data[i] {
			index = i
		}
	}

	//说明不在数组内
	if index == -1 {
		//如果数组满了，将第一个元素删除，这个元素放到末尾
		if la.count == la.length {
			for i := 1; i < la.count; i++ {
				la.data[i-1] = la.data[i]
			}

			la.data[la.count-1] = val
		} else {
			//数组还有位置，就把元素放到数组已有元素的最后
			la.data[la.count] = val
			la.count++
		}
	} else {
		//元素在数组内
		for i := index; i < la.count-1; i++ {
			la.data[i] = la.data[i+1]
		}
		la.data[la.count-1] = val
	}

	return val
}
