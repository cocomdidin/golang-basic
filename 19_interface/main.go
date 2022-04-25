package main

import "fmt"

type BangunDatar interface {
	HitungLuas() int
}

type Persegi struct {
	Sisi int
}

func (persegi Persegi) HitungLuas() int {
	return persegi.Sisi * persegi.Sisi
}

type PersegiPanjang struct {
	Panjang int
	Lebar   int
}

func (persegiPanjang PersegiPanjang) HitungLuas() int {
	return persegiPanjang.Panjang * persegiPanjang.Lebar
}

func main() {
	persegi := Persegi{Sisi: 5}
	fmt.Println("Luas Persegi:", luas(persegi))
	
	persegiPanjang := PersegiPanjang{Panjang: 5, Lebar: 10}
	fmt.Println("Luas Persegi Panjang:", luas(persegiPanjang))
}

func luas(bangunDatar BangunDatar) int {
	return bangunDatar.HitungLuas()
}