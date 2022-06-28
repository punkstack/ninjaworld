package ninja_world

const INITIALENERGY = 10

type Otsutsuki struct {
	name           string
	isTrapped      bool
	isAlive        bool
	energy         int
	currentVillage Village
}

func NewOtsutsuki(name string) *Otsutsuki {
	return &Otsutsuki{
		name:      name,
		isAlive:   true,
		isTrapped: false,
		energy:    INITIALENERGY,
	}
}

func (o *Otsutsuki) updateOtsutsuki(village Village) *Otsutsuki {
	o.energy -= 1
	o.currentVillage = village
	if o.energy == 0 {
		o.isAlive = false
	}
	return o
}
