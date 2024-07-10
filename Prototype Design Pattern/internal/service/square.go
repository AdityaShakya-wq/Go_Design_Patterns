package service

import (
	"ProtoTypeDesignPattern/models"
	"fmt"
)

type Square struct {
	Id      models.ShapeType
	Length  int
	Breadth int
}

func NewSquare(length, breadth int) Square {
	return Square{models.SquareType, length, breadth}
}

func (s Square) GetId() models.ShapeType {
	return s.Id
}

func (s Square) PrintProtoType() {
	fmt.Printf("Square Properties. Length -> %d, Breadth -> %d\n", s.Length, s.Breadth)
}

func (s Square) Clone() Shape {
	return NewSquare(s.Length, s.Breadth)
}
