package service

import (
	"ProtoTypeDesignPattern/models"
	"fmt"
)

type Circle struct {
	Id            models.ShapeType
	Radius        int
	Diameter      int
	Circumference int
}

func NewCircle(radius, diameter, circumference int) Circle {
	return Circle{models.CircleType, radius, diameter, circumference}
}

func (c Circle) GetId() models.ShapeType {
	return c.Id
}

func (c Circle) PrintProtoType() {
	fmt.Printf("Circle Properties. Radius -> %d, Diameter -> %d, Circumference -> %d\n", c.Radius, c.Diameter, c.Circumference)
}

func (c Circle) Clone() Shape {
	return NewCircle(c.Radius, c.Diameter, c.Circumference)
}
