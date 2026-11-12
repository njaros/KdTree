package sort

import (
	"reflect"
	"testing"

	"github.com/njaros/KdTree/point"
)

func collectArrayFromPoints(pts []Point, axis int) []float64 {
	collect := []float64{}
	for _, p := range pts {
		collect = append(collect, p.Axis(axis))
	}

	return collect
}

func TestQuickSelectMedian(t *testing.T) {
	pts := []Point{
		point.NewPoint2D(12, 3, nil),
		point.NewPoint2D(-11, 3, nil),
		point.NewPoint2D(22, 3, nil),
		point.NewPoint2D(12, 3, nil),
		point.NewPoint2D(7, 3, nil),
		point.NewPoint2D(-15, 3, nil),
		point.NewPoint2D(0, 3, nil),
		point.NewPoint2D(4, 3, nil),
		point.NewPoint2D(1, 3, nil),
	}

	a, median, b := QuickSelectPartitions(pts, len(pts)/2, 0)

	collectA := collectArrayFromPoints(a, 0)
	collectB := collectArrayFromPoints(b, 0)

	expectedA := []float64{-11, -15, 0, 1}
	expectedB := []float64{12, 7, 12, 22}
	expectedMedian := point.NewPoint2D(4, 3, nil)
	if !reflect.DeepEqual(median, expectedMedian) {
		t.Errorf("expected median: %v, found: %v\n", expectedMedian, median)
	}
	if !reflect.DeepEqual(collectA, expectedA) {
		t.Errorf("expected left array: %v, found %v", expectedA, collectA)
	}
	if !reflect.DeepEqual(collectB, expectedB) {
		t.Errorf("expected left array: %v, found %v", expectedB, collectB)
	}
}

func TestQuickSelectFirst(t *testing.T) {
	pts := []Point{
		point.NewPoint2D(12, 3, nil),
		point.NewPoint2D(-11, 3, nil),
		point.NewPoint2D(22, 3, nil),
		point.NewPoint2D(12, 3, nil),
		point.NewPoint2D(7, 3, nil),
		point.NewPoint2D(-15, 3, nil),
		point.NewPoint2D(0, 3, nil),
		point.NewPoint2D(4, 3, nil),
		point.NewPoint2D(1, 3, nil),
	}

	a, first, b := QuickSelectPartitions(pts, 0, 0)

	collectA := collectArrayFromPoints(a, 0)
	collectB := collectArrayFromPoints(b, 0)

	expectedA := []float64{}
	expectedB := []float64{-11, 0, 1, 7, 12, 22, 4, 12}
	expectedFirst := point.NewPoint2D(-15, 3, nil)
	if !reflect.DeepEqual(first, expectedFirst) {
		t.Errorf("expected first: %v, found: %v\n", expectedFirst, first)
	}
	if !reflect.DeepEqual(collectA, expectedA) {
		t.Errorf("expected left array: %v, found %v", expectedA, collectA)
	}
	if !reflect.DeepEqual(collectB, expectedB) {
		t.Errorf("expected left array: %v, found %v", expectedB, collectB)
	}
}

func TestQuickSelectLast(t *testing.T) {
	pts := []Point{
		point.NewPoint2D(12, 3, nil),
		point.NewPoint2D(-11, 3, nil),
		point.NewPoint2D(22, 3, nil),
		point.NewPoint2D(12, 3, nil),
		point.NewPoint2D(7, 3, nil),
		point.NewPoint2D(-15, 3, nil),
		point.NewPoint2D(0, 3, nil),
		point.NewPoint2D(4, 3, nil),
		point.NewPoint2D(1, 3, nil),
	}

	a, median, b := QuickSelectPartitions(pts, len(pts)-1, 0)

	collectA := collectArrayFromPoints(a, 0)
	collectB := collectArrayFromPoints(b, 0)

	expectedA := []float64{-11, -15, 0, 1, 7, 12, 4, 12}
	expectedB := []float64{}
	expectedMedian := point.NewPoint2D(22, 3, nil)
	if !reflect.DeepEqual(median, expectedMedian) {
		t.Errorf("expected median: %v, found: %v\n", expectedMedian, median)
	}
	if !reflect.DeepEqual(collectA, expectedA) {
		t.Errorf("expected left array: %v, found %v", expectedA, collectA)
	}
	if !reflect.DeepEqual(collectB, expectedB) {
		t.Errorf("expected right array: %v, found %v", expectedB, collectB)
	}
}
