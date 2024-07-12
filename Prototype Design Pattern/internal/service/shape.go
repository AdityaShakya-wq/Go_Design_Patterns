package service

import (
	"ProtoTypeDesignPattern/internal/models"
)

type Shape interface {
	GetId() models.ShapeType //get id
	PrintProtoType()         // print
	Clone() Shape            // used for getting Deepcopy
}
