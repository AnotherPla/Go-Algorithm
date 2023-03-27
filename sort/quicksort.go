package sort

import (
	"math/rand"
	"time"
)

func QuickSort(nums []int, left, right int) {
	if left < right {
		m := partition(nums, left, right)
		QuickSort(nums, left, m-1)
		QuickSort(nums, m+1, right)
	}
}

func partition(nums []int, l, r int) int {
	// 可以随机一个在l，r之间的元素做为key，并将它的值和nums[r]交换
	// 减少原来序列可能有序而带来的划分不均匀问题，排序规模大时有效果，规模小时几乎没有效果
	rand.NewSource(time.Now().Unix())
	k := rand.Intn(r-l) + l
	nums[r], nums[k] = nums[k], nums[r]

	key := nums[r]
	// j 是普通的遍历指针，指向当前元素
	// i 指向第一个还没有排序完成的元素
	i, j := l, l
	for j < r {
		// j遍历到比key值小的元素时，将它和i指向的元素值互换
		// i指向元素已经小于key值，向后移动
		// 此处大于小于号决定升序还是降序排列
		if nums[j] < key {
			nums[j], nums[i] = nums[i], nums[j]
			i++
		}
		j++
	}
	// i之前的元素已经都小于key，此时i是第一个不小于key的元素位置
	// 因为比key值小的元素已经被j遍历完，并都放在了i之前
	nums[i], nums[r] = nums[r], nums[i]
	// i此时就是key的正确位置
	return i
}
