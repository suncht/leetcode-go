package array

import (
	"math"
	"sort"
)

/**
347. 前 K 个高频元素
https://leetcode-cn.com/problems/top-k-frequent-elements/
给定一个非空的整数数组，返回其中出现频率前 k 高的元素。

示例 1:

输入: nums = [1,1,1,2,2,3], k = 2
输出: [1,2]
示例 2:

输入: nums = [1], k = 1
输出: [1]
说明：

你可以假设给定的 k 总是合理的，且 1 ≤ k ≤ 数组中不相同的元素的个数。
你的算法的时间复杂度必须优于 O(n log n) , n 是数组的大小。

*/
type pair struct {
	value  int
	number int
}

func topKFrequent(nums []int, k int) []int {
	var m = make(map[int]int)

	var length = len(nums)
	for i := 0; i < length; i++ {
		m[nums[i]]++
	}

	var pairs []pair
	for k, v := range m {
		pairs = append(pairs, pair{value: k, number: v})
	}
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].number > pairs[j].number
	})

	var res []int
	var count = (int)(math.Min(float64(k), float64(len(pairs))))
	for i := 0; i < count; i++ {
		res = append(res, pairs[i].value)
	}
	return res
}
