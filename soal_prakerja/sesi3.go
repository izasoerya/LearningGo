package soalprakerja

func ArrayMerge(arrayA, arrayB []string) []string {
	final := make([]string, 0, len(arrayA)+len(arrayB))
	seen := make(map[string]bool)

	for _, name := range arrayA {				// for arrayA
		if !seen[name] {						// if !seen add to final
			final = append(final, name)
			seen[name] = true					// flag the name that have been seen
		}
	}
	for _, name := range arrayB {				// for arrayB
		if !seen[name] {						// if !seen add to final
			final = append(final, name)
			seen[name] = true					// flag the name that have been seen
		}
	}

	return final
}

func Mapping(input []string) map[string]int {
	counts := make(map[string]int)
	output := make(map[string]int)
	
	for _, str := range input {				// make dummy map [string]int
		counts[str]++
	}
	for str, count := range counts {		// if count int is >= 2 then save
		if count >= 2 {
			output[str] = count
		}
	}

	return output
}

func MunculSekali(angka string) []string {
	arrchar := []rune(angka)
	final := make([]string, 0, len(angka))
	seen := make(map[rune]int)

	for _, number := range string(arrchar) {			// count the same number
		seen[number] +=1
	}
	for _, number := range string(arrchar) {			// only save the n = 1
		if seen[number] == 1 {
			final = append(final, string(number))
		}
	}

	return final
}
