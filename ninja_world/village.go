package ninja_world

import (
	"github.com/punkstack/ninjaworld/ninja_world/ninja_world_errors"
	utils "github.com/punkstack/ninjaworld/pkg/utils"
)

type VillageInterface interface {
	VillageDestroyed()
	AddNeighbour(direction *utils.Direction, village *Village) error
	removeAllNeighbourVillages()
	AddOtsutsuki(otsutsuki *Otsutsuki) *Village
	AreNeighboursAvailable() bool
	GetRandomNeighbourVillage() *Village
	RemoveOtsutsuki(otsutsuki *Otsutsuki) *Village
}

type Village struct {
	name        string
	neighbours  map[utils.Direction]*Village
	isDestroyed bool
	otsutsukies []*Otsutsuki
}

var _ VillageInterface = &Village{}

func NewVillage(name string) *Village {
	return &Village{
		name:        name,
		neighbours:  map[utils.Direction]*Village{},
		isDestroyed: false,
		otsutsukies: []*Otsutsuki{},
	}
}

func (v *Village) VillageDestroyed() {
	v.isDestroyed = true
	v.removeAllNeighbourVillages()
}

// AddNeighbour adds neighbouring village with cardinal direction
func (v *Village) AddNeighbour(direction *utils.Direction, village *Village) error {
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
	err := village.AddNeighbour(direction.GetOppositeDirection(), v)
	if err != nil {
		return err
	}
	return nil
}

// removeAllNeighbourVillages linked to village
func (v *Village) removeAllNeighbourVillages() {
	for key, neighbour := range v.neighbours {
		delete(neighbour.neighbours, *key.GetOppositeDirection())
	}
	v.neighbours = map[utils.Direction]*Village{}
}

func (v *Village) AddOtsutsuki(otsutsuki *Otsutsuki) *Village {
	v.otsutsukies = append(v.otsutsukies, otsutsuki)
	return v
}

func (v *Village) AreNeighboursAvailable() bool {
	return len(v.neighbours) > 0
}

func (v *Village) GetRandomNeighbourVillage() *Village {
	if len(v.neighbours) == 0 {
		return nil
	}
	randomIndex := utils.Pick(len(v.neighbours))
	for key := range v.neighbours {
		if randomIndex == 0 {
			if entry, ok := v.neighbours[key]; ok {
				return entry
			}
		}
		randomIndex--
	}
	return nil
}

func (v *Village) RemoveOtsutsuki(otsutsuki *Otsutsuki) *Village {
	for idx, currentOtsutsuki := range v.otsutsukies {
		if otsutsuki.name == currentOtsutsuki.name {
			v.otsutsukies[idx] = v.otsutsukies[len(v.otsutsukies)-1]
			v.otsutsukies = v.otsutsukies[:len(v.otsutsukies)-1]
			break
		}
	}
	return v
}
