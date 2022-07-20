package factory

import (
	"designMode/abstractFactory/Interface"
	"designMode/abstractFactory/shape"
)

type ShapeFactory struct{}

func (sh ShapeFactory) GetShape(shapeName string) Interface.Shape {
	switch shapeName {
	case "circle":
		return &shape.Circle{}
	case "square":
		return &shape.Square{}
	}
	return nil
}
func (sh ShapeFactory) GetColor(colorName string) Interface.Color {
	return nil
}
