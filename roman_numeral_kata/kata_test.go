package main

import (
	"fmt"
	"testing"
)

func TestConversionToRoman(t *testing.T) {
	cases := []struct {
		Arabic int
		Want   string
	}{
		{1, "I"},
		{2, "II"},
		{3, "III"},
		{4, "IV"},
		{5, "V"},
		{6, "VI"},
		{9, "IX"},
		{10, "X"},
		{11, "XI"},
		{14, "XIV"},
		{15, "XV"},
		{39, "XXXIX"},
		{40, "XL"},
		{41, "XLI"},
		{44, "XLIV"},
		{49, "XLIX"},
		{50, "L"},
		{100, "C"},
		{90, "XC"},
		{400, "CD"},
		{500, "D"},
		{900, "CM"},
		{1000, "M"},
		{1984, "MCMLXXXIV"},
		{3999, "MMMCMXCIX"},
		{2014, "MMXIV"},
		{1006, "MVI"},
		{798, "DCCXCVIII"},
	}

	for _, test := range cases {
		t.Run(fmt.Sprintf("%d -> %q", test.Arabic, test.Want), func(t *testing.T) {
			got := ConvertToRoman(test.Arabic)

			if got != test.Want {
				t.Errorf("got %q want %q", got, test.Want)
			}
		})
	}
}

func TestConversionToArabic(t *testing.T) {
	cases := []struct {
		Want  int
		Roman string
	}{
		{1, "I"},
		{2, "II"},
		{3, "III"},
		{4, "IV"},
		{5, "V"},
		{6, "VI"},
		{9, "IX"},
		{10, "X"},
		{11, "XI"},
		{14, "XIV"},
		{15, "XV"},
		{39, "XXXIX"},
		{40, "XL"},
		{41, "XLI"},
		{44, "XLIV"},
		{49, "XLIX"},
		{50, "L"},
		{100, "C"},
		{90, "XC"},
		{400, "CD"},
		{500, "D"},
		{900, "CM"},
		{1000, "M"},
		{1984, "MCMLXXXIV"},
		{3999, "MMMCMXCIX"},
		{2014, "MMXIV"},
		{1006, "MVI"},
		{798, "DCCXCVIII"},
	}
	for _, test := range cases {
		t.Run(fmt.Sprintf("%q -> %d", test.Roman, test.Want), func(t *testing.T) {
			got := ConvertToArabic(test.Roman)

			if got != test.Want {
				t.Errorf("got %d want %d", got, test.Want)
			}
		})
	}
}
