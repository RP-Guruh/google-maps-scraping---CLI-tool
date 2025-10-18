package scraper

import (
	"context"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/chromedp/chromedp"
	"github.com/olekukonko/tablewriter"
)

var URL = "https://www.google.com/maps"
var names, ratings, links []string
var data [][]string

const DEFAULTADATASCROLL = 14

func ScrapperRun(place, location string, limit int) {

	// cek koneksi internet
	ok := CheckInternetConnection()
	if !ok {
		log.Fatal("Program dihentikan karena tidak ada koneksi internet.")
	}

	fmt.Println("\nStarting scrapper...")
	fmt.Println("====================")

	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", false),
	)
	ctx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()
	ctx, cancel = chromedp.NewContext(ctx)
	defer cancel()

	if err := chromedp.Run(ctx,
		chromedp.Navigate(URL),
		chromedp.WaitReady(`#searchboxinput`, chromedp.ByID),
		chromedp.SetValue(`#searchboxinput`, place+" di "+location, chromedp.ByID),
		chromedp.Click(`#searchbox-searchbutton`, chromedp.ByID),
		chromedp.Sleep(5*time.Second),
	); err != nil {
		log.Fatal(err)
	}

	jumScroll := predictJumlahScroll(limit) + 1

	fmt.Println("Prosess scrolling akan dilakukan sebanyak", jumScroll, "kali, setiap scroll menunggu 2 detik... Mohon ditunggu...")
	panelSelector := `[role="feed"]`

	for i := 0; i < jumScroll; i++ {
		if err := chromedp.Run(ctx,
			chromedp.Evaluate(`
				var panel = document.querySelector('`+panelSelector+`');
				if (panel) {
					panel.scrollBy(0, 1000); // Gulir 1000 piksel
				} else {
					console.error('Panel hasil pencarian tidak ditemukan `+panelSelector+`.');
				}
			`, nil),
			chromedp.Sleep(2*time.Second),
		); err != nil {
			log.Fatal(err)
		}
		log.Println("Scroll ke-", i+1, "selesai.")
	}

	log.Println("Selesai scrolling.")
	fmt.Println("\nData akan otomatis terekspor menjadi csv (progress)")

	var names, ratings, links, address, phone []string

	namesScrap := fmt.Sprintf(`Array.from(document.querySelectorAll("a.hfpxzc")).slice(0, %d).map(e => e.getAttribute("aria-label"))`, limit)
	ratingsScrap := fmt.Sprintf(`Array.from(document.querySelectorAll(".MW4etd")).slice(0, %d).map(e => e.innerText)`, limit)
	linksScrap := fmt.Sprintf(`Array.from(document.querySelectorAll("a.hfpxzc")).slice(0, %d).map(e => e.href)`, limit)

	err := chromedp.Run(ctx,
		chromedp.WaitVisible(`a.hfpxzc`),
		chromedp.Evaluate(namesScrap, &names),
		chromedp.Evaluate(ratingsScrap, &ratings),
		chromedp.Evaluate(linksScrap, &links),
	)

	for i := range names {
		err := chromedp.Run(ctx, chromedp.Navigate(links[i]))
		if err != nil {
			address = append(address, "N/A")
			phone = append(phone, "N/A")
			continue
		}

		addressScrap := getTextIfExists(ctx, `button[data-item-id="address"] div.Io6YTe`, 5*time.Second)
		phoneScrap := getTextIfExists(ctx, `button[data-item-id^="phone:tel:"] div.Io6YTe`, 5*time.Second)

		address = append(address, addressScrap)
		phone = append(phone, phoneScrap)

	}

	if err != nil {
		log.Fatal(err)
	}

	if len(names) < limit {
		fmt.Println("Data yang tersedia pada google maps hanya", len(names), "dari", limit, "yang diminta")
	}

	table := tablewriter.NewTable(os.Stdout)
	table.Header("#", "Name", "Rate", "Link", "Address", "Phone", "Web")

	for i := range names {
		data = append(data, []string{fmt.Sprintf("%d", i+1), names[i], safeIndex(ratings, i), beautifyLink(safeIndex(links, i)), potongString(safeIndex(address, i), 20), safeIndex(phone, i), "N/A"})
	}
	table.Bulk(data)
	table.Render()
}

func beautifyLink(link string) string {
	relink := strings.Split(link, "/data=")[0]
	return relink
}

func safeIndex(arr []string, i int) string {
	if i < len(arr) {
		return arr[i]
	}
	return "N/A"
}

func extractEmails(text string) []string {
	re := regexp.MustCompile(`[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}`)
	return re.FindAllString(text, -1)
}

func predictJumlahScroll(jumlahData int) int {
	var scrolls float64
	scrolls = math.Ceil(float64(jumlahData) / DEFAULTADATASCROLL)
	return int(scrolls)
}

func potongString(address string, maxLen int) string {
	if len(address) > maxLen {
		return address[:maxLen] + "..."
	}
	return address
}

func getTextIfExists(ctx context.Context, selector string, timeout time.Duration) string {
	var text string
	ctxWait, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	err := chromedp.Run(ctxWait,
		chromedp.WaitVisible(selector, chromedp.ByQuery),
		chromedp.Text(selector, &text, chromedp.ByQuery),
	)

	if err != nil {
		return "N/A"
	}

	return text
}
