package water

import (
	"fmt"
	"time"
)

type Tap struct {

}

func (t *Tap) GetWater() {
	_ = fmt.Sprint("Getting Electricity")
	time.Sleep(1 * time.Second)
}