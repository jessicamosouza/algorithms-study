package swap

func Swap(array []int) []int {
	lenArray := len(array)

	for i := 0; i < lenArray; i++ {
		for j := i + 1; j < lenArray; j++ {
			if array[i] > array[j] {
				a := array[i]
				array[i] = array[j]
				array[j] = a
			}
		}
	}

	return array
}
