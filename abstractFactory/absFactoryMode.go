package abstractFactory

type FactoryProducer struct{}

func (fp FactoryProducer) GetFactory(factoryname string) AbstractFactory {
	switch factoryname {
	case "color":
		return &ColorFactory{}
	case "shape":
		return &ShapeFactory{}
	default:
		return nil
	}
}
func NewFactoryProducer() *FactoryProducer {
	return &FactoryProducer{}
}
