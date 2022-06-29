package ninja_world

import (
	"testing"
)

func newOtsutsuki(name string) *Otsutsuki {
	return NewOtsutsuki(name)
}

func TestOtsutsuki_KillOtsutsuki(t *testing.T) {
	o := newOtsutsuki("test1")
	o.KillOtsutsuki()

	if o.isAlive == true {
		t.Errorf("not able to kill otsutsuki")
	}

	if o.energy > 0 {
		t.Errorf("not able to kill otsutsuki")
	}
}

func TestOtsutsuki_moveOtsutsuki(t *testing.T) {
	o := newOtsutsuki("test1")
	initialEnergy := o.energy
	o.moveOtsutsuki(NewVillage("test1"))

	if o.energy == initialEnergy {
		t.Errorf("otsutsuki didn't reduced")
	}

	if o.currentVillage.name != "test1" {
		t.Errorf("otsutsuki didn't reduced")
	}
}
