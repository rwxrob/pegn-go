package pegn

import (
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

// Convert changes the given string from one base to another.
func Convert(v string, from, to int) (string, error) {
	i, err := strconv.ParseInt(v, from, 64)
	if err != nil {
		return "", err
	}
	return strings.ToUpper(strconv.FormatInt(i, to)), nil
}

// Convert changes the given string in the given base to hex.
func ConvertToHex(v string, base int) (string, error) {
	hex, err := Convert(v, base, 16)
	if err != nil {
		return "", err
	}
	if len(hex)%2 != 0 {
		hex = fmt.Sprintf("0%s", hex)
	}
	return hex, nil
}

// ConvertToInt changes the given string in the given base to an integer.
func ConvertToInt(v string, base int) (int, error) {
	i, err := strconv.ParseInt(v, base, 64)
	return int(i), err
}

func ConvertToRuneString(v string, base int) (string, error) {
	// Hex: (00)07
	value, err := ConvertToHex(v, base)
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
