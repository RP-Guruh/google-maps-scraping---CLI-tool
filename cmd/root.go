/*
Copyright © 2025 GURUH RACHMAT PRIBADI <github.guruh@gmail.com>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gomaps",
	Short: "Scrape data dari Google Maps, berdasarkan keyword dan lokasi. Developed by Guruh Rachmat Pribadi © 2025",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {

}
