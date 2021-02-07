package pegn

import (
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

// convert changes the given string from one base to another.
func convert(v string, from, to int) (string, error) {
	i, err := strconv.ParseInt(v, from, 64)
	if err != nil {
		return "", err
	}
	return strings.ToUpper(strconv.FormatInt(i, to)), nil
}

// convert changes the given string in the given base to hex.
func convertToHex(v string, base int) (string, error) {
	hex, err := convert(v, base, 16)
	if err != nil {
		return "", err
	}
	if len(hex)%2 != 0 {
		hex = fmt.Sprintf("0%s", hex)
	}
	return hex, nil
}

// convertToInt changes the given string in the given base to an integer.
func convertToInt(v string, base int) (int, error) {
	i, err := strconv.ParseInt(v, base, 64)
	return int(i), err
}

func convertToRuneString(v string, base int) (string, error) {
	// Hex: (00)07
	value, err := convertToHex(v, base)
	if err != nil {
		return "", err
	}
	if len(value)%4 == 2 {
		// Add zeros if not present.
		value = fmt.Sprintf("00%s", value)
	}
	// Runes: '\u0000' || '\U00000000' || int32
	if len(value) <= 4 {
		value = fmt.Sprintf("0x%s", value)
	} else {
		i, err := strconv.ParseInt(value, 16, 32)
		if err != nil {
			return "", err
		}
		if i <= utf8.MaxRune {
			value = fmt.Sprintf("0x%s", value)
		} else {
			value = fmt.Sprintf("%d", i)
		}
	}
	return value, nil
}
