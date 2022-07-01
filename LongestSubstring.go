/*
 * https://leetcode.com/problems/longest-substring-without-repeating-characters/
 *
 * 3. Longest Substring Without Repeating Characters
 * Medium
 *
 * Runtime: 7 ms, faster than 93.04% of Go online submissions for Add Two Numbers.
 * Memory Usage: 4.5 MB, less than 81.54% of Go online submissions for Add Two Numbers.
 *
 * Given a string s, find the length of the longest substring without repeating characters.
 *
 * Example 1:
 *
 * Input: s = "abcabcbb"
 * Output: 3
 * Explanation: The answer is "abc", with the length of 3.
 *
 */
func findIndexIt(li []int, target int) bool {
	for i := 0; i < len(li); i++ {
		if target == li[i] {
			return true
		}
	}
	return false
}

func findIndex(li []string, target string) bool {
	for i := 0; i < len(li); i++ {
		if target == li[i] {
			return true
		}
	}
	return false
}

func retLarge(itens []int) int {
	vol := itens[0]
	for i := 0; i < len(itens); i++ {
		for j := i + 1; j < len(itens); j++ {
			if vol < itens[j] {
				vol = itens[j]
			}
		}
	}
	return vol
}
func lengthOfLongestSubstring(s string) int {
	if s == "" {
		return 0
	}
	set := []string{}
	setCount := []int{1}
	sli := strings.Split(s, "")

	for j := 0; j < len(sli); j++ {
		set = append(set, sli[j])
		for i := j + 1; i < len(sli); i++ {
			status := findIndex(set, sli[i])
			if status == false {
				set = append(set, sli[i])
				st := findIndexIt(setCount, len(set))
				if st == false {
					setCount = append(setCount, len(set))
				}
			} else {
				set = []string{}
				break
			}
		}
	}
	return retLarge(setCount)
}
