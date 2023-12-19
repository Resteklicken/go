package main

import (
	"slices"
	"strings"
)

type RomanNumeral struct {
	Value  int
	Symbol string
}

type RomanNumerals []RomanNumeral

func (r RomanNumerals) ValueOf(symbol string) int {
	idx := slices.IndexFunc(r, func(rn RomanNumeral) bool { return rn.Symbol == symbol })
	return r[idx].Value
}

var allRomanNumerals = RomanNumerals{
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
	var (
		result      int
		current_val int
		prev_val    int
	)
	for i, c := range roman {
		current_val = allRomanNumerals.ValueOf(string(c))
		if i > 0 && current_val > prev_val {
			result = result + current_val - prev_val*2
			continue
		}
		result += current_val
		prev_val = current_val
	}
	return result
}
