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

	scoreKoala := []float32{88, 91, 110}
	scoreLumbaLumba := []float32{96, 108, 89}

	avgScoreKoala := avgSlice(&scoreKoala)
	avgScoreLumbaLumba := avgSlice(&scoreLumbaLumba)

	// print average score each team (2 decimal)
	fmt.Printf("Rata-rata score koala: %.2f\n", avgScoreKoala)
	fmt.Printf("Rata-rata score lumba-lumba: %.2f\n", avgScoreLumbaLumba)

	fmt.Print("Hasilnya : ")
	// compare average score each team
	if avgScoreKoala == avgScoreLumbaLumba {
		fmt.Println("Seri")
	} else if avgScoreKoala > avgScoreLumbaLumba {
		fmt.Println("Koala menang")
	} else {
		fmt.Println("Lumba-lumba menang")
	}

	fmt.Println("-----Bonus 1: -----")

	scoreKoala = []float32{109, 95, 123}
	scoreLumbaLumba = []float32{97, 112, 101}

	avgScoreKoala = avgSlice(&scoreKoala)
	avgScoreLumbaLumba = avgSlice(&scoreLumbaLumba)

	// print average score each team (2 decimal) minimal score 100
	fmt.Printf("Rata-rata score koala: %.2f\n", avgScoreKoala)
	fmt.Printf("Rata-rata score lumba-lumba: %.2f\n", avgScoreLumbaLumba)

	fmt.Print("Hasilnya : ")
	// compare average score each team with minimal score 100
	if avgScoreKoala == avgScoreLumbaLumba {
		fmt.Println("Seri")
	} else if (avgScoreKoala > avgScoreLumbaLumba) && (avgScoreKoala >= 100) {
		fmt.Println("Koala menang")
	} else if (avgScoreKoala < avgScoreLumbaLumba) && (avgScoreLumbaLumba >= 100) {
		fmt.Println("Lumba-lumba menang")
	} else {
		fmt.Println("Tidak ada pemenang")
	}

	fmt.Println("-----Bonus 2: -----")

	scoreKoala = []float32{109, 95, 106}
	scoreLumbaLumba = []float32{97, 112, 101}

	avgScoreKoala = avgSlice(&scoreKoala)
	avgScoreLumbaLumba = avgSlice(&scoreLumbaLumba)

	// print average score each team (2 decimal) minimal score 100
	fmt.Printf("Rata-rata score koala: %.2f\n", avgScoreKoala)
	fmt.Printf("Rata-rata score lumba-lumba: %.2f\n", avgScoreLumbaLumba)

	fmt.Print("Hasilnya : ")
	// compare average score each team with minimal score 100
	if (avgScoreKoala == avgScoreLumbaLumba) && (avgScoreKoala >= 100) && (avgScoreLumbaLumba >= 100) {
		fmt.Println("Seri")
	} else if (avgScoreKoala > avgScoreLumbaLumba) && (avgScoreKoala >= 100) {
		fmt.Println("Koala menang")
	} else if (avgScoreKoala < avgScoreLumbaLumba) && (avgScoreLumbaLumba >= 100) {
		fmt.Println("Lumba-lumba menang")
	} else {
		fmt.Println("Tidak ada pemenang")
	}
}
