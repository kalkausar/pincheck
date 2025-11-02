package main

import (
	"fmt"
	"log"

	"service-check-date-pin/pincheck"
)

func main() {
	// Contoh input
	day := 12
	monthStr := "08" // bisa "8" atau "08"
	year := 1998
	pin := "120898"

	// Parsing bulan string ke int
	month, err := pincheck.ParseMonthString(monthStr)
	if err != nil {
		log.Fatal(err)
	}

	match, err := pincheck.ValidatePIN(day, month, year, pin)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("📅 DOB: %02d-%02d-%04d\n🔐 PIN: %s\n", day, month, year, pin)

	if match {
		fmt.Println("⚠️  PIN kamu mengandung kombinasi tanggal lahir, silahkan ganti dengan PIN yang lain.")
	} else {
		fmt.Println("✅ PIN kamu AMAN dan diizinkan.")
	}
}
