package energy

import (
	"fmt"
	"time"
)

type Electricity struct {

}

func (e *Electricity) GetElectricity() {
	fmt.Printf("Getting Electricity...\n")
	time.Sleep(1 * time.Second)
}

func NewElectricity() *Electricity {
	return &Electricity{}
}