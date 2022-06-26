/*
 * https://leetcode.com/problems/two-sum/
 *
 * 1. Two Sum
 * Easy
 *
 * Runtime: 72 ms
 * Memory Usage: 3.6 MB
 * 57 / 57 test cases passed.
 *
 * Given an array of integers nums and an integer target, return indices of
 * the two numbers such that they add up to target.
 *
 * You may assume that each input would have exactly one solution,
 * and you may not use the same element twice.
 *
 * You can return the answer in any order.
 *
 * Example 1:
 *
 * Input: nums = [2,7,11,15], target = 9
 * Output: [0,1]
 * Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].
 */

func twoSum(nums []int, target int) []int {
	rst := []int{}

	for i := 0; i < len(nums); i++ {
		for x := i + 1; x < len(nums); x++ {
			if (nums[i] + nums[x]) == target {
				rst = append(rst, i, x)
			}
		}
	}

	return rst
}