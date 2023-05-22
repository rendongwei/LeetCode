package main

func main() {
	twoSum([]int{2, 7, 11, 15}, 9)
}
func twoSum(nums []int, target int) []int {
	/**
	思路:	利用一个map存储当前位置和值,值作为key,位置作为value
			循环数组,判断如果map中存在target-num当前值,则该值的value和当前position为所需答案
			否则,将当前的位置和值存入map中,进行下一次循环
	*/
	var m = make(map[int]int)
	for i, num := range nums {
		if v, ok := m[target-num]; ok {
			return []int{i, v}
		}
		m[num] = i
	}
	return nil
}
