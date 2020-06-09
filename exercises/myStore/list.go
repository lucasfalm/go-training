package main

type printer interface {
	print()
	discount(d int)
}

type list []printer
