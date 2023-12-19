package main

import (
	"slices"
	"strings"
)

type RomanNumeral struct {
	Value  int
	Symbol string
}

var allRomanNumerals = []RomanNumeral{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func ConvertToRoman(arabic int) string {
	var result strings.Builder

	for _, numeral := range allRomanNumerals {
		for arabic >= numeral.Value {
			result.WriteString(numeral.Symbol)
			arabic -= numeral.Value
		}
	}

	return result.String()
}

func ConvertToArabic(roman string) int {
	var result int

	/* 	for _, c := range roman {
		idx := slices.IndexFunc(allRomanNumerals, func(rn RomanNumeral) bool { return rn.Symbol == string(c) })
		result += allRomanNumerals[idx].Value
	} */

	for i, c := range roman {
		if i == 0 {
			current := slices.IndexFunc(allRomanNumerals, func(rn RomanNumeral) bool { return rn.Symbol == string(c) })
			result += allRomanNumerals[current].Value
			continue
		}

		prev_idx := slices.IndexFunc(allRomanNumerals, func(rn RomanNumeral) bool { return rn.Symbol == string([]rune(roman)[i-1]) })
		current_idx := slices.IndexFunc(allRomanNumerals, func(rn RomanNumeral) bool { return rn.Symbol == string(c) })
		prev_val := allRomanNumerals[prev_idx].Value
		current_val := allRomanNumerals[current_idx].Value

		if current_val > prev_val {
			result = result + current_val - prev_val*2
			continue
		}

		result += current_val
	}
	return result
}
