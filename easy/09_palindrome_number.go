package easy

func isPalindrome(x int) bool {
	var numbers []int
	if x < 0 {
		return false
	}
	for x > 0 {
		numbers = append(numbers, x%10)
		x = x / 10
	}
	for i := 0; i < (len(numbers) / 2); i++ {
		if numbers[i] != numbers[len(numbers)-i-1] {
			return false
		}
	}
	return true
}
