package main

func main() {
	//nums := []int{3, 2, 4}
	//fmt.Println(twoSum(nums, 6))
}

// 计算2数之和，返回2个数字的下标
// [2, 7, 11, 8] 9 [0, 1]
func twoSum(nums []int, target int) []int {
	// 快排:
	QSort(nums, 0, len(nums)-1)

	return []int{}
}

/**
 * 输入一个整数，输出该数二进制表示中 1 的个数。例如，把 9 表示成二进制是 1001，有 2 位是 1。因此，如果输入 9，则该函数输出 2。
 */
func hammingWeight(num uint32) int {
	res := 0
	for num > 0 {
		if num%2 == 1 {
			res++
		}
		num = num / 2
	}
	return res
}

func intToBin(num uint32) uint64 {
	if num <= 1 {
		return uint64(num)
	}
	var (
		res   uint64
		times uint64 = 1
	)
	for num > 0 {
		res += uint64(num%2) * times
		num /= 2
		times *= 10
	}
	return res
}
