package qsort


func QuickSort(in []int) {
	if len(in) <= 1 {
		return
	}
	mid, i := in[0], 1
	head, tail := 0, len(in)-1
	for i = 1; i <= tail; {
		if in[i] > mid {
			in[i], in[tail] = in[tail], in[i]
			tail--
		} else {
			in[i], in[head] = in[head], in[i]
			head++
			i++
		}
	}
	in[head] = mid
	QuickSort(in[:head])
	QuickSort(in[head+1:])
}
