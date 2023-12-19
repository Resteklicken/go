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
	idx := slices.IndexFunc(allRomanNumerals, func(rn RomanNumeral) bool { return rn.Symbol == string([]rune(roman)[0]) })
	result := allRomanNumerals[idx].Value

	for _, c := range roman[1:] {
		idx := slices.IndexFunc(allRomanNumerals, func(rn RomanNumeral) bool { return rn.Symbol == string(c) })
		result += allRomanNumerals[idx].Value
	}
	return result

	/* for i, c := range []rune(roman[1:]) {
		prev := slices.IndexFunc(allRomanNumerals, func(rn RomanNumeral) bool { return rn.Symbol == string(c) })
		current := slices.IndexFunc(allRomanNumerals, func(rn RomanNumeral) bool { return rn.Symbol == string([]rune(roman)[i]) })
		switch {
		case allRomanNumerals[prev].Value >= allRomanNumerals[current].Value:
			result += allRomanNumerals[current].Value
		default:
			result = result + allRomanNumerals[current].Value - allRomanNumerals[prev].Value
		}
	} */
}
