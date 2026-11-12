package kdtree

import (
	"fmt"
	"log"

	"github.com/njaros/KdTree/sort"
)

type Point interface {
	Content() interface{}
	Dimension() int
	Axis(i int) float64
}

type KdTree struct {
	root *kdnode
	size int
}

// Usefull for a small amount of points (like thousands)
// And when the kdtree needs to be implemented often.
func NewKdTreeUnsorted(pts []Point) *KdTree {
	kd := KdTree{nil, 0}

	for _, p := range pts {
		kd.Insert(p)
	}
	return &kd
}

// Usefull for a big amount of points
func NewKdTree(pts []Point) *KdTree {
	kd := KdTree{nil, len(pts)}

	l, m, r := sort.GetMedianPartitions(pts, 0)
	kd.root = newKdNode(m, 0)
	kd.root.insertPoints(l, r)
	return &kd
}

func (t *KdTree) Insert(p Point) {
	if t.root != nil && p.Dimension() != t.root.point.Dimension() {
		log.Fatalf(
			"Fatal: wrong dimension point. Tried to insert dim %d unstead of %d\n",
			p.Dimension(),
			t.root.point.Dimension(),
		)
	}
	t.size += 1
	depth := 0
	current := &t.root
	for *current != nil {
		current = (*current).next(p)
		depth += 1
	}
	*current = newKdNode(p, depth)
}

type kdnode struct {
	point Point
	left  *kdnode
	right *kdnode
	depth int
}

func newKdNode(p Point, depth int) *kdnode {
	return &kdnode{
		point: p,
		left:  nil,
		right: nil,
		depth: depth,
	}
}

func (n *kdnode) next(p Point) **kdnode {
	axis := n.depth % p.Dimension()
	if p.Axis(axis) <= n.point.Axis(axis) {
		return &n.left
	}
	return &n.right
}

func (n *kdnode) insertPoints(ptsLeft []Point, ptsRight []Point) {
	depth := n.depth + 1
	axis := depth % n.point.Dimension()
	if len(ptsLeft) > 0 {
		l, m, r := sort.GetMedianPartitions(ptsLeft, axis)
		n.left = newKdNode(m, depth)
		n.left.insertPoints(l, r)
	}
	if len(ptsRight) > 0 {
		l, m, r := sort.GetMedianPartitions(ptsRight, axis)
		n.right = newKdNode(m, depth)
		n.right.insertPoints(l, r)
	}
}

func (t KdTree) ToString() string {
	toString := "["
	if t.root != nil {
		t.root.toString(&toString)
	}
	toString += "]"

	return toString
}

func (t KdTree) Print() {
	to_print := t.ToString()
	log.Print(to_print)
}

func (t KdTree) Nearest(p Point) *Point {
	if t.root == nil {
		return nil
	}

	bestDist := squareDistPoints(p, t.root.point);
	bestPoint := t.root.point;
	t.root.nearest(p, &bestPoint, &bestDist);
	return &bestPoint;
}

func (n *kdnode) nearest(p Point, best *Point, bestDist *float64) {
	axis := n.depth % n.point.Dimension()
	leftSide := n.point.Axis(axis) >= p.Axis(axis)

	if leftSide {
		if n.left != nil {
			nextNodePoint := n.left.point;
			nextDist := squareDistPoints(p, nextNodePoint);
			if nextDist < *bestDist {
				*best = nextNodePoint;
				*bestDist = nextDist;
			}
			n.left.nearest(p, best, bestDist)
		}
		if n.right != nil && squareDistHyperplan(axis, n.point, p) < *bestDist {
			nextNodePoint := n.right.point;
			nextDist := squareDistPoints(p, nextNodePoint);
			if nextDist < *bestDist {
				*best = nextNodePoint;
				*bestDist = nextDist;
			}
			n.right.nearest(p, best, bestDist)
		}
	} else {
		if n.right != nil {
			nextNodePoint := n.right.point;
			nextDist := squareDistPoints(p, nextNodePoint);
			if nextDist < *bestDist {
				*best = nextNodePoint;
				*bestDist = nextDist;
			}
			n.right.nearest(p, best, bestDist)
		}
		if n.left != nil && squareDistHyperplan(axis, n.point, p) < *bestDist {
			nextNodePoint := n.left.point;
			nextDist := squareDistPoints(p, nextNodePoint);
			if nextDist < *bestDist {
				*best = nextNodePoint;
				*bestDist = nextDist;
			}
			n.left.nearest(p, best, bestDist)
		}
	}
}

func (t KdTree) GetAllInRange(p Point, r int) []*Point {
	//TODO
	return nil
}

func (t KdTree) KNN(p Point, n int) []*Point {
	//TODO
	return nil
}

func squareDistPoints(a Point, b Point) float64 {
	distance := 0.
	for i := range a.Dimension() {
		diff := a.Axis(i) - b.Axis(i)
		distance += diff * diff
	}

	return distance
}

func squareDistHyperplan(axis int, treePoint Point, p Point) float64 {
	d := treePoint.Axis(axis) - p.Axis(axis)
	return d * d
}

func (n *kdnode) toString(to_print *string) {
	*to_print += "("
	for axis := range n.point.Dimension() {
		*to_print += fmt.Sprintf("%.2f", n.point.Axis(axis))
		if axis < n.point.Dimension()-1 {
			*to_print += ", "
		}
	}
	*to_print += ")"
	if n.left != nil || n.right != nil {
		*to_print += " ["
		if n.left != nil {
			n.left.toString(to_print)
		}
		*to_print += "] "
		*to_print += "["
		if n.right != nil {
			n.right.toString(to_print)
		}
		*to_print += "]"
	}
}
