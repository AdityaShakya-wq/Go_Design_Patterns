package main

import (
	"ProtoTypeDesignPattern/models"
	"ProtoTypeDesignPattern/registry"
	"ProtoTypeDesignPattern/service"
	"fmt"
)

func main() {
	registry.LoadRegistry()

	square := registry.RegistryList[int(models.SquareType)]
	sq, sok := square.(service.Square) // if type can be asserted
	if sok {
		fmt.Println("Old square properties:")
		sq.PrintProtoType()
		newSquare := sq.Clone()
		fmt.Println("Cloned square object properties:")
		newSquare.PrintProtoType()
	}

	circle := registry.RegistryList[int(models.CircleType)]
	c, cok := circle.(service.Circle)
	if cok {
		fmt.Println("Old circle properties")
		c.PrintProtoType()
		newCircle := c.Clone().(service.Circle) // type assertion to modify data
		newCircle.Radius = 6
		fmt.Println("Cloned circle object with modified properties:")
		newCircle.PrintProtoType()
	}
}
