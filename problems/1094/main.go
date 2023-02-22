package main

//func carPooling(trips [][]int, capacity int) bool {
//	arr := make([]int, 1000)
//	for _, trip := range trips {
//		for i := trip[1]; i < trip[2]; i++ {
//			arr[i] += trip[0]
//			if arr[i] > capacity {
//				return false
//			}
//		}
//	}
//	return true
//}

func carPooling(trips [][]int, capacity int) bool {
	arr := make([]int, 1001)
	for _, trip := range trips {
		arr[trip[1]] += trip[0]
		arr[trip[2]] -= trip[0]
	}
	if arr[0] > capacity {
		return false
	}
	for i := 1; i < 1001; i++ {
		arr[i] = arr[i-1] + arr[i]
		if arr[i] > capacity {
			return false
		}
	}
	return true
}

func main() {
	carPooling([][]int{{8, 2, 3}, {4, 1, 3}, {1, 3, 6}, {8, 4, 6}, {4, 4, 8}}, 12)
}
