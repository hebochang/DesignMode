package Builder

type Builder interface {
	buildDisk()
	buildCPU()
	buildRom()
}
