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
	e := energy.New()
	c := coffeemachine.New(b, w, e)
	c.GetCoffee()
}
