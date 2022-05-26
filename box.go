package golang_united_school_homework

import (
	"errors"
	"fmt"
)

var (
	errorOutOfRange      = errors.New("Box is full.Out of shapeCapacity.")
	errorIndexOutOfRange = errors.New("Shape by this index doesn't exist or index went out of the range.")
)

// box contains list of shapes and able to perform operations on them
type box struct {
	shapes         []Shape
	shapesCapacity int // Maximum quantity of shapes that can be inside the box.
}

func (b *box) checkIsFull() (bool, error) {

	if b.shapesCapacity-len(b.shapes) >= 1 {
		return false, nil
	}

	return true, errorOutOfRange
}

func (b *box) findByIndex(i int) (found bool, err error) {
	if i < len(b.shapes) {
		return true, nil
	}

	return false, errorIndexOutOfRange
}

// NewBox creates new instance of box
func NewBox(shapesCapacity int) *box {
	return &box{
		shapesCapacity: shapesCapacity,
	}
}

// AddShape adds shape to the box
// returns the error in case it goes out of the shapesCapacity range.
func (b *box) AddShape(shape Shape) error {

	_, err := b.checkIsFull()

	if err != nil {
		return err
	}

	b.shapes = append(b.shapes, shape)

	return nil
}

// GetByIndex allows getting shape by index
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) GetByIndex(i int) (Shape, error) {

	_, err := b.findByIndex(i)
	if err != nil {
		return nil, err
	}

	return b.shapes[i], nil

}

// ExtractByIndex allows getting shape by index and removes this shape from the list.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ExtractByIndex(i int) (Shape, error) {
	_, err := b.findByIndex(i)
	if err != nil {
		return nil, err
	}

	extShape := b.shapes[i]

	//remove form shapes
	b.shapes = append(b.shapes[:i], b.shapes[i+1:]...)

	return extShape, nil

}

// ReplaceByIndex allows replacing shape by index and returns removed shape.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ReplaceByIndex(i int, shape Shape) (Shape, error) {
	_, err := b.findByIndex(i)
	if err != nil {
		return nil, err
	}

	replacedShape := b.shapes[i]
	b.shapes[i] = shape

	return replacedShape, nil

}

// SumPerimeter provides sum perimeter of all shapes in the list.
func (b *box) SumPerimeter() float64 {

	var sum float64
	for i := range b.shapes {
		sum += b.shapes[i].CalcPerimeter()
	}

	return sum

}

// SumArea provides sum area of all shapes in the list.
func (b *box) SumArea() float64 {

	var sum float64
	for i := range b.shapes {
		sum += b.shapes[i].CalcArea()
	}

	return sum

}

// RemoveAllCircles removes all circles in the list
// whether circles are not exist in the list, then returns an error
func (b *box) RemoveAllCircles() error {

	circleCount := 0

	//Here we need to iterate from back to front so you don't have to worry about indexes that are deleted.
	for i := len(b.shapes) - 1; i >= 0; i-- {
		if _, ok := b.shapes[i].(*Circle); ok {
			circleCount++
			//remove that shape from slice
			b.shapes = append(b.shapes[:i], b.shapes[i+1:]...)

		}
	}

	if circleCount == 0 {
		return fmt.Errorf("No circles found.")
	}

	return nil

}
