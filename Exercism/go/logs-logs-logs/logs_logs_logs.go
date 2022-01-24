package logs

// Application identifies the application emitting the given log.
func Application(log string) string {
	application := map[rune]string{
		'\u2757':     "recommendation",
		'\U0001f50d': "search",
		'\u2600':     "weather",
	}
	for _, character := range log {
		if char, ok := application[character]; ok {
			return char
		}
	}
	return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	runeList := []rune(log)
	for idx, char := range runeList {
		if char == oldRune {
			runeList[idx] = newRune
		}
	}
	return string(runeList)
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	runeList := []rune(log)
	return len(runeList) <= limit
}
