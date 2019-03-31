//题目地址：https://leetcode-cn.com/explore/interview/card/top-interview-questions-easy/1/array/26/
package 数组

func intersect0(nums1 []int, nums2 []int) []int {
	length1 := len(nums1)
	var nums3 []int
	for i:=0; i <length1; i++ {
		for j:=0;j<len(nums2); j++ {
			if nums2[j] == nums1[i] {
				nums3 = append(nums3, nums1[i])
				nums2 = append(nums2[0:j], nums2[j+1:]...)
				break
			}
		}
	}

	return nums3
}


func intersect(nums1 []int, nums2 []int) []int {
	var nums3 []int
	hasMap := make(map[int]int)
	for _,val := range nums1 {
		hasMap[val]++
	}

 	for _, value := range nums2 {
		times, ok := hasMap[value]

		if ok && times>0 {
			nums3 = append(nums3, value)
			hasMap[value]--
		}
	}

	return nums3
}