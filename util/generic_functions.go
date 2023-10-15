package util

import "regexp"

func IsValidPhoneNumber(phone string) bool {
	re := regexp.MustCompile(`^[0-9]+$`)
	return re.MatchString(phone)
}
