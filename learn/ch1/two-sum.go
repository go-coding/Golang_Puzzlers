package ch1

import "log"

func twoSum(nums []int, target int) []int {
	// index 负责保存map[整数]整数的序列号
	index := make(map[int]int, len(nums))

	log.Printf("index ==>> %v num ==>> %v ", index, nums)

	// 通过 for 循环，获取b的序列号
	for i, b := range nums {

		// 通过查询map，获取a = target - b的序列号

		a := target - b


		j, ok := index[a]

		log.Printf("target ==>> %v i ==>> %v a ==>> %v b ==>> %v  j ==>> %v ok ===>> %v", target, i, b, a, j, ok)

		if ok {
			// ok 为 true
			// 说明在i之前，存在 nums[j] == a
			return []int{j, i}
			// 注意，顺序是j，i
		}

		// 把b和i的值，存入map
		index[b] = i
	}

	log.Printf("index ==>> %v num ==>> %v ", index, nums)

	return nil
}
