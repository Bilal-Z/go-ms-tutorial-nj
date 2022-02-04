package handlers

import (
	"log"
)

type Products struct{
	l *log.Logger
}

func NewProduct(l *log.Logger) *Products{
	return &Products{l}
}

type KeyProduct struct{}