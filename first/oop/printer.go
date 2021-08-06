package main

type printer interface {
	print()
	discount(ratio float64)
}