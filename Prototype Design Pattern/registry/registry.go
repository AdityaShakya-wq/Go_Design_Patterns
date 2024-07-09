package registry

import (
	"ProtoTypeDesignPattern/service"
)

var RegistryList = make(map[int]service.Shape)

func LoadRegistry() {
	circle := service.NewCircle(2, 3, 4)
	RegistryList[int(circle.GetId())] = circle

	square := service.NewSquare(3, 4)
	RegistryList[int(square.GetId())] = square
}
