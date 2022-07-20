package Builder

type Drictor struct {
	builder Builder
}

func NewConstruct(b Builder) *Drictor {
	return &Drictor{
		builder: b,
	}
}

func (this *Drictor) Consturct() {
	this.builder.buildDisk()
	this.builder.buildCPU()
	this.builder.buildRom()
}
