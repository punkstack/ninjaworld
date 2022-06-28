package simulation

import (
	"github.com/punkstack/ninjaworld/ninja_world/logger"
	"testing"
)

func TestNewSimulation(t *testing.T) {
	logger.InitializeLogger("../../simulation.json")
	NewSimulation(1, "../../tests/input.txt", "../../tests/result.txt")
}
