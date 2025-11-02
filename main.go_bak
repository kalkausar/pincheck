package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func twoDigit(n int) string {
	return fmt.Sprintf("%02d", n)
}

func yearParts(year int) (string, string) {
	y := fmt.Sprintf("%04d", year)
	return y[:2], y[2:]
}

func generate6FromDOB(day, month int, yy string) []string {
	dd := twoDigit(day)
	mm := twoDigit(month)

	return []string{
		dd + mm + yy,
		dd + yy + mm,
		mm + dd + yy,
		mm + yy + dd,
		yy + dd + mm,
		yy + mm + dd,
	}
}

func isSixDigitNumeric(s string) bool {
	match, _ := regexp.MatchString(`^\d{6}$`, s)
	return match
}

func main() {
	day := flag.Int("d", -1, "Tanggal lahir (1-31)")
	monthStr := flag.String("m", "", "Bulan lahir (01-12, bisa dengan awalan nol)")
	year := flag.Int("y", -1, "Tahun lahir (4-digit, misal 1998)")
	pin := flag.String("pin", "", "PIN 6 digit yang akan dicek")
	flag.Parse()

	// Konversi bulan string ke int
	month, err := strconv.Atoi(*monthStr)
	if err != nil || month < 1 || month > 12 {
		fmt.Fprintln(os.Stderr, "❌ Bulan tidak valid. Gunakan 01-12.")
		os.Exit(1)
	}

	if *day < 1 || *day > 31 || *year < 1000 {
		fmt.Fprintln(os.Stderr, "❌ Argumen tanggal tidak valid. Contoh: -d 12 -m 05 -y 1998")
		os.Exit(1)
	}

	if !isSixDigitNumeric(*pin) {
		fmt.Fprintln(os.Stderr, "❌ PIN harus berupa 6 digit numerik. Contoh: -pin 120598")
		os.Exit(1)
	}

	yyHead, yyTail := yearParts(*year)
	var variants []string
	variants = append(variants, generate6FromDOB(*day, month, yyTail)...)
	variants = append(variants, generate6FromDOB(*day, month, yyHead)...)

	isMatch := false
	for _, v := range variants {
		if v == *pin {
			isMatch = true
			break
		}
	}

	fmt.Println("────────────────────────────")
	fmt.Printf("📅 Tanggal Lahir: %02d-%02d-%04d\n", *day, month, *year)
	fmt.Printf("🔐 PIN Diperiksa : %s\n", *pin)
	fmt.Println("────────────────────────────")

	if isMatch {
		fmt.Println("⚠️  PIN kamu mengandung kombinasi tanggal lahir, silahkan ganti dengan PIN yang lain.")
	} else {
		fmt.Println("✅ PIN kamu AMAN dan diizinkan.")
	}
}
