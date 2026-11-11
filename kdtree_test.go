package kdtree

import (
	"testing"
	"github.com/njaros/KdTree/point"
)

func TestKdTreeUnsorted(t *testing.T) {
	tree := NewKdTreeUnsorted(nil);
	tree.Insert(point.NewPoint2D(2, 3, nil));
	tree.Insert(point.NewPoint2D(4, 4, nil));
	tree.Insert(point.NewPoint2D(8, 8, nil));
	tree.Insert(point.NewPoint2D(-2, 3, nil));
	tree.Insert(point.NewPoint2D(-2, 3, nil));
	tree.Insert(point.NewPoint2D(2, 3, nil));
	tree.Insert(point.NewPoint2D(2, 3, nil));
	expected := "[(2.00, 3.00) [(-2.00, 3.00) [(-2.00, 3.00) [] [(2.00, 3.00) [(2.00, 3.00)] []]] []] [(4.00, 4.00) [] [(8.00, 8.00)]]]";
	result := tree.ToString();
	if result != expected {
		t.Errorf("Got '%s', expected '%s'", result, expected);
	}
}

func TestNewKdTree(t *testing.T) {
	tree := NewKdTree([]Point{
		point.NewPoint2D(1.2, 3.1, nil),
		point.NewPoint2D(1.1, 8.1, nil),
		point.NewPoint2D(0.2, 31.1, nil),
		point.NewPoint2D(-1.2, -3.1, nil),
		point.NewPoint2D(1.21, 13.1, nil),
		point.NewPoint2D(11.2, 43.1, nil),
		point.NewPoint2D(7., -8.1, nil),
		point.NewPoint2D(17.2, -13.1, nil),
		point.NewPoint2D(5.2, 4, nil),
		point.NewPoint2D(8, 8, nil),
		point.NewPoint2D(1, 1, nil),
	});
	expected := "[(1.21, 13.10) [(1.20, 3.10) [(1.00, 1.00) [(-1.20, -3.10)] []] [(1.10, 8.10) [(0.20, 31.10)] []]] [(5.20, 4.00) [(17.20, -13.10) [(7.00, -8.10)] []] [(11.20, 43.10) [(8.00, 8.00)] []]]]";
	found := tree.ToString();
	if found != expected {
		t.Errorf("expected: '%s', found '%s'\n", expected, found);
	}
}