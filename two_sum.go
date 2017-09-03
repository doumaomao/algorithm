func twoSum(nums []int, target int) []int {
	res := []int{}
	var m map[int]int
	m = make(map[int]int)
	for index := 0; index < len(nums); index++ {
		find_num := target-nums[index]
		if m[find_num] > 0 {
			res := []int{m[find_num]-1, index}
     			return res
		}
		m[nums[index]] = index+1
	}
	return res
}
