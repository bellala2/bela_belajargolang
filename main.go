package main

import (
	"fmt"
	"math"
)

// Struct 
type Lingkaran struct {
	JariJari float64
}

// luas lingkaran
func (l Lingkaran) HitungLuas() float64 {
	return math.Pi * math.Pow(l.JariJari, 2)
}

func main() {
	// 3. VARIABEL
	var nama string = "meilan"
	var umur int = 18

	kota := "Malang"
	isStudent := true

	
	fmt.Println("Data Diri")
	fmt.Println("Nama:", nama)
	fmt.Println("Umur:", umur)
	fmt.Println("Kota:", kota)
	fmt.Println("Status Mahasiswa:", isStudent)

	fmt.Println("\nLUAS LINGKARAN")
	lingkaranSaya := Lingkaran{JariJari: 10}
	luas := lingkaranSaya.HitungLuas()

	fmt.Println("Jari-jari:", lingkaranSaya.JariJari)
	fmt.Println("Luas Lingkaran:", luas)
}