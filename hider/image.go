package hider

func HideMessage(msg string, jpg []byte) []byte {
	return append(jpg, []byte(msg))
}

func ReadMessage(jpg []byte) string {
	var start int
	for x := len(jpg) - 2; x > 0; x = x - 1 {
		if jpg[x] == 0xFF && jpg[x+1] == 0xD9 {
			start = x + 2
			break
		}
	}
	if start < len(jpg) {
		return string(jpg[start:])
	}

	return "No Message."
}