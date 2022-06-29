package ninja_world

const INITIALENERGY = 10000

type OtsutsukiInterface interface {
	moveOtsutsuki(village *Village) *Otsutsuki
	KillOtsutsuki()
	IsMovable() bool
}

type Otsutsuki struct {
	name           string
	isAlive        bool
	energy         int
	currentVillage *Village
}

var _ OtsutsukiInterface = &Otsutsuki{}

func NewOtsutsuki(name string) *Otsutsuki {
	return &Otsutsuki{
		name:           name,
		isAlive:        true,
		energy:         INITIALENERGY,
		currentVillage: nil,
	}
}

// moveOtsutsuki update the status of otsutsuki
func (o *Otsutsuki) moveOtsutsuki(village *Village) *Otsutsuki {
	o.energy -= 1
	o.currentVillage = village
	return o
}

func (o *Otsutsuki) KillOtsutsuki() {
	o.energy = 0
	o.isAlive = false
}

// IsMovable check weather otsutsuki can move or not
func (o *Otsutsuki) IsMovable() bool {
	return len(o.currentVillage.neighbours) > 0 && o.isAlive && o.energy > 0
}
