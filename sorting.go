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

func BubbleSort(s []int) []int {
	for i := len(s) - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if s[i] < s[j] {
				s[i], s[j] = s[j], s[i]
			}
		}
	}
	return s[:]
}

func InsertionSort(s []int) []int {
	for i := 1; i < len(s); i++ {
		value := s[i]
		j := i - 1
		for j >= 0 && s[j] > value {
			s[j+1] = s[j]
			j -= 1
		}
		s[j+1] = value
	}
	return s[:]
}
