// Package pincheck menyediakan fungsi untuk memeriksa
// apakah PIN 6 digit mengandung kombinasi dari tanggal lahir pengguna.
package pincheck

import (
	"fmt"
	"regexp"
	"strconv"
)

// twoDigit mengubah angka menjadi string 2 digit, misalnya 5 -> "05"
func twoDigit(n int) string {
	return fmt.Sprintf("%02d", n)
}

// yearParts memisahkan tahun menjadi 2 digit awal dan 2 digit akhir
func yearParts(year int) (string, string) {
	y := fmt.Sprintf("%04d", year)
	return y[:2], y[2:]
}

// generate6FromDOB menghasilkan 6 kombinasi 6 digit dari dd, mm, yy
func generate6FromDOB(day, month int, yy string) []string {
	dd := twoDigit(day)
	mm := twoDigit(month)

	return []string{
		dd + mm + yy, // 1) ddmmyy
		dd + yy + mm, // 2) ddyy mm
		mm + dd + yy, // 3) mmddyy
		mm + yy + dd, // 4) mmyydd
		yy + dd + mm, // 5) yyddmm
		yy + mm + dd, // 6) yymmdd
	}
}

// isSixDigitNumeric memvalidasi apakah string berupa 6 digit numerik
func isSixDigitNumeric(s string) bool {
	match, _ := regexp.MatchString(`^\d{6}$`, s)
	return match
}

// ValidatePIN memeriksa apakah PIN 6 digit mengandung kombinasi tanggal lahir.
// Parameter:
// - day: tanggal lahir (1–31)
// - month: bulan lahir (1–12)
// - year: tahun lahir (4 digit)
// - pin: PIN 6 digit
//
// Return:
// - bool: true jika PIN mengandung kombinasi tanggal lahir
// - error: jika input tidak valid
func ValidatePIN(day, month, year int, pin string) (bool, error) {
	if day < 1 || day > 31 {
		return false, fmt.Errorf("tanggal tidak valid")
	}
	if month < 1 || month > 12 {
		return false, fmt.Errorf("bulan tidak valid")
	}
	if year < 1000 {
		return false, fmt.Errorf("tahun tidak valid (harus 4 digit)")
	}
	if !isSixDigitNumeric(pin) {
		return false, fmt.Errorf("PIN harus berupa 6 digit numerik")
	}

	yyHead, yyTail := yearParts(year)
	variants := append(
		generate6FromDOB(day, month, yyTail),
		generate6FromDOB(day, month, yyHead)...,
	)

	for _, v := range variants {
		if v == pin {
			return true, nil
		}
	}
	return false, nil
}

// ParseMonthString mengonversi string bulan "08" atau "9" ke int dengan aman
func ParseMonthString(monthStr string) (int, error) {
	month, err := strconv.Atoi(monthStr)
	if err != nil || month < 1 || month > 12 {
		return 0, fmt.Errorf("bulan tidak valid: %s", monthStr)
	}
	return month, nil
}
