package io

import (
	"fmt"
	"github.com/punkstack/ninjaworld/ninja_world"
	"github.com/punkstack/ninjaworld/pkg/logger"
	"testing"
)

func TestWorldReadFromFile(t *testing.T) {
	logger.InitializeLogger("../../logs/test.json")
	ninjaWorld := ninja_world.NewWorld()
	err := ParseInputFile("../../tests/input.txt", ninjaWorld)
	if err != nil {
		fmt.Println(err)
	}
}
