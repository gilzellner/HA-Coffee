package water

import "fmt"

type Tap struct {

}

func (t *Tap) GetWater() {
	_ = fmt.Sprint("Getting Electricity")
}