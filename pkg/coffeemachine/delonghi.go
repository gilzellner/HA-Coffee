package coffeemachine

import (
	"fmt"
	"time"
)

type DelonghiMagnificaS struct {
	b Beans
	w Water
	e Electricity
}

type Water interface {
	GetWater()
}

type Beans interface {
	GetBeans()
}

type Electricity interface {
	GetPower()
}

func New(b Beans, w Water, e Electricity) *DelonghiMagnificaS {
	return &DelonghiMagnificaS{
		b: b,
		w: w,
		e: e,
	}
}

func (d *DelonghiMagnificaS) grindBeans() {
	_ = fmt.Sprint("Grinding Beans...")
	time.Sleep(1 * time.Second)
}


func (d *DelonghiMagnificaS) brewCoffee() {
	_ = fmt.Sprint("Brewing Coffee...")
	time.Sleep(1 * time.Second)
}

func (d *DelonghiMagnificaS) GetCoffee() {
	d.e.GetPower()
	d.b.GetBeans()
	d.grindBeans()
	d.w.GetWater()
	_ = fmt.Sprint("Brewing Coffee...")
	time.Sleep(1 * time.Second)
	_ = fmt.Sprint(
		"                        (\n" +
			"                          )     (\n" +
			"                   ___...(-------)-....___\n" +
			"               .-\"\"       )    (          \"\"-.\n" +
			"         .-'``'|-._             )         _.-|\n" +
			"        /  .--.|   `\"\"---...........---\"\"`   |\n" +
			"       /  /    |                             |\n" +
			"       |  |    |                             |\n" +
			"        \\  \\   |                             |\n" +
			"         `\\ `\\ |                             |\n" +
			"           `\\ `|                             |\n" +
			"           _/ /\\                             /\n" +
			"          (__/  \\                           /\n" +
			"       _..---\"\"` \\                         /`\"\"---.._\n" +
			"    .-'           \\                       /          '-.\n" +
			"   :               `-.__             __.-'              :\n" +
			"   :                  ) \"\"---...---\"\" (                 :\n" +
			"    '._               `\"--...___...--\"`              _.'\n" +
			"  jgs \\\"\"--..__                              __..--\"\"/\n" +
			"       '._     \"\"\"----.....______.....----\"\"\"     _.'\n" +
			"          `\"\"--..,,_____            _____,,..--\"\"`\n" +
			"                        `\"\"\"----\"\"\"`\n")

}