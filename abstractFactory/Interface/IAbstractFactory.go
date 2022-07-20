package Interface

//AbstractFactory 抽象工厂接口
type AbstractFactory interface {
	GetShape(shapeName string) Shape
	GetColor(colorName string) Color
}
