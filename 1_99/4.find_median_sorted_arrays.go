package __99

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var l1, l2, length = len(nums1), len(nums2), (len(nums1) + len(nums2)) / 2
	var n1, n2 = 0, 0
	var left, right = 0, 0
	for n1+n2 <= length {
		left = right
		if n1 < l1 && (n2 >= l2 || nums1[n1] < nums2[n2]) {
			right = nums1[n1]
			n1++
		} else {
			right = nums2[n2]
			n2++
		}
	}
	if (len(nums1)+len(nums2))%2 != 0 {
		return float64(right)
	}
	return float64(right+left) / 2
}
