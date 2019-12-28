package array

/**
350. 两个数组的交集 II
https://leetcode-cn.com/problems/intersection-of-two-arrays-ii/
给定两个数组，编写一个函数来计算它们的交集。

示例 1:
输入: nums1 = [1,2,2,1], nums2 = [2,2]
输出: [2,2]

示例 2:
输入: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
输出: [4,9]
说明：

输出结果中每个元素出现的次数，应与元素在两个数组中出现的次数一致。
我们可以不考虑输出结果的顺序。
*/
func intersect(nums1 []int, nums2 []int) []int {
	if len(nums1) == 0 || len(nums2) == 0 {
		return make([]int, 0)
	}
	res := make([]int, 0)
	map1 := make(map[int]int)
	for _, v := range nums1 {
		map1[v] += 1
	}

	for _, v := range nums2 {
		if map1[v] > 0 {
			res = append(res, v)
			map1[v] -= 1
		}
	}
	return res
}

// 如果给定的数组已经排好序呢？你将如何优化你的算法？
// 假设数组都是升序排序
func intersect2(nums1 []int, nums2 []int) []int {
	res := make([]int, 0)
	len1 := len(nums1)
	len2 := len(nums2)
	if len1 == 0 || len2 == 0 {
		return res
	}

	if nums1[len1-1] < nums2[0] || nums2[len1-1] < nums1[0] {
		return res
	}

	if len1 <= len2 {
		for i := 0; i < len1; i++ {
			if nums1[i] == nums2[i] {
				res = append(res, nums1[i])
			}
		}
	} else {
		for i := 0; i < len2; i++ {
			if nums2[i] == nums1[i] {
				res = append(res, nums2[i])
			}
		}
	}

	return res
}
