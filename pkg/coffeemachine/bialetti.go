package coffeemachine

import (
	"fmt"
	"time"
)

type Bialetti struct {
	b Beans
	w Water
	e Heat
}

type Heat interface {
	GetElectricity()
}

func NewBialetti(b Beans, w Water, e Electricity) *DelonghiMagnificaS {
	return &DelonghiMagnificaS{
		b: b,
		w: w,
		e: e,
	}
}


func (d *Bialetti) brewCoffee() {
	fmt.Printf("Brewing Coffee...\n")
	time.Sleep(1 * time.Second)
}

func (d *Bialetti) GetCoffee() {
	d.e.GetElectricity()
	d.b.GetBeans()
	d.w.GetWater()
	fmt.Printf("Brewing Coffee...\n")
	time.Sleep(1 * time.Second)
	fmt.Printf(
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