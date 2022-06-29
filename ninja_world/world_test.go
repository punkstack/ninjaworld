package ninja_world

import (
	"github.com/punkstack/ninjaworld/ninja_world/ninja_world_errors"
	"github.com/punkstack/ninjaworld/pkg/utils"
	"testing"
)

func initNinjaWorld() *World {
	return NewWorld()
}

func TestWorld_AddVillage(t *testing.T) {
	ninjaWorld := initNinjaWorld()
	err := ninjaWorld.AddVillage("test")
	if err != nil {
		t.Errorf("got an error in adding a village")
	}

	if len(ninjaWorld.villages) != 1 {
		t.Errorf("Expected 1 village but got %q", len(ninjaWorld.villages))
	}
}

func TestDuplicateVillageAddition(t *testing.T) {
	ninjaWorld := initNinjaWorld()
	err := ninjaWorld.AddVillage("test")
	if err != nil {
		t.Errorf("got an error in adding a village")
	}
	err = ninjaWorld.AddVillage("test")
	if err != ninja_world_errors.VILLAGEALREADYEXISTS {
		t.Errorf("Ninja world already exists but got overrided")
	}
}

func TestWorld_DestroyVillage(t *testing.T) {
	ninjaWorld := initNinjaWorld()
	err := ninjaWorld.AddVillage("test")
	if err != nil {
		t.Errorf("got an error in adding a village")
	}
	err = ninjaWorld.DestroyVillage("test")
	if err != nil {
		t.Errorf("Not able to destroy a village")
	}
	if len(ninjaWorld.destroyedVillages) != 1 {
		t.Errorf("village not added to destroyed village map")
	}
	if len(ninjaWorld.villages) != 0 {
		t.Errorf("village not removed from villages map")
	}
}

func TestWorld_GetVillageByName(t *testing.T) {
	ninjaWorld := initNinjaWorld()

	err := ninjaWorld.AddVillage("test")
	if err != nil {
		t.Errorf("got an error in adding a village")
	}

	village, err := ninjaWorld.GetVillageByName("test")
	if err != nil {
		t.Errorf("not able to get village by name")
	}

	if village.name != "test" {
		t.Errorf("invalid village fetched")
	}
}

func TestWorld_GetRandomVillage(t *testing.T) {
	ninjaWorld := initNinjaWorld()

	err := ninjaWorld.AddVillage("test1")
	if err != nil {
		t.Errorf("got an error in adding a village")
	}

	err = ninjaWorld.AddVillage("test2")
	if err != nil {
		t.Errorf("got an error in adding a village")
	}

	err = ninjaWorld.AddVillage("test3")
	if err != nil {
		t.Errorf("got an error in adding a village")
	}

	village := ninjaWorld.GetRandomVillage()
	if village == nil {
		t.Errorf("random village not found")
	}
}

func TestWorld_GetRemainingVillageString(t *testing.T) {
	ninjaWorld := initNinjaWorld()
	err := ninjaWorld.AddVillage("test")
	if err != nil {
		t.Errorf("got an error in adding a village")
	}
	if ninjaWorld.GetRemainingVillageString()[0] != "test" {
		t.Errorf("remaining village string not matched")
	}
}

func TestWorld_GetRemainingVillageStringWithDestroyedVillage(t *testing.T) {
	ninjaWorld := initNinjaWorld()

	err := ninjaWorld.AddVillage("test1")
	if err != nil {
		t.Errorf("got an error in adding a village")
	}

	err = ninjaWorld.AddVillage("test2")
	if err != nil {
		t.Errorf("got an error in adding a village")
	}

	err = ninjaWorld.DestroyVillage("test2")
	if err != nil {
		t.Errorf("got an error while destroying a village")
	}

	if ninjaWorld.GetRemainingVillageString()[0] != "test1" {
		t.Errorf("remaining village string not matched")
	}
}

func TestWorld_AddOtsutsuki(t *testing.T) {
	ninjaWorld := initNinjaWorld()

	err := ninjaWorld.AddOtsutsuki()
	if err != nil {
		t.Errorf("not able to add otsutsuki")
	}

	err = ninjaWorld.AddOtsutsuki()
	if err != nil {
		t.Errorf("not able to add otsutsuki")
	}

	if len(ninjaWorld.otsutsukies) != 2 {
		t.Errorf("otsutsuki count miss match")
	}
}

func TestWorld_GetOtsutsukies(t *testing.T) {
	ninjaWorld := initNinjaWorld()

	err := ninjaWorld.AddOtsutsuki()
	if err != nil {
		t.Errorf("not able to add otsutsuki")
	}

	err = ninjaWorld.AddOtsutsuki()
	if err != nil {
		t.Errorf("not able to add otsutsuki")
	}

	otsutsukies := ninjaWorld.GetOtsutsukies()
	if len(otsutsukies) != 2 {
		t.Errorf("otsutsuki count miss match")
	}
}

func TestWorld_DeployOtsutsukies(t *testing.T) {
	ninjaWorld := initNinjaWorld()

	err := ninjaWorld.AddOtsutsuki()
	if err != nil {
		t.Errorf("not able to add otsutsuki")
	}

	err = ninjaWorld.AddOtsutsuki()
	if err != nil {
		t.Errorf("not able to add otsutsuki")
	}

	err = ninjaWorld.DeployOtsutsukies()
	if err != ninja_world_errors.NOVILLAGELEFTFOROTSUTSUKI {
		t.Errorf("otsutsuki deployed in unknown village fatal")
	}

	err = ninjaWorld.AddVillage("test1")
	if err != nil {
		t.Errorf("got an error in adding a village")
	}

	err = ninjaWorld.DeployOtsutsukies()
	if err != nil {
		t.Errorf("otsutsuki random deployment failed")
	}

	for _, value := range ninjaWorld.otsutsukies {
		if value.currentVillage.name != "test1" {
			t.Errorf("otsutsuki deployed in unknown village fatal")
		}
	}
}

func TestWorld_ExecuteWar(t *testing.T) {
	ninjaWorld := initNinjaWorld()

	err := ninjaWorld.AddOtsutsuki()
	if err != nil {
		t.Errorf("not able to add otsutsuki")
	}

	err = ninjaWorld.AddOtsutsuki()
	if err != nil {
		t.Errorf("not able to add otsutsuki")
	}

	err = ninjaWorld.AddVillage("test1")
	if err != nil {
		t.Errorf("got an error in adding a village")
	}

	err = ninjaWorld.DeployOtsutsukies()
	if err != nil {
		t.Errorf("otsutsuki deployed in unknown village fatal")
	}

	ninjaWorld.ExecuteWar()

	if len(ninjaWorld.villages) != 0 {
		t.Errorf("war simulation failed")
	}

	for _, o := range ninjaWorld.otsutsukies {
		if o.isAlive {
			t.Errorf("war simulation failed")
		}
	}
}

func TestWorld_MoveOtsutukies(t *testing.T) {
	ninjaWorld := initNinjaWorld()

	err := ninjaWorld.AddOtsutsuki()
	if err != nil {
		t.Errorf("not able to add otsutsuki")
	}

	err = ninjaWorld.AddVillage("test1")
	if err != nil {
		t.Errorf("got an error in adding a village")
	}

	err = ninjaWorld.AddVillage("test2")
	if err != nil {
		t.Errorf("got an error in adding a village")
	}

	test1Village, err := ninjaWorld.GetVillageByName("test1")
	if err != nil {
		t.Errorf("failed to get village")
	}
	test2Village, err := ninjaWorld.GetVillageByName("test2")
	if err != nil {
		t.Errorf("failed to get village")
	}

	err = test1Village.AddNeighbour(utils.GetNorthCardinalDirection(), test2Village)
	if err != nil {
		t.Errorf("failed to add neighbouring village")
	}

	err = ninjaWorld.DeployOtsutsukies()
	if err != nil {
		t.Errorf("otsutsuki deployed in unknown village fatal")
	}

	currentOtsutsukiVillage := ""
	for _, o := range ninjaWorld.otsutsukies {
		currentOtsutsukiVillage = o.currentVillage.name
	}

	ninjaWorld.MoveOtsutukies()

	for _, o := range ninjaWorld.otsutsukies {
		if currentOtsutsukiVillage == o.currentVillage.name {
			t.Errorf("otsutsuki movement error")
		}
	}

}

func TestWorld_HasWarEnded(t *testing.T) {
	ninjaWorld := initNinjaWorld()

	if ninjaWorld.HasWarEnded() == false {
		t.Errorf("otsutsuki check failed")
	}

	err := ninjaWorld.AddOtsutsuki()
	if err != nil {
		t.Errorf("error deploying otsutsuki")
	}

	err = ninjaWorld.AddVillage("test1")
	if err != nil {
		t.Errorf("got an error in adding a village")
	}

	err = ninjaWorld.AddVillage("test2")
	if err != nil {
		t.Errorf("got an error in adding a village")
	}

	err = ninjaWorld.villages["test1"].AddNeighbour(utils.GetNorthCardinalDirection(), ninjaWorld.villages["test2"])
	if err != nil {
		t.Errorf("got an error in adding neighbour village")
	}

	var currentOtsutsuki *Otsutsuki
	for _, o := range ninjaWorld.otsutsukies {
		currentOtsutsuki = o
	}

	currentOtsutsuki.currentVillage = ninjaWorld.GetRandomVillage()
	currentOtsutsuki.isAlive = false

	if ninjaWorld.HasWarEnded() == false {
		t.Errorf("war ended logic check issue")
	}

	currentOtsutsuki.isAlive = true
	if ninjaWorld.HasWarEnded() == true {
		t.Errorf("war ended logic check issue")
	}

	currentOtsutsuki.energy = 0
	if ninjaWorld.HasWarEnded() == false {
		t.Errorf("war ended logic check issue")
	}

}
