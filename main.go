/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
	"github.com/punkstack/ninjaworld/cmd"
	"github.com/punkstack/ninjaworld/pkg/logger"
)

func main() {
	logger.InitializeLogger("logs/simulation_log.json")
	cmd.Execute()
}
