package energy

import (
	"fmt"
	"time"
)

type Electricity struct {

}

func (e *Electricity) GetPower() {
	_ = fmt.Sprint("Getting Electricity")
	time.Sleep(1 * time.Second)
}