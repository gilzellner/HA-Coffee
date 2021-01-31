package main

import (
	"github.com/gilzellner/HA-Coffee/pkg/beans"
	"github.com/gilzellner/HA-Coffee/pkg/coffeemachine"
	"github.com/gilzellner/HA-Coffee/pkg/energy"
	"github.com/gilzellner/HA-Coffee/pkg/water"
)

func main() {
	b := beans.New()
	w := water.New()
	//e := energy.NewElectricity()
	e := energy.NewGas()
	s := energy.NewStove(e)
	c := coffeemachine.NewBialetti(b, w, s)
	c.GetCoffee()
}
