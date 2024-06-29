package utils

import (
	"math/rand"
	"regexp"
	"strings"
	"time"
)

func IsValidURL(url *string) bool {
	var re = regexp.MustCompile(`^(https?|ftp):\/\/(www\.)?([a-zA-Z0-9-]{1,63}\.){1,}[a-zA-Z]{2,6}(:[0-9]{1,5})?(\/[^\s]*)?$`)
	return re.MatchString(*url)
}

func IsValidExpiry(exp *string) bool {
	var numberRegex = regexp.MustCompile(`^-?\d+(\.\d+)?$`)
	return numberRegex.MatchString(*exp)
}

func GenerateShortURL() string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, 6)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func TrimURL(url string) string {
	lastIndex := strings.LastIndex(url, "/getshort/")
	if lastIndex == -1 {
		return url
	}
	trimmedURL := url[:lastIndex]
	return trimmedURL + "/getorigin/"
}
