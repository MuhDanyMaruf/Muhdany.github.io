package main

import (
	"fmt"
)

func main() {
	// Operasi Penjumlahan
	hasilPenjumlahan := penjumlahan(5, 3)
	fmt.Println("Hasil Penjumlahan:", hasilPenjumlahan)

	// Operasi Pengurangan
	hasilPengurangan := pengurangan(8, 3)
	fmt.Println("Hasil Pengurangan:", hasilPengurangan)

	// Operasi Perkalian
	hasilPerkalian := perkalian(4, 6)
	fmt.Println("Hasil Perkalian:", hasilPerkalian)

	// Operasi Pembagian
	hasilPembagian := pembagian(10, 2)
	fmt.Println("Hasil Pembagian:", hasilPembagian)
}

func penjumlahan(a, b int) int {
	return a + b
}

func pengurangan(a, b int) int {
	return a - b
}

func perkalian(a, b int) int {
	return a * b
}

func pembagian(a, b int) float64 {
	return float64(a) / float64(b)
}
