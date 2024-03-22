package soal

import (
	"fmt"
)

// function to calculate BMI with parameter weight and height (pointer) return BMI
func calculateBMI(weight, height *float32) float32 {
	return *weight / (*height * *height)
}

func SoalDua() {
	fmt.Println("\n-----Soal 2-----")
	// Data Uji 1
	var markWeight float32 = 78.0
	var markHeight float32 = 1.69
	var johnWeight float32 = 92.0
	var johnHeight float32 = 1.95

	// calculate BMI for Mark and John
	markBMI := calculateBMI(&markWeight, &markHeight)
	johnBMI := calculateBMI(&johnWeight, &johnHeight)

	markHigherBMI := markBMI > johnBMI

	// print BMI for Mark and John and comparation who has higher BMI
	fmt.Println("-----Data Uji 1-----")
	fmt.Printf("BMI Mark: %.2f\n", markBMI)
	fmt.Printf("BMI John: %.2f\n", johnBMI)
	fmt.Printf("Apakah Mark memiliki BMI lebih tinggi dari John? %t\n\n", markHigherBMI)

	// Data uji 2
	markWeight = 95.0
	markHeight = 1.88
	johnWeight = 85.0
	johnHeight = 1.76

	// calculate BMI for Mark and John
	markBMI = calculateBMI(&markWeight, &markHeight)
	johnBMI = calculateBMI(&johnWeight, &johnHeight)

	markHigherBMI = markBMI > johnBMI

	// print BMI for Mark and John and comparation who has higher BMI
	fmt.Println("-----Data Uji 1-----")
	fmt.Printf("BMI Mark: %.2f\n", markBMI)
	fmt.Printf("BMI John: %.2f\n", johnBMI)
	fmt.Printf("Apakah Mark memiliki BMI lebih tinggi dari John? %t\n", markHigherBMI)
}
