package dao

type SportFacade struct {
	basketball Basketball
	football   Football
}

func NewSportFacade() *SportFacade {
	return &SportFacade{
		basketball: Basketball{},
		football:   Football{},
	}
}

func (f *SportFacade) PlayBasketball() {
	f.basketball.Run()
}

func (f *SportFacade) PlayFootball() {
	f.football.Run()
}
