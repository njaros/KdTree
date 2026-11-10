package kdtree

import (
	"testing"
	"github.com/njaros/KdTree/point"
)

func TestKdTree(t *testing.T) {
	tree := NewKdTreeUnsorted(nil);
	tree.Insert(point.NewPoint2D(2, 3, nil));
	tree.Insert(point.NewPoint2D(4, 4, nil));
	tree.Insert(point.NewPoint2D(8, 8, nil));
	tree.Insert(point.NewPoint2D(-2, 3, nil));
	tree.Insert(point.NewPoint2D(-2, 3, nil));
	tree.Insert(point.NewPoint2D(2, 3, nil));
	tree.Insert(point.NewPoint2D(2, 3, nil));
	result := tree.ToString();
	want := "[(2.00, 3.00) [(-2.00, 3.00) [(-2.00, 3.00) [] [(2.00, 3.00) [(2.00, 3.00)] []]] []] [(4.00, 4.00) [] [(8.00, 8.00)]]]";
	if result != want {
		t.Errorf("Got '%s', expected '%s'", result, want);
	}
}