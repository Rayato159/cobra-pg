package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "external-service",
	Short: "Run external service",
	Long:  `To run the external service.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("radius", "r", false, "Radius of the circle")
	rootCmd.Flags().BoolP("dimension", "d", false, "Dimension of the square")
}
