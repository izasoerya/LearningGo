package soalprakerja

func ArrayMerge(arrayA, arrayB []string) []string {
	/*
		Works by read element side by side, then
		delete the j (behind) element from the array
	*/
	final := append(arrayA, arrayB...)
	var temp int

	for i := 0; i < len(final); i++ { // Counter for A
		for j := i + 1; j < len(final)-1; j++ { // Counter for B
			if final[i] == final[j] {
				temp = j // Detect element that hold same value
				if temp != 0 {
					final = append(final[:temp], final[temp+1:]...) // Remove element that hold same value
				}
			}
		}
	}

	return final
}

func Mapping(input []string) map[string]int {
	counts := make(map[string]int)
	for _, str := range input {
		counts[str]++
	}
	output := make(map[string]int)

	for str, count := range counts {
		if count >= 2 {
			output[str] = count
		}
	}

	return output
}
