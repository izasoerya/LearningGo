package soalprakerja

func ArrayMerge(arrayA, arrayB []string) []string {
	final := make([]string, 0, len(arrayA)+len(arrayB))
	seen := make(map[string]bool)

	for _, name := range arrayA {
		if !seen[name] {
			final = append(final, name)
			seen[name] = true
		}
	}
	for _, name := range arrayB {
		if !seen[name] {
			final = append(final, name)
			seen[name] = true
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
