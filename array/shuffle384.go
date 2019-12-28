package array

/**
384. 打乱数组
https://leetcode-cn.com/problems/shuffle-an-array/submissions/
打乱一个没有重复元素的数组。
*/
import "math/rand"

type Solution struct {
	oldNums       []int
	numsToShuffle []int
}

func Constructor(nums []int) Solution {
	numsCopy := make([]int, len(nums))
	numsToShuffle := make([]int, len(nums))
	copy(numsCopy, nums)
	copy(numsToShuffle, nums)
	solution := Solution{oldNums: numsCopy, numsToShuffle: numsToShuffle}
	return solution
}

/** Resets the array to its original configuration and return it. */
func (this *Solution) Reset() []int {
	return this.oldNums
}

/** Returns a random shuffling of the array. */
func (this *Solution) Shuffle() []int {
	length := len(this.oldNums)
	for i := 0; i < length; i++ {
		pos := rand.Intn(length-i) + i
		this.numsToShuffle[i], this.numsToShuffle[pos] = this.numsToShuffle[pos], this.numsToShuffle[i]
	}

	return this.numsToShuffle
}
