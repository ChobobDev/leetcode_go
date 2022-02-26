package easy

// I can be placed before V (5) and X (10) to make 4 and 9.
// X can be placed before L (50) and C (100) to make 40 and 90.
// C can be placed before D (500) and M (1000) to make 400 and 900.

// Runtime: 8 ms, faster than 77.36% of Go online submissions for Roman to Integer.
// Memory Usage: 2.8 MB, less than 100.00% of Go online submissions for Roman to Integer.

func romanToInt(s string) int {
	romanInInt := map[string]int{"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000}
	currentVal, prevVal, result := 0, 0, 0
	for i := 0; i < len(s); i++ {
		char := s[len(s)-(i+1) : len(s)-i]
		currentVal = romanInInt[char]
		if currentVal < prevVal {
			result = result - currentVal
		} else {
			result = result + currentVal
		}
		prevVal = currentVal
	}
	return (result)
}
