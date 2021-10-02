package main

func main() {
	getMinDeletions("abcab")
}

func getMinDeletions(s string) int {
	var m map[int32]int32 = map[int32]int32{}
	for _, v := range s {
		if _, ok := m[v]; ok {
			m[v]++
		} else {
			m[v] = 0
		}
	}
	var sum int32 = 0
	for _, v := range m {
		sum += v
	}
	return int(sum)
}
