package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	var num float64

	fmt.Println("Enter a number: ")
	fmt.Scanf("%f", &num)
	fmt.Printf("Result: %s\n", extenso(num))
}

var singluarClasses = []string{
	"mil",
	"milhão",
	"bilhão",
	"trilhão",
	"quatrilhão",
}

var pluralClasses = []string{
	"mil",
	"milhões",
	"bilhões",
	"trilhões",
	"quatrilhões",
}

var units = map[int]string{
	1:  "um",
	2:  "dois",
	3:  "três",
	4:  "quatro",
	5:  "cinco",
	6:  "seis",
	7:  "sete",
	8:  "oito",
	9:  "nove",
	10: "dez",
	11: "onze",
	12: "doze",
	13: "treze",
	14: "catorze",
	15: "quinze",
	16: "dezesseis",
	17: "dezessete",
	18: "dezoito",
	19: "dezenove",
}

var tens = map[int]string{
	2: "vinte",
	3: "trinta",
	4: "quarenta",
	5: "cinquenta",
	6: "sessenta",
	7: "setenta",
	8: "oitenta",
	9: "noventa",
}

var hundreds = map[int]string{
	0: "cem",
	1: "cento",
	2: "duzentos",
	3: "trezentos",
	4: "quatrocentos",
	5: "quinhentos",
	6: "seiscentos",
	7: "setecentos",
	8: "oitocentos",
	9: "novecentos",
}

func extenso(n float64) string {
	if n == 0 {
		return "zero"
	}

	intParts := []string{}
	decParts := []string{}

	nAbs := math.Abs(n)
	nAbsStr := fmt.Sprintf("%f", nAbs)

	nInt := int(math.Trunc(nAbs))

	decIndex := strings.Index(nAbsStr, ".")

	if decIndex > -1 {
		decStr := nAbsStr[decIndex+1:]

		decStrInt, err := strconv.Atoi(decStr)

		if err != nil {
			fmt.Println("Error:", err)

			return ""
		}

		if decStrInt > 0 {

			for strings.HasSuffix(decStr, "0") {
				decStr = decStr[:len(decStr)-1]
			}

			if len(decStr) > 2 {
				decStr = decStr[:2] + "." + decStr[2:]
			}

			dec, err := strconv.ParseFloat(decStr, 64)

			if err != nil {
				fmt.Println("Error:", err)

				return ""
			}

			decInt := int(math.Round(dec))

			if decInt >= 100 {
				decInt = 0
				nInt += 1
			}

			words := []string{}

			if decInt >= 20 {
				ten := int(math.Trunc(float64(decInt) / 10))

				words = append(words, tens[ten])

				decInt -= ten * 10
			}

			if decInt > 0 {
				words = append(words, units[decInt])
			}

			if len(words) > 0 {

				verbalDec := strings.Join(words, " e ") + " centavo"

				if int(math.Round(dec)) > 1 {
					verbalDec += "s"
				}

				if nInt == 0 {
					verbalDec += " de real"
				} else {
					verbalDec = "e " + verbalDec
				}

				decParts = append(decParts, verbalDec)
			}
		}
	}

	if nInt > 0 {
		nIntStr := fmt.Sprintf("%d", nInt)
		nIntStrLen := len(nIntStr)

		n_parts := math.Ceil(float64(nIntStrLen) / 3)

		for i := 0; i < int(n_parts); i++ {
			end := nIntStrLen - 3*i
			start := end - 3

			var partStr string

			if nIntStrLen < 3 {
				partStr = nIntStr
			} else {
				partStr = nIntStr[start:end]
			}

			part, err := strconv.Atoi(partStr)

			if err != nil {
				fmt.Println("Error:", err)

				break
			}

			if part == 0 {
				continue
			}

			words := []string{}

			if part == 100 {
				words = append(words, hundreds[0])
				part = 0
			}

			if part > 100 {
				hundred := int(math.Trunc(float64(part) / 100))

				words = append(words, hundreds[hundred])

				part -= hundred * 100
			}

			if part >= 20 {
				ten := int(math.Trunc(float64(part) / 10))

				words = append(words, tens[ten])

				part -= ten * 10
			}

			if part > 0 {
				words = append(words, units[part])
			}

			verbalPart := strings.Join(words, " e ")

			if i == 1 && verbalPart == "um" {
				verbalPart = singluarClasses[0]
			} else if i > 1 && verbalPart == "um" {
				verbalPart += fmt.Sprintf(" %s", singluarClasses[i-1])
			} else if i >= 1 {
				verbalPart += fmt.Sprintf(" %s", pluralClasses[i-1])
			}

			intParts = append([]string{verbalPart}, intParts...)
		}

		if n < 0 {
			intParts = append([]string{"menos"}, intParts...)
		}

		if nInt > 1 {
			intParts = append(intParts, "reais")
		} else {
			intParts = append(intParts, "real")
		}
	}

	parts := append(intParts, decParts...)

	return strings.Join(parts, " ")
}
