package soal

import "fmt"

func avgSlice(slice *[]float32) float32 {
	var total float32
	for _, value := range *slice {
		total += value
	}
	return total / float32(len(*slice))
}

func SoalSatu() {
	fmt.Println("-----Soal 1-----")

	skorkoala := []float32{88, 91, 110}
	skorlumbaLumba := []float32{96, 108, 89}

	skorRataKoala := avgSlice(&skorkoala)
	skorRataLumbaLumba := avgSlice(&skorlumbaLumba)

	// print average score each team (2 decimal)
	fmt.Printf("Rata-rata skor koala: %.2f\n", skorRataKoala)
	fmt.Printf("Rata-rata skor lumba-lumba: %.2f\n", skorRataLumbaLumba)

	fmt.Print("Hasilnya : ")
	// compare average score each team
	if skorRataKoala == skorRataLumbaLumba {
		fmt.Println("Seri")
	} else if skorRataKoala > skorRataLumbaLumba {
		fmt.Println("Koala menang")
	} else {
		fmt.Println("Lumba-lumba menang")
	}

	fmt.Println("-----Bonus 1: -----")

	skorkoala = []float32{109, 95, 123}
	skorlumbaLumba = []float32{97, 112, 101}

	skorRataKoala = avgSlice(&skorkoala)
	skorRataLumbaLumba = avgSlice(&skorlumbaLumba)

	// print average score each team (2 decimal) minimal score 100
	fmt.Printf("Rata-rata skor koala: %.2f\n", skorRataKoala)
	fmt.Printf("Rata-rata skor lumba-lumba: %.2f\n", skorRataLumbaLumba)

	fmt.Print("Hasilnya : ")
	// compare average score each team with minimal score 100
	if skorRataKoala == skorRataLumbaLumba {
		fmt.Println("Seri")
	} else if (skorRataKoala > skorRataLumbaLumba) && (skorRataKoala >= 100) {
		fmt.Println("Koala menang")
	} else if (skorRataKoala < skorRataLumbaLumba) && (skorRataLumbaLumba >= 100) {
		fmt.Println("Lumba-lumba menang")
	} else {
		fmt.Println("Tidak ada pemenang")
	}

	fmt.Println("-----Bonus 2: -----")

	skorkoala = []float32{109, 95, 106}
	skorlumbaLumba = []float32{97, 112, 101}

	skorRataKoala = avgSlice(&skorkoala)
	skorRataLumbaLumba = avgSlice(&skorlumbaLumba)

	// print average score each team (2 decimal) minimal score 100
	fmt.Printf("Rata-rata skor koala: %.2f\n", skorRataKoala)
	fmt.Printf("Rata-rata skor lumba-lumba: %.2f\n", skorRataLumbaLumba)

	fmt.Print("Hasilnya : ")
	// compare average score each team with minimal score 100
	if (skorRataKoala == skorRataLumbaLumba) && (skorRataKoala >= 100) && (skorRataLumbaLumba >= 100) {
		fmt.Println("Seri")
	} else if (skorRataKoala > skorRataLumbaLumba) && (skorRataKoala >= 100) {
		fmt.Println("Koala menang")
	} else if (skorRataKoala < skorRataLumbaLumba) && (skorRataLumbaLumba >= 100) {
		fmt.Println("Lumba-lumba menang")
	} else {
		fmt.Println("Tidak ada pemenang")
	}
}
