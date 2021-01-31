package energy

import (
	"fmt"
	"time"
)

type Gas struct {

}

func (e *Gas) GetGas() {
	fmt.Printf("Getting Gas...\n")
	time.Sleep(1 * time.Second)
}

func NewGas() *Gas {
	return &Gas{}
}