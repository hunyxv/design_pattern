package strategy

// RemoveRepByLoop 通过两重循环过滤重复元素
func RemoveRepByLoop(slc []int) []int {
	result := make([]int, 0, len(slc)) // 存放结果
	for i := range slc {
		flag := true
		for j := range result {
			if slc[i] == result[j] {
				flag = false // 存在重复元素，标识为false
				break
			}
		}
		if flag { // 标识为false，不添加进结果
			result = append(result, slc[i])
		}
	}
	return result
}

// RemoveRepByMap 通过map主键唯一的特性过滤重复元素
func RemoveRepByMap(slc []int) []int {
	result := make([]int, 0, len(slc))
	temp := map[int]struct{}{}
	for _, item := range slc {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}

// RemoveRep 元素去重
func RemoveRep(slc []int) []int {
	if len(slc) < 1024 {
		// 切片长度小于1024的时候，循环来过滤
		return RemoveRepByLoop(slc)
	}
	// 大于的时候，通过map来过滤
	return RemoveRepByMap(slc)
}
