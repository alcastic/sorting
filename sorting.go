package sorting

func SelectSort(s []int) []int {
	for i := len(s) - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if s[i] < s[j] {
				s[i], s[j] = s[j], s[i]
			}
		}
	}
	return s[:]
}
