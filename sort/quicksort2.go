// 快排的第二种写法
// 最经典普通的写法
package sort

func Qsort(nums []int, left, right int) {
	if left < right {
		m := part(nums, left, right)
		Qsort(nums, left, m-1)
		Qsort(nums, m+1, right)
	}
}

func part(nums []int, left, right int) int {
	key := nums[left]
	for left < right {
		for left < right && nums[right] >= key {
			right--
		}
		nums[left] = nums[right]
		for left < right && nums[left] < key {
			left++
		}
		nums[right] = nums[left]
	}
	nums[left] = key
	return left
}
