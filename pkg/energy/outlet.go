package energy

import (
	"fmt"
	"time"
)

type Electricity struct {

}

func (e *Electricity) GetPower() {
	fmt.Printf("Getting Electricity...\n")
	time.Sleep(1 * time.Second)
}

func New() *Electricity {
	return &Electricity{}
}