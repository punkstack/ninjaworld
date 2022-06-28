package ninja_world

import (
	"github.com/punkstack/ninjaworld/ninja_world/ninja_world_errors"
	"github.com/punkstack/ninjaworld/ninja_world/utils"
)

type Village struct {
	name        string
	neighbours  map[utils.Direction]Village
	isDestroyed bool
	otsutsukies []Otsutsuki
}

func NewVillage(name string) *Village {
	return &Village{
		name:        name,
		neighbours:  map[utils.Direction]Village{},
		isDestroyed: false,
		otsutsukies: []Otsutsuki{},
	}
}

func (v *Village) SetVillageDestroyed() {
	v.isDestroyed = true
	v.handleCardinalVillages()
}

func (v *Village) AddNeighbour(direction *utils.Direction, village Village) error {
	if value, exists := v.neighbours[*direction]; exists {
		if value.name == village.name {
			return nil
		}
	}
	for key, value := range v.neighbours {
		if key == *direction || village.name == value.name {
			return ninja_world_errors.VILLAGEALREADYINLINK
		}
	}
	v.neighbours[*direction] = village
	err := village.AddNeighbour(direction.GetOppositeDirection(), *v)
	if err != nil {
		return err
	}
	return nil
}

func (v *Village) handleCardinalVillages() {
	for key, neighbour := range v.neighbours {
		delete(neighbour.neighbours, *key.GetOppositeDirection())
	}
	v.neighbours = nil
}

func (v *Village) AddOtsutsuki(otsutsuki Otsutsuki) *Village {
	v.otsutsukies = append(v.otsutsukies, otsutsuki)
	return v
}

func (v *Village) AreNeighboursAvailable() bool {
	return len(v.neighbours) > 0
}

func (v *Village) GetRandomNeighbourVillage() Village {
	randomIndex := utils.Pick(len(v.neighbours))

	for key := range v.neighbours {
		if randomIndex == 0 {
			if entry, ok := v.neighbours[key]; ok {
				return entry
			}
		}
		randomIndex--
	}

	return Village{}
}

func (v *Village) RemoveOtsutsuki(otsutsuki Otsutsuki) *Village {
	for idx, currentOtsutsuki := range v.otsutsukies {
		if otsutsuki.name == currentOtsutsuki.name {
			v.otsutsukies[idx] = v.otsutsukies[len(v.otsutsukies)-1]
			v.otsutsukies = v.otsutsukies[:len(v.otsutsukies)-1]
			break
		}
	}
	return v
}
