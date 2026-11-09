package point

type Point struct {
	position []float64
	content interface{}
}

func NewPoint(position []float64, content interface{}) *Point {
	return &Point{
		position: position,
		content: content,
	}
}

func NewPoint2D(x float64, y float64, content interface{}) *Point {
	return NewPoint([]float64{x, y}, content);
}

func NewPoint3D(x float64, y float64, z float64, content interface{}) *Point {
	return NewPoint([]float64{x, y, z}, content);
}

func (p Point) Dimension() int {
	return len(p.position);
}

func (p Point) Axis(i int) float64 {
	return p.position[i];
}

func (p Point) Content() interface{} {
	return p.content;
}