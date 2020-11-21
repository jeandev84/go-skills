package main

import (
	"fmt"
)

type Instrumenter interface {
	Play()
}

type Amplifier interface {
	Connect(amp string)
}

type Guitar struct {
	Strings int
}

func (g Guitar) Play() {
	fmt.Printf("Tzouuuuuinng with %d strings\n", g.Strings)
}

func (g Guitar) Connect(amp string) {
	fmt.Printf("Connected to %v\n", amp)
}

type Piano struct {
	Keys int
}

func (p Piano) Play() {
	fmt.Printf("Plip Plip %d keys\n", p.Keys)
}

func main() {
	var instr Instrumenter

	instr = &Guitar{5}
	instr.Play()
	// instr.Connect("Marshall") // instr is *NOT* of type Amplifier!

	instr = &Piano{88}
	instr.Play()

	g := Guitar{12}
	g.Play()
	g.Connect("Marshall")
}
