package registry

import (
	service2 "ProtoTypeDesignPattern/internal/service"
)

var RegistryList = make(map[int]service2.Shape)

func LoadRegistry() {
	circle := service2.NewCircle(2, 3, 4)
	RegistryList[int(circle.GetId())] = circle

	square := service2.NewSquare(3, 4)
	RegistryList[int(square.GetId())] = square
}
