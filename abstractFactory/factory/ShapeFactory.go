package factory

import (
	"designMode/abstractFactory/Interface"
	"designMode/abstractFactory/color"
)

type ColorFactory struct{}

func (cf ColorFactory) GetColor(colorName string) Interface.Color {
	switch colorName {
	case "red":
		return &color.Red{}
	case "green":
		return &color.Green{}
	default:
		return nil
	}
}
func (cf ColorFactory) GetShape(shapeName string) Interface.Shape {
	return nil
}
