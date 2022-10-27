func pickingNumbers(a []int32) int32 {
	// Write your code here
	var sum int32 = 0
	km := make(map[int32]int32)
	for _, val := range a {
		if _, ok := km[val]; ok {
			km[val]++
		} else {
			km[val] = 1
		}
	}
	for key, val := range km {
		var tt int32 = val
		for k, v := range km {
			if key == k {
				continue
			}
			if math.Abs(float64(key-k)) <= 1 {
				tt = val + v
			}
		}
		if tt > sum {
			sum = tt
		}
	}
	return sum

}