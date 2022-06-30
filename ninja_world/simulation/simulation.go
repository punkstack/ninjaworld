package simulation

import (
	"fmt"
	"github.com/punkstack/ninjaworld/ninja_world"
	"github.com/punkstack/ninjaworld/ninja_world/io"
	"github.com/punkstack/ninjaworld/pkg/logger"
)

type Simulation struct {
}

// NewSimulation this is a runner function which expects the input
// otsutsukiCount: count of aliens
// inputFileName: pass the input map filename with full path
// outputFilename: pass the output filename with full path where we want to store the leftover villages/*
func NewSimulation(otsutsukiCount int32, inputFilename, outputFileName string) error {
	fmt.Println("Simulation started")
	ninjaWorld := ninja_world.NewWorld()

	err := io.ParseInputFile(inputFilename, ninjaWorld)
	if err != nil {
		logger.Sugar.Error(err)
		return err
	}

	for idx := 0; idx < int(otsutsukiCount); idx++ {
		err = ninjaWorld.AddOtsutsuki()
		if err != nil {
			logger.Sugar.Error(err)
			return err
		}
	}

	launchSimulation(ninjaWorld)

	err = io.WriteResultToOutputFile(outputFileName, ninjaWorld)
	if err != nil {
		logger.Sugar.Error(err)
		return err
	}
	fmt.Println("Simulation ended")
	return nil
}

// launchSimulation handles the execution of the ninja world simulation
func launchSimulation(ninjaWorld *ninja_world.World) {
	err := ninjaWorld.DeployOtsutsukies()
	if err != nil {
		logger.Sugar.Error("Error ", err.Error())
		return
	}
	ninjaWorld.ExecuteWar()
	for !ninjaWorld.HasWarEnded() {
		ninjaWorld.MoveOtsutukies()
		ninjaWorld.ExecuteWar()
	}
}
