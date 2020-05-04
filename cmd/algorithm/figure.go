package main


func main() {
	//nums := []int{3, 2, 4}
	//fmt.Println(twoSum(nums, 6))

	//l1 := makeListNode(9)
	//l1.Next = makeListNode(9)
	//l2 := makeListNode(1)
	//res := addTwoNumbers(l1, l2)
	//fmt.Printf("%v", res)


}

// 计算2数之和，返回2个数字的下标
// [2, 7, 11, 8] 9 [0, 1]
func twoSum(nums []int, target int) []int {
	h := make(map[int]int)
	for k, v := range nums {
		h[v] = k
	}
	for k, v := range nums {
		expect := target - v
		if ek, ok := h[expect]; ok && ek != k {
			return []int{k, ek}
		}
	}
	return []int{}
}

