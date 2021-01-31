package water

import (
	"fmt"
	"time"
)

type Tap struct {

}

func (t *Tap) GetWater() {
	fmt.Printf("Getting Water...\n")
	time.Sleep(1 * time.Second)
}

func New() *Tap {
	return &Tap{}
}