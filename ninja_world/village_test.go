package ninja_world

import (
	"github.com/punkstack/ninjaworld/pkg/utils"
	"testing"
)

func initVillage(name string) *Village {
	return NewVillage(name)
}

func TestVillage_SetVillageDestroyed(t *testing.T) {

	village := initVillage("test1")
	village2 := initVillage("test2")

	err := village.AddNeighbour(utils.GetNorthCardinalDirection(), village2)
	if err != nil {
		t.Errorf("neighbour addition error")
	}

	village.VillageDestroyed()

	if village.isDestroyed == false {
		t.Errorf("destroying village error")
	}

	if len(village.neighbours) != 0 {
		t.Errorf("destroying village error")
	}
}

func TestVillage_AddNeighbour(t *testing.T) {
	village := initVillage("test1")
	village2 := initVillage("test2")

	err := village.AddNeighbour(utils.GetNorthCardinalDirection(), village2)
	if err != nil {
		t.Errorf("neighbour addition error")
	}

	if village.neighbours[*utils.GetNorthCardinalDirection()].name != village2.name {
		t.Errorf("neighbour addition error")
	}
}

func TestVillage_removeAllNeighbourVillages(t *testing.T) {
	village := initVillage("test1")
	village2 := initVillage("test2")

	err := village.AddNeighbour(utils.GetNorthCardinalDirection(), village2)
	if err != nil {
		t.Errorf("neighbour addition error")
	}

	village.removeAllNeighbourVillages()

	if len(village.neighbours) != 0 {
		t.Errorf("neighbour deletion error")
	}

	if len(village2.neighbours) != 0 {
		t.Errorf("neighbour deletion error")
	}
}

func TestVillage_AreNeighboursAvailable(t *testing.T) {
	village := initVillage("test1")
	village2 := initVillage("test2")

	err := village.AddNeighbour(utils.GetNorthCardinalDirection(), village2)
	if err != nil {
		t.Errorf("neighbour addition error")
	}

	if village.AreNeighboursAvailable() == false {
		t.Errorf("neighbour validation error")
	}

	village.removeAllNeighbourVillages()

	if village.AreNeighboursAvailable() == true {
		t.Errorf("neighbour validation error")
	}
}

func TestVillage_GetRandomNeighbourVillage(t *testing.T) {
	village := initVillage("test1")
	village2 := initVillage("test2")

	err := village.AddNeighbour(utils.GetNorthCardinalDirection(), village2)
	if err != nil {
		t.Errorf("neighbour addition error")
	}

	if village.GetRandomNeighbourVillage() == nil {
		t.Errorf("random neighbour error")
	}

	village.removeAllNeighbourVillages()

	if village.GetRandomNeighbourVillage() != nil {
		t.Errorf("random neighbour error")
	}
}

func TestVillage_AddOtsutsuki(t *testing.T) {
	o := NewOtsutsuki("test1")
	village := initVillage("test1")

	village.AddOtsutsuki(o)

	if len(village.otsutsukies) == 0 {
		t.Errorf("addition of otsutsuki error")
	}
}

func TestVillage_RemoveOtsutsuki(t *testing.T) {
	o := NewOtsutsuki("test1")
	village := initVillage("test1")

	village.AddOtsutsuki(o)

	if len(village.otsutsukies) == 0 {
		t.Errorf("remove of otsutsuki error")
	}

	village.RemoveOtsutsuki(o)

	if len(village.otsutsukies) > 0 {
		t.Errorf("remove of otsutsuki error")
	}
}
