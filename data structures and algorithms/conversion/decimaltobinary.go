

package conversion

// Importing necessary package.
import (
	"errors"
	"strconv"
)

// Reverse() function that will take string,
// and returns the reverse of that string.
func Reverse(str string) string {
	rStr := []rune(str)
	for i, j := 0, len(rStr)-1; i < len(rStr)/2; i, j = i+1, j-1 {
		rStr[i], rStr[j] = rStr[j], rStr[i]
	}
	return string(rStr)
}

// DecimalToBinary() function that will take Decimal number as int,
// and return it's Binary equivalent as string.
func DecimalToBinary(num int) (string, error) {
	if num < 0 {
		return "", errors.New("integer must have +ve value")
	}
	if num == 0 {
		return "0", nil
	}
	var result string = ""
	for num > 0 {
		result += strconv.Itoa(num & 1)
		num >>= 1
	}
	return Reverse(result), nil
}
