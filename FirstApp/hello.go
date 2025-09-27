package firstapp

const englishHelloPrefix = "Hello "
const turkish = "Turkish"
const turkishPrefix = "Merhaba "
const french = "French"
const frechPrefix = "Bonjour"

func Hello(key string, language string) string {
	if key == "" {
		key = "World"
	}

	return greetingprefix(language) + key
}

func greetingprefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frechPrefix
	case turkish:
		prefix = turkishPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}
