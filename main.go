/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
	"github.com/punkstack/ninjaworld/cmd"
	"github.com/punkstack/ninjaworld/ninja_world/logger"
)

func main() {
	logger.InitializeLogger("simulation_log.json")
	cmd.Execute()
}
