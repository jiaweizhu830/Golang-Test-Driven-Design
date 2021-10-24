package numeral

import (
	"strings"
)

type RomanNumeral struct {
	Value  uint16
	Symbol string
}

var allRomanNumerals = []RomanNumeral{
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func ConvertToRoman(arabic uint16) string {

	// efficiently build a string using Write methods
	// minimize memory copying
	var result strings.Builder

	// for i := arabic; i > 0; i-- {
	// 	if i == 5 {
	// 		result.WriteString("V")
	// 		break
	// 	}
	// 	if i == 4 {
	// 		result.WriteString("IV")
	// 		break
	// 	}
	// 	result.WriteString("I")
	// }

	// for arabic > 0 {
	// 	switch {
	// 	case arabic > 9:
	// 		result.WriteString("X")
	// 		arabic -= 10
	// 	case arabic > 8:
	// 		result.WriteString("IX")
	// 		arabic -= 9
	// 	case arabic > 4:
	// 		// append the current letter to the tail
	// 		result.WriteString("V")
	// 		arabic -= 5
	// 	case arabic > 3:
	// 		result.WriteString("IV")
	// 		arabic -= 4
	// 	default:
	// 		result.WriteString("I")
	// 		arabic--
	// 	}
	// }

	// be more OOP
	for _, numeral := range allRomanNumerals {
		for arabic >= numeral.Value {
			result.WriteString(numeral.Symbol)
			arabic -= numeral.Value
		}
	}

	return result.String()
}

type RomanNumerals []RomanNumeral

/*
In Go, when indexing strings, you get "byte"
*/
func (r RomanNumerals) ValueOf(symbols ...byte) uint16 {
	symbol := string(symbols)
	for _, s := range r {
		if s.Symbol == symbol {
			return s.Value
		}
	}
	return 0
}

func ConvertToArabic(roman string) uint16 {
	// need to change the logic
	var total uint16 = 0
	for range roman {
		total++
	}
	return total
}
