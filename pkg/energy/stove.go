package energy

import (
	"fmt"
	"time"
)

type Stove struct {
	g gas
}

type gas interface {
	GetGas()
}


func (s *Stove) GetHeat() {
	fmt.Printf("Heating up...\n")
	time.Sleep(1 * time.Second)
}

func NewStove(g gas) *Stove {
	return &Stove{g: g}
}