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

var d float64

// findAreaSquareCmd represents the findAreaSquare command
var findAreaSquareCmd = &cobra.Command{
	Use:   "square",
	Short: "Calculate square area",
	Long:  `To calculate a square area followed by the given dimension.`,
	Run: func(cmd *cobra.Command, args []string) {
		if d == 0.0 {
			errMsg := src.BuildExceptionMessage(src.SquareDimensionIsRequiredException, "dimension is required")
			log.Fatal(errMsg)
		}

		result := src.CalculateAreaSquare(d)
		fmt.Printf("Area: %.6f units\n", result)
	},
}

func init() {
	findAreaSquareCmd.Flags().Float64VarP(&d, "dimension", "d", 0.0, "Dimension of the square")

	rootCmd.AddCommand(findAreaSquareCmd)
}
