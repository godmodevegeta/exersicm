package logs


// Application identifies the application emitting the given log.
func Application(log string) string {
	table := make(map[int]rune)
	table[0] = '❗'
	table[1] = '🔍'
	table[2] = '☀'

	for _, runeValue := range log {
		if runeValue == table[0] {
			return "recommendation"
		} else if runeValue == table[1]{
			return "search"
		} else if runeValue == table[2] {
			return "weather"
		}
	}
	return "default"

}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	runeSlice := []rune(log)
	for index, runeValue := range runeSlice {
		if runeValue == oldRune {
			runeSlice[index] = newRune
		}
	}
	return string(runeSlice)
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	runeSlice := []rune(log)
	return len(runeSlice) <= limit
}
