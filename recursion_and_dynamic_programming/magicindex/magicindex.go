package magicindex

func FindMagicIndex(nums []int) int {
	for i, n := range nums {
		if i == n {
			return i
		}
	}
	return -1
}

func FindMagicIndex2(nums []int) int {
	return findMagicIndex2(nums, 0, len(nums)-1)
}

func findMagicIndex2(nums []int, start, end int) int {
	if start > end {
		return -1
	}
	mid := (start + end) / 2
	if mid < nums[mid] {
		return findMagicIndex2(nums, start, mid-1)
	}
	if mid > nums[mid] {
		return findMagicIndex2(nums, mid+1, end)
	}
	return mid
}
