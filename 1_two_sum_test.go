/*
Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
You may assume that each input would have exactly one solution, and you may not use the same element twice.
You can return the answer in any order.

Example 1:

Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].
Example 2:

Input: nums = [3,2,4], target = 6
Output: [1,2]
Example 3:

Input: nums = [3,3], target = 6
Output: [0,1]

Constraints:

2 <= nums.length <= 10^4
-10^9 <= nums[i] <= 10^9
-10^9 <= target <= 10^9
Only one valid answer exists.
*/

package leetcode

import (
	"fmt"
	"reflect"
	"testing"
)

func TwoSum(nums []int, target int) []int {
	m := make(map[int][]int, len(nums))
	// For each element of the array, create a slice of its indices in the array
	for idx, val := range nums {
		m[val] = append(m[val], idx)
	}

	for idx, val := range nums {
		// remove processed values ​​from map
		m[val] = m[val][1:]
		if len(m[val]) == 0 {
			delete(m, val)
		}
		// Knowing the sum, for each element of the array find the second addend
		// and check it in the map
		addend := target - val
		if _, ok := m[addend]; ok {
			return []int{idx, m[addend][0]}
		}
	}

	return []int{0, 0}
}

func TestTwoSum(t *testing.T) {
	testCases := []struct {
		nums     []int
		sum      int
		expected []int
	}{
		{
			nums:     []int{2, 7, 11, 15},
			sum:      9,
			expected: []int{0, 1},
		},
		{
			nums:     []int{3, 2, 4},
			sum:      6,
			expected: []int{1, 2},
		},
		{
			nums:     []int{3, 3},
			sum:      6,
			expected: []int{0, 1},
		},
	}

	for _, tc := range testCases {
		res := TwoSum(tc.nums, tc.sum)
		fmt.Println(res) // print res for debug
		if !reflect.DeepEqual(res, tc.expected) {
			t.Errorf("Expected %v, but got %v", tc.expected, res)
		}
	}
}
