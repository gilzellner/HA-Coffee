package beans

import (
	"fmt"
	"time"
)

type Beans struct {

}

func (b *Beans) GetBeans() {
	_ = fmt.Sprint("Getting Beans")
	time.Sleep(1 * time.Second)
}