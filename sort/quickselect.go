package sort

type Point interface {
	Axis(i int) float64
}

func GetMedian[T Point](pts []T, axis int) T {
	k := len(pts) / 2
	return QuickSelect(pts, k, axis)[k]
}

func GetMedianPartitions[T Point](pts []T, axis int) ([]T, T, []T) {
	k := len(pts) / 2
	return QuickSelectPartitions(pts, k, axis)
}

// Do a QuickSelect, then returns elements before k, the k'st element, and the elements after k.
func QuickSelectPartitions[T Point](pts []T, k int, axis int) ([]T, T, []T) {
	pts = QuickSelect(pts, k, axis)
	return pts[0:k], pts[k], pts[k+1:]
}

// Organise array to put the k'st shortest element in position k
func QuickSelect[T Point](pts []T, k int, axis int) []T {
	if len(pts) == 0 {
		return pts
	}
	start := 0
	left := 0
	i := 0
	finish := false
	for pivot := len(pts) - 1; !finish; {
		pivot_value := pts[pivot].Axis(axis)
		for ; i != pivot; i++ {
			if pts[i].Axis(axis) <= pivot_value {
				if i != left {
					swap(&pts[i], &pts[left])
				}
				left++
			}
		}
		swap(&pts[left], &pts[pivot])
		if left > k {
			pivot = left - 1
		} else if left < k {
			start = left + 1
		} else {
			finish = true
		}
		left = start
		i = start
	}
	return pts
}

func swap[T Point](a *T, b *T) {
	temp := *a
	*a = *b
	*b = temp
}
