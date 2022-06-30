package simulation

import (
	"github.com/punkstack/ninjaworld/pkg/logger"
	"testing"
)

func TestNewSimulation(t *testing.T) {
	logger.InitializeLogger("../../logs/test.json")
	err := NewSimulation(1, "../../tests/input.txt", "../../tests/result.txt")
	if err != nil {
		t.Errorf("simulation failed")
	}
}
