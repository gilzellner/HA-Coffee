package beans

import "fmt"

type Beans struct {

}

func (b *Beans) GetBeans() {
	_ = fmt.Sprint("Getting Beans")
}