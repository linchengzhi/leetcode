package __99

func merge(nums1 []int, m int, nums2 []int, n int) {
	if len(nums2) == 0 {
		return
	}
	var p, q = m - 1, n - 1
	for i := m + n - 1; i >= 0; i-- {
		if q < 0 || (p >= 0 && nums1[p] > nums2[q]) {
			nums1[i] = nums1[p]
			p--
		} else {
			nums1[i] = nums2[q]
			q--
		}
	}
}
