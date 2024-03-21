package utils

func AvgSlice(slice *[]float32) float32 {
	var total float32
	for _, value := range *slice {
		total += value
	}
	return total / float32(len(*slice))
}
