
package conversion

// Importing necessary package.
import (
	"errors"
	"regexp"
)

var isValid = regexp.MustCompile("^[0-1]{1,}$").MatchString

// BinaryToDecimal() function that will take Binary number as string,
// and return it's Decimal equivalent as integer.
func BinaryToDecimal(binary string) (int, error) {
	if !isValid(binary) {
		return -1, errors.New("not a valid binary string")
	}
	if len(binary) > 32 {
		return -1, errors.New("binary number must be in range 0 to 2^(31-1)")
	}
	var result, base int = 0, 1
	for i := len(binary) - 1; i >= 0; i-- {
		if binary[i] == '1' {
			result += base
		}
		base *= 2
	}
	return result, nil
}
