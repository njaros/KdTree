package sort

type Point interface {
	Axis(i int) float64
}

func GetMedian(pts []Point, axis int) Point {
	k := len(pts)/2;
	return QuickSelect(pts, k, axis)[k];
}

// Do a QuickSelect, then returns elements before k, the k'st element, and the elements after k.
func QuickSelectPartitions(pts []Point, k int, axis int) ([]Point, Point, []Point) {
	if len(pts) == 0 {
		return nil, nil, nil;
	}
	pts = QuickSelect(pts, k, axis);
	return pts[0:k], pts[k], pts[k+1:]
}

// Organise array to put the k'st shortest element in position k
func QuickSelect(pts []Point, k int, axis int) []Point {
	if len(pts) == 0 {
		return pts;
	}
	p_left := 0;
	p_right := len(pts) - 1;
	left := 0;
	i := 0;
	for pivot := p_right ;pivot != k; {
		pivot_value := pts[pivot].Axis(axis);
		for ;i != pivot; i++ {
			if (pts[i].Axis(axis) < pivot_value) {
				swap(&pts[i], &pts[left]);
				left++;
			}
		}
		swap(&pts[left], &pts[pivot]);
		if pivot > k {
			p_right = pivot - 1;
		} else {
			p_left = pivot + 1;
		}
		left = p_left;
		i = left;
		pivot = p_right;
	}
	return pts;
}

func swap(a *Point, b *Point) {
	temp := *a;
	*a = *b;
	*b = temp;
}