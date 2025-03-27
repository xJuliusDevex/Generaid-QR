package service

import (
	"strconv"
	"strings"
)

func Status(name, nro string) bool {
	if name != "" && nro != "" {
		return true
	}
	return false
}
func Verify_int(num string) bool {
	_, err := strconv.Atoi(num)
	if err != nil {
		return false
	}
	return true
}
func Full_space(letra string) string {
	if 3 > len(letra) {
		for i := len(letra); i < 3; i++ {
			letra = "0" + letra
		}

	}
	return letra
}
func Mayuscula(letra string) string {
	return strings.ToUpper(letra)
}
