package main

import "fmt"

type book struct {
	title, author string
	price         money
}

func (b *book) print() {
	fmt.Printf("%s by %s costs %s\n", b.title, b.author, b.price.format())
}

func (b *book) discount(ratio float64) {
	if ratio > 1 {
		ratio = 1
	} else if ratio < 0 {
		ratio = 0
	}

	b.price -= b.price * money(ratio)
}
