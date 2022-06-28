package ninja_world

import (
	"fmt"
	"github.com/punkstack/ninjaworld/ninja_world/logger"
	"github.com/punkstack/ninjaworld/ninja_world/ninja_world_errors"
	"github.com/punkstack/ninjaworld/ninja_world/utils"
	"strings"
	"syreclabs.com/go/faker"
)

type WorldInterface interface {
	AddVillage(name string) error
	DestroyVillage(name string) error
	GetVillage(name string) (Village, error)
	AddOtsutsuki() error
	GetOtsutsukies() map[string]Otsutsuki
	GetRandomVillage() Village
	DeployOtsutsukies()
	ExecuteWar()
	GetRemainingVillageString() []string
	IsAnyOtsutsukiAlive() bool
	MoveOtsutukies()
}

type World struct {
	villages          map[string]Village
	otsutsukies       map[string]Otsutsuki
	destroyedVillages map[string]Village
}

var _ WorldInterface = &World{}

func NewWorld() *World {
	return &World{
		villages:          map[string]Village{},
		destroyedVillages: map[string]Village{},
		otsutsukies:       map[string]Otsutsuki{},
	}
}

func (w *World) AddVillage(name string) error {
	if _, exists := w.destroyedVillages[name]; exists {
		return ninja_world_errors.VILLAGEDESTROYED
	}
	if _, exists := w.villages[name]; exists {
		return nil
	} else {
		w.villages[name] = *NewVillage(name)
	}
	return nil
}

func (w *World) DestroyVillage(name string) error {
	village, _ := w.villages[name]
	if _, exists := w.destroyedVillages[name]; exists {
		return ninja_world_errors.VILLAGEDESTROYED
	} else {
		village.SetVillageDestroyed()
		delete(w.villages, name)
		w.destroyedVillages[name] = village
	}
	return nil
}

func (w *World) GetVillage(name string) (Village, error) {
	if _, exists := w.villages[name]; exists {
		return w.villages[name], nil
	} else {
		return Village{}, ninja_world_errors.VILLAGEDOESNOTEXISTS
	}
}

func (w *World) AddOtsutsuki() error {
	otsutsukiName := faker.Name().FirstName()
	logger.Sugar.Info("creating otsutsuki ", otsutsukiName)
	if _, exists := w.otsutsukies[otsutsukiName]; exists {
		return w.AddOtsutsuki()
	} else {
		w.otsutsukies[otsutsukiName] = *NewOtsutsuki(otsutsukiName)
		return nil
	}
}

func (w *World) GetOtsutsukies() map[string]Otsutsuki {
	return w.otsutsukies
}

func (w *World) GetRandomVillage() Village {
	randomIndex := utils.Pick(len(w.villages))
	for key := range w.villages {
		if randomIndex == 0 {
			if entry, ok := w.villages[key]; ok {
				return entry
			}
		}
		randomIndex--
	}

	return Village{}
}

func (w *World) DeployOtsutsukies() {
	for _, ots := range w.GetOtsutsukies() {
		randomVillage := w.GetRandomVillage()
		if entry, ok := w.otsutsukies[ots.name]; ok {
			entry.updateOtsutsuki(randomVillage)
			w.otsutsukies[ots.name] = entry
		}
		if entry, ok := w.villages[randomVillage.name]; ok {
			entry.AddOtsutsuki(ots)
			w.villages[randomVillage.name] = entry
		}
	}
}

func (w *World) ExecuteWar() {
	for _, village := range w.villages {
		if len(village.otsutsukies) > 1 {
			otsutsukies := []string{}
			if entry, ok := w.villages[village.name]; ok {
				entry.isDestroyed = true
				for _, ots := range village.otsutsukies {
					otsutsukies = append(otsutsukies, ots.name)
					if e, ok := w.otsutsukies[ots.name]; ok {
						e.energy = 0
						e.isAlive = false
						w.otsutsukies[ots.name] = e
					}
				}
				for _, n := range village.neighbours {
					for key, value := range n.neighbours {
						if value.name == entry.name {
							delete(n.neighbours, key)
						}
					}
				}
				delete(w.villages, village.name)
				w.destroyedVillages[village.name] = entry
			}
			fmt.Printf("%s is destroyed by otsutsukies %s and %s", village.name, strings.Join(otsutsukies[:len(otsutsukies)-1], `,`), otsutsukies[len(otsutsukies)-1])
			fmt.Println()
		}
	}
}

func (w *World) GetRemainingVillageString() []string {
	result := []string{}
	for key, value := range w.villages {
		currentVillageStatus := key
		if !value.isDestroyed {
			for direction, neighbour := range value.neighbours {
				currentVillageStatus += fmt.Sprintf(" %s=%s ", direction.String(), neighbour.name)
			}
			result = append(result, currentVillageStatus)
		}
	}
	return result
}

func (w *World) IsAnyOtsutsukiAlive() bool {
	for _, otsutsuki := range w.otsutsukies {
		if otsutsuki.isAlive {
			return true
		}
	}
	return false
}

func (w *World) MoveOtsutukies() {
	for _, otsutsuki := range w.otsutsukies {
		if otsutsuki.isAlive {
			if len(otsutsuki.currentVillage.neighbours) > 0 {
				village, err := w.GetVillage(otsutsuki.currentVillage.name)
				if err != nil {
					logger.Sugar.Error("Village not found ", err.Error())
				}
				randomVillage := village.GetRandomNeighbourVillage()
				if entry, ok := w.otsutsukies[otsutsuki.name]; ok {
					entry.updateOtsutsuki(randomVillage)
					w.otsutsukies[otsutsuki.name] = entry
				}
				if entry, ok := w.villages[randomVillage.name]; ok {
					entry.AddOtsutsuki(otsutsuki)
					w.villages[randomVillage.name] = entry
				}
				if entry, ok := w.villages[village.name]; ok {
					entry.RemoveOtsutsuki(otsutsuki)
					w.villages[village.name] = entry
				}
			}
		}
	}
}
