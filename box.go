package golang_united_school_homework

import "fmt"

// box contains list of shapes and able to perform operations on them
type box struct {
	shapes         []Shape
	shapesCapacity int // Maximum quantity of shapes that can be inside the box.
}

// NewBox creates new instance of box
//goland:noinspection GoUnusedExportedFunction
func NewBox(shapesCapacity int) *box {
	return &box{
		shapesCapacity: shapesCapacity,
	}
}

// AddShape adds shape to the box
// returns the error in case it goes out of the shapesCapacity range.
func (b *box) AddShape(shape Shape) error {
	if len(b.shapes) >= b.shapesCapacity {
		return fmt.Errorf("got out of capacity range")
	}
	b.shapes = append(b.shapes, shape)
	return nil
}

// GetByIndex allows getting shape by index
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) GetByIndex(i int) (Shape, error) {
	if i < 0 || i >= len(b.shapes) {
		return nil, fmt.Errorf("out of index range")
	}
	return b.shapes[i], nil
}

// ExtractByIndex allows getting shape by index and removes this shape from the list.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ExtractByIndex(i int) (Shape, error) {
	shape, err := b.GetByIndex(i)
	if err != nil {
		return nil, fmt.Errorf("shape does't exists: %w", err)
	}
	b.shapes = append(b.shapes[:i], b.shapes[i+1:]...)
	return shape, nil
}

// ReplaceByIndex allows replacing shape by index and returns removed shape.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ReplaceByIndex(i int, shape Shape) (Shape, error) {
	shape, err := b.GetByIndex(i)
	if err != nil {
		return nil, fmt.Errorf("shape does't exists: %w", err)
	}
	b.shapes[i] = shape
	return shape, nil
}

// SumPerimeter provides sum perimeter of all shapes in the list.
func (b *box) SumPerimeter() (sum float64) {
	for _, s := range b.shapes {
		sum += s.CalcPerimeter()
	}
	return
}

// SumArea provides sum area of all shapes in the list.
func (b *box) SumArea() (sum float64) {
	for _, s := range b.shapes {
		sum += s.CalcArea()
	}
	return
}

// RemoveAllCircles removes all circles in the list
// whether circles are not exist in the list, then returns an error
func (b *box) RemoveAllCircles() error {
	shapes := make([]Shape, 0)
	for i := range b.shapes {
		if _, ok := b.shapes[i].(*Circle); !ok {
			shapes = append(shapes, b.shapes[i])
		}
	}

	if len(b.shapes) > 0 && len(b.shapes) == len(shapes) {
		return fmt.Errorf("no circles exists")
	}

	b.shapes = shapes
	return nil
}
