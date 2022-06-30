package ninja_world

import (
	"fmt"
	"github.com/punkstack/ninjaworld/ninja_world/ninja_world_errors"
	"github.com/punkstack/ninjaworld/pkg/utils"
	"strings"
	"syreclabs.com/go/faker"
)

type WorldInterface interface {
	AddVillage(name string) error
	GetVillageByName(name string) (*Village, error)
	DestroyVillage(name string) error
	GetRandomVillage() *Village
	GetRemainingVillageString() []string
	AddOtsutsuki() error
	GetOtsutsukies() map[string]*Otsutsuki
	DeployOtsutsukies() error
	ExecuteWar()
	HasWarEnded() bool
}

type World struct {
	villages          map[string]*Village
	otsutsukies       map[string]*Otsutsuki
	destroyedVillages map[string]*Village
}

var _ WorldInterface = &World{}

func NewWorld() *World {
	return &World{
		villages:          map[string]*Village{},
		destroyedVillages: map[string]*Village{},
		otsutsukies:       map[string]*Otsutsuki{},
	}
}

// AddVillage Create and add new village to ninja world if village doesn't exist
func (w *World) AddVillage(name string) error {
	if _, exists := w.destroyedVillages[name]; exists {
		return ninja_world_errors.VILLAGEDESTROYED
	}
	if _, exists := w.villages[name]; exists {
		return ninja_world_errors.VILLAGEALREADYEXISTS
	} else {
		w.villages[name] = NewVillage(name)
	}
	return nil
}

// DestroyVillage Destroys the village and updates the required villages state
func (w *World) DestroyVillage(name string) error {
	village := w.villages[name]
	if _, exists := w.destroyedVillages[name]; exists {
		return ninja_world_errors.VILLAGEDESTROYED
	} else {
		village.VillageDestroyed()
		delete(w.villages, name)
		w.destroyedVillages[name] = village
	}
	return nil
}

func (w *World) GetVillageByName(name string) (*Village, error) {
	if _, exists := w.villages[name]; exists {
		return w.villages[name], nil
	} else {
		return nil, ninja_world_errors.VILLAGEDOESNOTEXISTS
	}
}

// GetRandomVillage returns a village of non-destroyed state
func (w *World) GetRandomVillage() *Village {
	randomIndex := utils.Pick(len(w.villages))
	for key := range w.villages {
		if randomIndex == 0 {
			if entry, ok := w.villages[key]; ok {
				return entry
			}
		}
		randomIndex--
	}

	return nil
}

// GetRemainingVillageString returns the remaining villages with the village and neighbouring villages with direction in string
func (w *World) GetRemainingVillageString() []string {
	result := make([]string, 0)
	for key, value := range w.villages {
		currentVillageStatus := key
		for direction, neighbour := range value.neighbours {
			currentVillageStatus += fmt.Sprintf(" %s=%s ", direction.String(), neighbour.name)
		}
		result = append(result, currentVillageStatus)
	}
	return result
}

// AddOtsutsuki adds a new otsutsuki to ninja world
func (w *World) AddOtsutsuki() error {
	otsutsukiName := faker.Name().FirstName()

	// this is to maintain unique otsutsukiName
	if _, exists := w.otsutsukies[otsutsukiName]; exists {
		return w.AddOtsutsuki()
	} else {
		w.otsutsukies[otsutsukiName] = NewOtsutsuki(otsutsukiName)
		return nil
	}
}

func (w *World) GetOtsutsukies() map[string]*Otsutsuki {
	return w.otsutsukies
}

// DeployOtsutsukies deploys all available Otsutsukies to random villages, updates village and otsutsuki states
func (w *World) DeployOtsutsukies() error {
	if len(w.villages) == 0 {
		return ninja_world_errors.NOVILLAGELEFTFOROTSUTSUKI
	}
	for _, ots := range w.GetOtsutsukies() {
		randomVillage := w.GetRandomVillage()
		ots.moveOtsutsuki(randomVillage)
		randomVillage.AddOtsutsuki(ots)
	}
	return nil
}

// ExecuteWar checks otsutsuki presence in villages, updates village and otsutsuki status
func (w *World) ExecuteWar() {
	for _, village := range w.villages {
		if len(village.otsutsukies) > 1 {
			otsutsukies := make([]string, 0)
			village.isDestroyed = true
			for _, ots := range village.otsutsukies {
				otsutsukies = append(otsutsukies, ots.name)
				ots.KillOtsutsuki()
			}
			for _, n := range village.neighbours {
				for key, value := range n.neighbours {
					if value.name == village.name {
						delete(n.neighbours, key)
					}
				}
			}
			delete(w.villages, village.name)
			fmt.Printf("%s is destroyed by otsutsukies %s and %s", village.name, strings.Join(otsutsukies[:len(otsutsukies)-1], `,`), otsutsukies[len(otsutsukies)-1])
			fmt.Println()
		}
	}
}

// MoveOtsutukies Move all otsutsukies to random neighbour
func (w *World) MoveOtsutukies() {
	for _, otsutsuki := range w.otsutsukies {
		if otsutsuki.IsMovable() {
			village, err := w.GetVillageByName(otsutsuki.currentVillage.name)
			if err != nil {
				panic(fmt.Sprintf("Village not found %s", err.Error()))
			}
			randomVillage := village.GetRandomNeighbourVillage()

			// updating otsutsuki status
			otsutsuki.moveOtsutsuki(randomVillage)

			// fetching a random village and updating village otsutsukies
			randomVillage.AddOtsutsuki(otsutsuki)

			// Removing otsutsuki from current village
			village.RemoveOtsutsuki(otsutsuki)
		}
	}
}

// HasWarEnded returns if the war ends
func (w *World) HasWarEnded() bool {
	// check base condition where no otsutsuki is present
	if len(w.otsutsukies) == 0 {
		return true
	}

	// check for any movable otsutsuki
	for _, otsutsuki := range w.otsutsukies {
		if otsutsuki.IsMovable() {
			return false
		}
	}
	return true
}
