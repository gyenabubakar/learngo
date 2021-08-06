package main

type list []printer

func (l list) print() {
	for _, item := range l {
		item.print()
	}
}

func (l list) discount(ratio float64) {
	for _, item := range l {
		item.discount(ratio)
	}
}