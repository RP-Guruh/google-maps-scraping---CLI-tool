/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"gomaps/models"
	"gomaps/scraper"
	"os"

	"github.com/spf13/cobra"
)

var place string
var location string
var limit int
var importFile string
var Input models.ScrapeInput

// scrapingCmd represents the scraping command
var scrapingCmd = &cobra.Command{
	Use:   "scraping",
	Short: "Scrape data dari Google Maps, berdasarkan keyword dan lokasi",
	Long: `Command scraping digunakan untuk melakukan pencarian & scraping data dari Google Maps.
	Gunakan flag --place/-p untuk menentukan keyword pencarian
	GUnakan flag --location/-l untuk menentukan lokasi target pencarian
	Gunakan flag --limit untuk menentukan batas hasil pencarian (default 20) (progress)
	Gunakan flag --import/-i untuk menentukan format input source (default csv) (progress)

Contoh penggunaan:
  gomaps scraping --place="sd negeri" --location="depok jawa barat" --limit=100 --import=csv

flag --place/-p dan --location/-l wajib diisi, flag --limit dan --import/-i opsional

Selamat mencoba ( ͡° ͜ʖ ͡°)
`,

	RunE: func(cmd *cobra.Command, args []string) error {

		limit, _ := cmd.Flags().GetInt("limit")
		// validasi limit jika kurang dari sama dengan 0 maka set ke 1
		if limit <= 0 {
			limit = 1
		}

		importFile, _ := cmd.Flags().GetString("import")
		// validasi importFile jika bukan csv atau json maka set ke csv
		if importFile != "csv" && importFile != "json" {
			importFile = "csv"
		}

		placeFlags, _ := cmd.Flags().GetString("place")
		// validasi place jika kosong maka keluarin error
		if placeFlags == "" {
			fmt.Println("Error: --place flag is required")
			os.Exit(1)
		}

		locationFlags, _ := cmd.Flags().GetString("location")
		// validasi location jika kosong maka keluarin error
		if locationFlags == "" {
			fmt.Println("Error: --location flag is required")
			os.Exit(1)
		}

		Input = models.ScrapeInput{
			Place:      placeFlags,
			Location:   locationFlags,
			Limit:      limit,
			ImportFile: importFile,
		}

		// kirim hasil input ke parser
		scraper.ParserInit(Input)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(scrapingCmd)
	scrapingCmd.PersistentFlags().StringVarP(&place, "place", "p", "", "Keyword or category of place to search (e.g., 'cafe', 'sd negeri')")
	scrapingCmd.PersistentFlags().StringVarP(&location, "location", "l", "", "Target location or region (e.g., 'depok jawa barat')")
	scrapingCmd.PersistentFlags().IntVarP(&limit, "limit", "", 20, "Limit the number of results to fetch")
	scrapingCmd.PersistentFlags().StringVarP(&importFile, "import", "i", "csv", "Input source format (e.g., csv, json)")
}
