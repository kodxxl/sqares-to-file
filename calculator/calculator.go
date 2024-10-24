package calculator

func GetSqaresOf(source []int) []int {
	result := []int{}

	for _, value := range source {
		sqare := value * value
		result = append(result, sqare)
	}

	return result
}
