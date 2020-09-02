package Array

//Given an integer array nums sorted in ascending order, and an integer target.
//
//Suppose that nums is rotated at some pivot unknown to you beforehand (i.e., [0, 1, 2, 4, 5, 6, 7] might become [4, 5, 6, 7, 0,1, 2]).
//
//You should search for target in nums and if you found return its index, otherwise return -1.
//
// 
//
//Example 1:
//
//Input: nums = [4, 5,6, 7, 0, 1, 2], target = 0
//Output: 4
//Example 2:
//
//Input: nums = [4, 5, 6, 7, 0, 1, 2], target = 3
//Output: -1
//Example 3:
//
//Input: nums = [1], target = 0
//Output: -1
// 
//
//Constraints:
//
//1 <= nums.length <= 5000
//-10^4 <= nums[i] <= 10^4
//All values of nums are unique.
//nums is guranteed to be rotated at some pivot.
//-10^4 <= target <= 10^4

func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := (left + right) / 2

		if nums[mid] == target {
			return mid
		}

		// 判断是否在前半部分查找
		if (nums[left] <= target && target <= nums[mid]) ||
			(nums[mid] <= nums[right] && (target < nums[mid] ||
				target > nums[right])) {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return -1
}
