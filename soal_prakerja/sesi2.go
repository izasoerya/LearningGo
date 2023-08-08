package soalprakerja

func DetectPrimeNumber(input int) string {
	yes := "Prime number"
	no := "Not prime number"

	if input == 0 || input == 1 {
		return no
	}
	for i := 2; i < input; i++ {
		if input%i == 0 {
			return no
		}
	}
	return yes
}

func Detect7Multiplication(input int) string {
	yes := "Seven multiplication"
	no := "Not seven multiplication"

	if input % 7 != 0 {
		return no
	}
	return yes
}

func CalculateAreaTrapeziod(bottom, up, height int) int {
	return (bottom + up )/2 * (height)
}
