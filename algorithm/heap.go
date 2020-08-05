package algorithm

type Heap struct {
	data   []int
	length int
	count  int
}

//这里实现的是大顶堆
func NewHeap(len int) (heap *Heap) {
	heap = &Heap{
		data:   make([]int, len+1),
		length: len,
		count:  0,
	}

	return
}

//从下向上堆化
func (h *Heap) insert(val int) {
	if h.length == h.count {
		return
	}

	//数组从1开始， 0不存储数据
	i := h.count + 1
	h.data[i] = val
	h.count++
	//当在堆内，并且子节点比父节点大 进行交换处理
	for i/2 > 0 && h.data[i] > h.data[i/2] {
		h.data[i], h.data[i/2] = h.data[i/2], h.data[i]
		i = i / 2
	}
}

//从上向下堆化
func (h *Heap) delete() {
	if h.count == 0 {
		return
	}

	h.data[1] = h.data[h.count]
	h.count--

	index := 1
	maxPos := 1
	for {
		if 2*index <= h.count && h.data[2*index] > h.data[index] {
			maxPos = 2 * index
		}

		if 2*index+1 <= h.count && h.data[2*index+1] > h.data[maxPos] {
			maxPos = 2*index + 1
		}

		if maxPos == index {
			break
		}

		h.data[index], h.data[maxPos] = h.data[maxPos], h.data[index]

		index = maxPos
	}
}
