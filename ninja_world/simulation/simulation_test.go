package simulation

import (
	"github.com/punkstack/ninjaworld/pkg/logger"
	"testing"
)

func TestNewSimulation(t *testing.T) {
	logger.InitializeLogger("../../logs/test.json")
	NewSimulation(4, "../../tests/input.txt", "../../tests/result.txt")
}
