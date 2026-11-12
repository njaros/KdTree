package kdtree

import (
	"reflect"
	"sync"
	"testing"

	"github.com/njaros/KdTree/point"
)

// "Fixture" instances.

var kdTreeSingleInstance *KdTree
var lock1 = &sync.Mutex{}
var kdTreeUnsortedInstance *KdTree
var lock2 = &sync.Mutex{}

func getKdTreeInstance() *KdTree {
	if kdTreeSingleInstance == nil {
		lock1.Lock()
		defer lock1.Unlock()
		if kdTreeSingleInstance == nil {
			kdTreeSingleInstance = NewKdTree([]Point{
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
			})
		}
	}

	return kdTreeSingleInstance
}

func getKdTreeUnsortedInstance() *KdTree {
	if kdTreeUnsortedInstance == nil {
		lock2.Lock()
		defer lock2.Unlock()
		if kdTreeUnsortedInstance == nil {
			kdTreeUnsortedInstance = NewKdTreeUnsorted([]Point{
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
			})
		}
	}

	return kdTreeUnsortedInstance
}

// Tests

func TestKdTreeUnsorted(t *testing.T) {
	tree := NewKdTreeUnsorted(nil)
	tree.Insert(point.NewPoint2D(2, 3, nil))
	tree.Insert(point.NewPoint2D(4, 4, nil))
	tree.Insert(point.NewPoint2D(8, 8, nil))
	tree.Insert(point.NewPoint2D(-2, 3, nil))
	tree.Insert(point.NewPoint2D(-2, 3, nil))
	tree.Insert(point.NewPoint2D(2, 3, nil))
	tree.Insert(point.NewPoint2D(2, 3, nil))
	expected := "[(2.00, 3.00) [(-2.00, 3.00) [(-2.00, 3.00) [] [(2.00, 3.00) [(2.00, 3.00)] []]] []] [(4.00, 4.00) [] [(8.00, 8.00)]]]"
	result := tree.ToString()
	if result != expected {
		t.Errorf("Got '%s', expected '%s'", result, expected)
	}
}

func TestNewKdTree(t *testing.T) {
	tree := getKdTreeInstance()
	expected := "[(1.21, 13.10) [(1.20, 3.10) [(1.00, 1.00) [(-1.20, -3.10)] []] [(1.10, 8.10) [(0.20, 31.10)] []]] [(5.20, 4.00) [(17.20, -13.10) [(7.00, -8.10)] []] [(11.20, 43.10) [(8.00, 8.00)] []]]]"
	found := tree.ToString()
	if found != expected {
		t.Errorf("expected: '%s', found '%s'\n", expected, found)
	}
}

func TestNearest(t *testing.T) {
	tree := getKdTreeInstance()
	testPoints := []Point{
		point.NewPoint2D(8, 8, nil),
		point.NewPoint2D(1.21, 13.10, nil),
		point.NewPoint2D(2, 2, nil),
		point.NewPoint2D(-1000, -1000, nil),
		point.NewPoint2D(20, -15, nil),
	}

	expected := []Point{
		point.NewPoint2D(8, 8, nil),
		point.NewPoint2D(1.21, 13.1, nil),
		point.NewPoint2D(1.2, 3.1, nil),
		point.NewPoint2D(-1.2, -3.1, nil),
		point.NewPoint2D(17.2, -13.1, nil),
	}

	for i := range 5 {
		found := tree.Nearest(testPoints[i])
		if !reflect.DeepEqual(*found, expected[i]) {
			t.Errorf("%d: with testPoint %v -> expected '%v', found '%v'", i, testPoints[i], expected[i], *found)
		}
	}
}

func TestNearestUnsorted(t *testing.T) {
	tree := getKdTreeUnsortedInstance()
	testPoints := []Point{
		point.NewPoint2D(8, 8, nil),
		point.NewPoint2D(1.21, 13.10, nil),
		point.NewPoint2D(2, 2, nil),
		point.NewPoint2D(-1000, -1000, nil),
		point.NewPoint2D(20, -15, nil),
	}

	expected := []Point{
		point.NewPoint2D(8, 8, nil),
		point.NewPoint2D(1.21, 13.1, nil),
		point.NewPoint2D(1.2, 3.1, nil),
		point.NewPoint2D(-1.2, -3.1, nil),
		point.NewPoint2D(17.2, -13.1, nil),
	}

	for i := range 5 {
		found := tree.Nearest(testPoints[i])
		if !reflect.DeepEqual(*found, expected[i]) {
			t.Errorf("%d: with testPoint %v -> expected '%v', found '%v'", i, testPoints[i], expected[i], *found)
		}
	}
}
