package main

import "fmt"

type game struct {
	title string
	price money
}

func (g *game) print() {
	fmt.Printf("%s costs %s\n", g.title, g.price.format())
}

func (g *game) discount(ratio float64) {
	if ratio > 1 {
		ratio = 1
	} else if ratio < 0 {
		ratio = 0
	}

	g.price -= g.price * money(ratio)
}