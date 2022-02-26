package easy

import "fmt"

// Runtime: 59 ms, faster than 5.49% of Go online submissions for Palindrome Number.
// Memory Usage: 6.6 MB, less than 8.81% of Go online submissions for Palindrome Number.
// func isPalindrome(x int) bool {
// 	var numbers []int
// 	if x < 0 {
// 		return false
// 	}
// 	for x > 0 {
// 		numbers = append(numbers, x%10)
// 		x = x / 10
// 	}
// 	for i := 0; i < (len(numbers) / 2); i++ {
// 		if numbers[i] != numbers[len(numbers)-i-1] {
// 			return false
// 		}
// 	}
// 	return true
// }

//Runtime: 36 ms, faster than 33.97% of Go online submissions for Palindrome Number.
//Memory Usage: 4.8 MB, less than 62.15% of Go online submissions for Palindrome Number.
func isPalindrome(x int) bool {
	reversedNum := 0
	if x < 0 {
		return false
	}
	for n := x; n > 0; n /= 10 {
		reversedNum = (reversedNum + n%10) * 10
	}
	fmt.Println(reversedNum / 10)
	return (reversedNum / 10) == x
}
