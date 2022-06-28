package simulation

import (
	"github.com/punkstack/ninjaworld/ninja_world"
	"github.com/punkstack/ninjaworld/ninja_world/io"
	"github.com/punkstack/ninjaworld/ninja_world/logger"
	"os"
)

type Simulation struct {
}

func NewSimulation(otsutsukiCount int32, inputFilename, outputFileName string) {
	ninjaWorld := ninja_world.NewWorld()

	err := io.ParseInputFile(inputFilename, ninjaWorld)
	if err != nil {
		logger.Sugar.Error(err)
		os.Exit(1)
	}

	for idx := 0; idx < int(otsutsukiCount); idx++ {
		err = ninjaWorld.AddOtsutsuki()
		if err != nil {
			logger.Sugar.Error(err)
			panic(err)
		}
	}

	launchSimulation(ninjaWorld)
	err = io.WriteResultToOutputFile(outputFileName, ninjaWorld)
	if err != nil {
		logger.Sugar.Error(err)
		os.Exit(1)
	}
}

func launchSimulation(ninjaWorld *ninja_world.World) {
	ninjaWorld.DeployOtsutsukies()
	ninjaWorld.ExecuteWar()
	for ninjaWorld.IsAnyOtsutsukiAlive() {
		ninjaWorld.MoveOtsutukies()
		ninjaWorld.ExecuteWar()
	}
}
