/*
 * https://leetcode.com/problems/reverse-integer/
 *
 * 7. Reverse Integer
 * Medium
 *
 * Runtime: 0 ms, faster than 100.00% of Go online submissions for Reverse Integer.
 * Memory Usage: 2.1 MB, less than 30.49% of Go online submissions for Reverse Integer.
 *
 * Given a signed 32-bit integer x, return x with its digits reversed.
 * If reversing x causes the value to go outside the signed 32-bit integer range [-231, 231 - 1], then return 0.
 *
 * Assume the environment does not allow you to store 64-bit integers (signed or unsigned).
 *
 * Example 1:
 *
 * Input: x = -123
 * Output: -321
 *
**/
func reverse(x int) int {
	isNeg := false
	if x < 0 {
		isNeg = !isNeg
		x *= -1
	}

	current := 0
	for x > 0 {
		remember := x % 10
		current *= 10
		current += remember
		x /= 10
	}
	if current <= math.MinInt32 {
		return 0
	}
	if current >= math.MaxInt32 {
		return 0
	}
	if isNeg {
		current *= -1
	}
	return current
}