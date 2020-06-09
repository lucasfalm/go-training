package main

type printer interface {
	print()
}

type list []printer
