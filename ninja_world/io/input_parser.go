package io

import (
	"bufio"
	"github.com/punkstack/ninjaworld/ninja_world"
	"github.com/punkstack/ninjaworld/ninja_world/logger"
	"github.com/punkstack/ninjaworld/ninja_world/ninja_world_errors"
	"github.com/punkstack/ninjaworld/ninja_world/utils"
	"log"
	"os"
	"strings"
)

type InputParser struct {
}

// ParseInputFile Expects ninja world text file and sets up the ninja world map otsutsuki
func ParseInputFile(filename string, ninjaWorld *ninja_world.World) error {
	logger.Sugar.Info("Parsing input")
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer func(f *os.File) {
		err = f.Close()
		if err != nil {
			panic(err)
		}
	}(f)
	// Init scanner to scan lines
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	// Scanning input file
	for scanner.Scan() {
		currentRow := scanner.Text()
		chunks := strings.Fields(currentRow)
		villageName := chunks[0]
		err = ninjaWorld.AddVillage(villageName)
		if err != nil {
			return err
		}
		village, err := ninjaWorld.GetVillage(villageName)
		if err != nil {
			return err
		}
		for idx := 1; idx < len(chunks); idx++ {
			currentStringChunk := strings.Split(chunks[idx], "=")
			direction, currentVillageName := currentStringChunk[0], currentStringChunk[1]
			currentVillage, err := ninjaWorld.GetVillage(currentVillageName)
			if err != nil {
				if err.Error() == ninja_world_errors.VILLAGEDOESNOTEXISTS.Error() {
					err = ninjaWorld.AddVillage(currentVillageName)
					if err != nil {
						return err
					}
					currentVillage, err = ninjaWorld.GetVillage(currentVillageName)
					if err != nil {
						return err
					}
				} else {
					return err
				}
			}
			err = village.AddNeighbour(utils.GetDirectionByString(direction), currentVillage)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func WriteResultToOutputFile(filename string, ninjaWorld *ninja_world.World) error {
	results := ninjaWorld.GetRemainingVillageString()
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		log.Fatalf("failed creating file: %s", err)
		return err
	}
	defer file.Close()
	datawriter := bufio.NewWriter(file)

	for _, data := range results {
		_, _ = datawriter.WriteString(data + "\n")
	}

	datawriter.Flush()
	return nil
}
