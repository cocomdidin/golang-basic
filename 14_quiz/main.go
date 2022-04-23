package main

import "fmt"

func main() {
	scores := [8]int{100, 80, 92, 70, 93, 60, 75, 85}

	// Soal 1 : hitung rata-rata nilai
	fmt.Println("Rata-rata nilai :", average(scores[:]))

	// Soal 2 : cari nilai >= 90
	fmt.Println("Nilai >= 90 :", goodScores(scores[:]))
}

func average(scores []int) float64 {
	var total int
	for _, v := range scores {
		total += v
	}
	return float64(total) / float64(len(scores))
}

func goodScores(scores []int) []int {
	var goodScores []int
	for _, v := range scores {
		if v >= 90 {
			goodScores = append(goodScores, v)
		}
	}
	return goodScores
}