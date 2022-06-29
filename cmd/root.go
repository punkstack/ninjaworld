package cmd

import (
	"fmt"
	"github.com/punkstack/ninjaworld/ninja_world/simulation"
	"github.com/punkstack/ninjaworld/pkg/logger"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ninjaworld",
	Short: "A simulation of otsutsuki invading ninja world",
	Long:  `Power hungry otsutsuki are back to capture ninja world`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		inputFilename, err := cmd.Flags().GetString("input-filename")
		if err != nil {
			fmt.Println(err)
			return
		}
		outputFilename, err := cmd.Flags().GetString("output-filename")
		if err != nil {
			fmt.Println(err)
			return
		}
		otsutsukiCount, err := cmd.Flags().GetInt32("otsutsuki")
		if err != nil {
			fmt.Println(err)
			return
		}
		err = simulation.NewSimulation(otsutsukiCount, inputFilename, outputFilename)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

// setting input flags and validating cmd flags
func init() {
	rootCmd.Flags().StringP("input-filename", "i", "FILE NAME", "map input file name with file path")
	rootCmd.Flags().StringP("output-filename", "o", "FILE NAME", "map input file name with file path")
	rootCmd.Flags().Int32P("otsutsuki", "n", 2, "count of otsutsuki invading ninja world")
	if err := rootCmd.MarkFlagRequired("input-filename"); err != nil {
		logger.Sugar.Error(err)
	}
	if err := rootCmd.MarkFlagRequired("output-filename"); err != nil {
		logger.Sugar.Error(err)
	}
	if err := rootCmd.MarkFlagRequired("otsutsuki"); err != nil {
		logger.Sugar.Error(err)
	}
}
