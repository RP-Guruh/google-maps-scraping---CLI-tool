package scraper

import (
	"fmt"
	"net"
	"time"
)

func CheckInternetConnection() bool {
	timeout := 2 * time.Second
	_, err := net.DialTimeout("tcp", "8.8.8.8:53", timeout) // ping ke google
	if err != nil {
		fmt.Println("Tidak terdapat koneksi internet pada perangkat anda")
		return false
	}
	return true
}
