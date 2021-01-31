package beans

import (
	"fmt"
	"time"
)

type Beans struct {

}

func (b *Beans) GetBeans() {
	fmt.Printf("Getting Beans...\n")
	time.Sleep(1 * time.Second)
}

func New() *Beans {
	return &Beans{}
}