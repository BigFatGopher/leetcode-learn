func merge(nums1 []int, m int, nums2 []int, n int) {
	a1, a2 := m-1, n-1
	x := m + n - 1
	for a2 >= 0 {
		if a1 >= 0 && nums1[a1] > nums2[a2] {
			nums1[x] = nums1[a1]
			a1--
		} else {
			nums1[x] = nums2[a2]
			a2--
		}
		x--
	}
}