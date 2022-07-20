package main

import "designMode/FacadePattern/dao"

func main() {
	sf := dao.NewSportFacade()
	sf.PlayBasketball()
	sf.PlayFootball()
}
