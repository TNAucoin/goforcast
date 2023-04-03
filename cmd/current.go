package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/tnaucoin/goforcast/pkg/config"
	"github.com/tnaucoin/goforcast/pkg/weather"
	"log"
)

var currentCmd = &cobra.Command{
	Use:   "current [zip code]",
	Short: "Get a forecast for your location.",
	Long:  "A simple cli for getting the current weather forcast for your location.",
	Run: func(cmd *cobra.Command, args []string) {
		apiKey := config.ReadEnvValue("API_KEY")
		location, err := weather.GetLocationFromZip(args[0], apiKey)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Location: %v", *location)
	},
}

func init() {
	rootCmd.AddCommand(currentCmd)
}
