# 🗺️ gomaps - Google Maps Scraper CLI

![Go Version](https://img.shields.io/badge/Go-1.21%2B-00ADD8?logo=go)
![License](https://img.shields.io/badge/license-MIT-green)
![Status](https://img.shields.io/badge/status-Active-success)
![Author](https://img.shields.io/badge/developer-Guruh_Rachmat_P-blue)

**gomaps** adalah *Command Line Interface (CLI)* untuk melakukan **scraping data dari Google Maps** berdasarkan *keyword* dan *lokasi*.  
Dikembangkan dengan oleh **Guruh Rachmat Pribadi © 2025**, tool ini cocok untuk pengumpulan data bisnis, tempat umum, pendidikan, atau riset lokasi.

**gomaps** is a *Command Line Interface (CLI) tool* designed to **scrape data from Google Maps** based on a *keyword* and *location*.
Developed by **Guruh Rachmat Pribadi © 2025**, this tool is ideal for collecting data on businesses, public places, educational institutions, or location-based research.
---

## ✨ Fitur Utama

✅ Scrape data tempat dari Google Maps secara otomatis  
✅ Tentukan *keyword* pencarian (`--place`) dan *lokasi target* (`--location`)  
✅ Batas hasil pencarian dengan `--limit` (default: 20)  
✅ Mendukung berbagai format input (`csv`, `json`) *(progress)*  
✅ Dibangun menggunakan **Go Cobra+ Chromedp** — cepat, ringan, dan efisien   

---

## ⚙️ Instalasi

Pastikan sudah menginstal Go versi ≥ 1.21.

### Opsi 1 - Install langsung
```bash
go install github.com/username/gomaps@latest
```

### Opsi 2 - Jalankan dari source
```bash
git clone https://github.com/username/gomaps.git
cd gomaps
go run main.go scraping --place="cafe" --location="bandung"
```

## 🧭 Cara Penggunaan
```bash
Usage:
  gomaps [command]

Available Commands:
  completion  Generate autocompletion script
  help        Help about any command
  scraping    Scrape data dari Google Maps

gomaps scraping --help

Command scraping digunakan untuk melakukan pencarian & scraping data dari Google Maps.
Gunakan flag --place/-p untuk menentukan keyword pencarian.
Gunakan flag --location/-l untuk menentukan lokasi target pencarian.
Gunakan flag --limit untuk menentukan batas hasil pencarian (default 20).
Gunakan flag --import/-i untuk menentukan format input source (default csv).

Contoh penggunaan:
  gomaps scraping --place="sd negeri" --location="depok jawa barat" --limit=100 --import=csv

Flag --place/-p dan --location/-l wajib diisi.
Flag --limit dan --import/-i opsional.

Selamat mencoba ( ͡° ͜ʖ ͡°)
```



