/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/Rayato159/cobra-pg/src"
	"github.com/spf13/cobra"
)

var r float64

// findAreaCircleCmd represents the findAreaCircle command
var findAreaCircleCmd = &cobra.Command{
	Use:   "circle",
	Short: "Calculate circle area",
	Long:  `To calculate a circle area followed by the given radius.`,
	Run: func(cmd *cobra.Command, args []string) {
		if r == 0.0 {
			errMsg := src.BuildExceptionMessage(src.RadiusIsRequiredException, "radius is required")
			log.Fatal(errMsg)
		}

		result := src.CalculateAreaCircle(r)
		fmt.Printf("Area: %.6f units\n", result)
	},
}

func init() {
	findAreaCircleCmd.Flags().Float64VarP(&r, "radius", "r", 0.0, "Radius of the circle")

	rootCmd.AddCommand(findAreaCircleCmd)
}
