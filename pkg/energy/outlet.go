package energy

import "fmt"

type Electricity struct {

}

func (e *Electricity) GetPower() {
	_ = fmt.Sprint("Getting Electricity")
}