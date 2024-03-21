package soal

import (
	"fmt"
)

func calculateBMI(weight, height float64) float64 {
	return weight / (height * height)
}

func SoalDua() {
	fmt.Println("\n-----Soal 2-----")
	// Data Uji 1
	markWeight := 78.0
	markHeight := 1.69
	johnWeight := 92.0
	johnHeight := 1.95

	markBMI := calculateBMI(markWeight, markHeight)
	johnBMI := calculateBMI(johnWeight, johnHeight)

	markHigherBMI := markBMI > johnBMI

	fmt.Println("-----Data Uji 1-----")
	fmt.Printf("BMI Mark: %.2f\n", markBMI)
	fmt.Printf("BMI John: %.2f\n", johnBMI)
	fmt.Printf("Apakah Mark memiliki BMI lebih tinggi dari John? %t\n\n", markHigherBMI)

	// Data uji 2
	markWeight = 95.0
	markHeight = 1.88
	johnWeight = 85.0
	johnHeight = 1.76

	markBMI = calculateBMI(markWeight, markHeight)
	johnBMI = calculateBMI(johnWeight, johnHeight)

	markHigherBMI = markBMI > johnBMI

	fmt.Println("-----Data Uji 1-----")
	fmt.Printf("BMI Mark: %.2f\n", markBMI)
	fmt.Printf("BMI John: %.2f\n", johnBMI)
	fmt.Printf("Apakah Mark memiliki BMI lebih tinggi dari John? %t\n", markHigherBMI)
}
