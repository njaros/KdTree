package main

import (
	"github.com/njaros/KdTree/point"
	"github.com/njaros/KdTree"
)

func main() {
	tree := kdtree.NewKdTree();
	tree.Insert(point.NewPoint2D(2, 3, nil));
	tree.Insert(point.NewPoint2D(4, 4, nil));
	tree.Insert(point.NewPoint2D(8, 8, nil));
	tree.Insert(point.NewPoint2D(-2, 3, nil));
	tree.Insert(point.NewPoint2D(-2, 3, nil));
	tree.Insert(point.NewPoint2D(2, 3, nil));
	tree.Insert(point.NewPoint2D(2, 3, nil));
	tree.Print();
}