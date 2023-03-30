package sort

type Heap struct {
	value []int
	size  int
}

/*
建堆的过程就是从最后一个子堆调整的过程
依次向前调整直到所有子堆调整完成

向前调整时影响到已经调整的子堆的父节点，就需要再次调整该父节点的子堆
再次调整是在调整堆时需要考虑的情况
*/
func createHeap(nums []int) (h Heap) {
	h.value = nums
	h.size = len(nums)

	//从最后一个堆开始
	for i := h.size / 2; i >= 0; i-- {
		adjustHeap(h, i)
	}
	return
}

/*
	调整子堆的过程就是将父，左子，右子的最大值换到父节点
	同时调整以被换的子节点为父节点的子堆的过程
*/

func adjustHeap(h Heap, parentNode int) {
	// 元素以0开始，因此左孩子为*2+1，右孩子为*2+2
	left := parentNode*2 + 1
	right := parentNode*2 + 2

	max := parentNode
	if left < h.size && h.value[left] > h.value[max] {
		max = left
	}
	if right < h.size && h.value[right] > h.value[max] {
		max = right
	}
	if max != parentNode {
		// 将三者最大的值换到父节点，并去调整被换子节点的堆
		h.value[max], h.value[parentNode] = h.value[parentNode], h.value[max]
		adjustHeap(h, max)
	}
}

func HeapSort(nums []int) {
	h := createHeap(nums)

	for h.size > 0 {
		// 每次都将最大的值加入无序序列的最后一个，并加入有序区
		// size-1的意义在于有序区每次多一位，无序区每次少一位
		h.value[0], h.value[h.size-1] = h.value[h.size-1], h.value[0]
		h.size--
		adjustHeap(h, 0)
	}
}
