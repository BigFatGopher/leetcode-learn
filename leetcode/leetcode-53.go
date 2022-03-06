func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	maxSum := math.MinInt64
	curSum := 0
	for _, cur := range nums {
		temp := curSum + cur
		curSum = max(temp, cur)
		if curSum > maxSum {
			maxSum = curSum
		}
	}
	return maxSum
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}