package kdtree

import (
	"fmt"
	"log"
)

type Point interface {
	Content() interface{}
	Dimension() int
	Axis(i int) float64
}

type KdTree struct {
	root *kdnode
}

// Usefull for a small amount of points (like thousands)
// And when the kdtree needs to be implemented often.
func NewKdTreeUnsorted(pts []Point) *KdTree {
	kd := KdTree{nil};

	for _, p := range pts {
		kd.Insert(p);
	}
	return &kd;
}

// Usefull for a big amount of points
func NewKdTree(pts []Point) *KdTree {
	kd := KdTree{nil};

	//TODO
	return &kd;
}

func (t *KdTree) Insert(p Point) {
	if t.root != nil && p.Dimension() != t.root.point.Dimension() {
		log.Fatalf(
			"Fatal: wrong dimension point. Tried to insert dim %d unstead of %d\n",
			p.Dimension(),
			t.root.point.Dimension(),
		);
	}
	depth := 0;
	current := &t.root;
	for ;*current != nil; {
		current = (*current).next(p);
		depth += 1;
	}
	*current = newKdNode(p, depth);
}

type kdnode struct {
	point Point
	left *kdnode
	right *kdnode
	depth int
}

func newKdNode(p Point, depth int) *kdnode {
	return &kdnode{
		point: p,
		left: nil,
		right: nil,
		depth: depth,
	};
}

func (n *kdnode) next(p Point) **kdnode {
	axis := n.depth % p.Dimension();
	if p.Axis(axis) <= n.point.Axis(axis) {
		return &n.left;
	}
	return &n.right;
}

func (t KdTree) ToString() string {
	toString := "[";
	if t.root != nil {
		t.root.toString(&toString);
	}
	toString += "]";

	return toString;
}

func (t KdTree) Print() {
	to_print := t.ToString();
	log.Print(to_print);
}

func (n *kdnode) toString(to_print *string) {
	*to_print += "(";
	for axis := range n.point.Dimension() {
		*to_print += fmt.Sprintf("%.2f", n.point.Axis(axis))
		if axis < n.point.Dimension() - 1 {
			*to_print += ", ";
		}
	}
	*to_print += ")";
	if (n.left != nil || n.right != nil) {
		*to_print += " [";
		if n.left != nil {
			n.left.toString(to_print);
		}
		*to_print += "] ";
		*to_print += "[";
		if n.right != nil {
			n.right.toString(to_print);
		}
		*to_print += "]";
	}
}