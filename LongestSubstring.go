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
func checkExistInt(li []int, target int) bool {
	for i := 0; i < len(li); i++ {
		if target == li[i] {
			return true
		}
	}
	return false
}

func checkExistsStr(li []string, target string) bool {
	for i := 0; i < len(li); i++ {
		if target == li[i] {
			return true
		}
	}
	return false
}

func maxLength(itens []int) int {
	vol := itens[0]
	for i := 0; i < len(itens); i++ {
		if itens[i] > vol {
			vol = itens[i]
		}
	}
	return vol
}

func lengthOfLongestSubstring(s string) int {
    if s == ""{
        return 0
    }
    set := []string{}
	setCount := []int{1}

	for j := 0; j < len(s); j++ {
        set = append(set, string(s[j]))
		for i := j + 1; i < len(s); i++ {
            status := checkExistsStr(set, string(s[i]))
			if status == false {
                set = append(set, string(s[i]))
                st := checkExistInt(setCount, len(set))
				if st == false {
					setCount = append(setCount, len(set))
				}
			} else {
				set = []string{}
                break
			}
		}
	}
    return maxLength(setCount)
}
