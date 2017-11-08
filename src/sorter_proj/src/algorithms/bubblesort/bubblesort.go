package bubblesort

func BubbleSort(in []int) {
	var sorted bool
	for i := 0; i < len(in) - 1; i++ {
		sorted = true
		for j := 0; j < len(in) -i -1; j++ {
			if in[j] > in[j+1] {
				in[j], in[j+1] = in[j+1], in[j]
				sorted = false
			}
		}
		if sorted {
			break
		}
	}
}
