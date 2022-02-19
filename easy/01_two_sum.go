package easy

// 02/09/2022 Initial Solution
func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := 1; j < len(nums); j++ {
			if nums[j]+nums[i] == target && i != j {
				return []int{i, j}
			}
		}
	}
	return nil
}
