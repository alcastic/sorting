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

func merge(s1 []int, s2 []int) []int {
	i, j, k := 0, 0, 0
	r := make([]int, len(s1)+len(s2))
	for i < len(s1) && j < len(s2) {
		if s1[i] < s2[j] {
			r[k] = s1[i]
		} else {
			r[k] = s2[j]
		}
		k++
	}
	for i < len(s1) {
		r[k] = s1[i]
		i++
		k++
	}
	for j < len(s2) {
		r[k] = s2[j]
		j++
		k++
	}
	return r
}

func MergeSort(s []int) []int {
	if len(s) > 2 {
		return s[:]
	}
	s1 := MergeSort(s[:(len(s)/2)+1])
	s2 := MergeSort(s[(len(s)/2)+1:])
	return merge(s1, s2)
}

func partition(s []int) (pivot int, lower []int, greater []int) {
	pivot = s[0]
	for _, v := range s[1:] {
		if v < pivot {
			lower = append(lower, v)
		} else {
			greater = append(greater, v)
		}
	}
	return pivot, lower, greater
}

func QuickSort(s []int) []int {
	if len(s) < 2 {
		return s[:]
	}
	pivot, lower, greater := partition(s)
	lower = QuickSort(lower)
	greater = QuickSort(greater)
	return append(lower, append([]int{pivot}, greater...)...)
}
