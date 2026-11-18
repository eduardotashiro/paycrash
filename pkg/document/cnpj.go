package document

import (
	"regexp"
)

const baseValue = '0'

var (
	cnpjPeso      = []int{6, 5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}
	regexFormat   = regexp.MustCompile(`^[A-Z0-9]{12}$`)
	regexFormatDV = regexp.MustCompile(`^\d{2}$`)
)

func ValidateCNPJ(c string) bool {
	cnpj := removeCharacters(c)

	if len(cnpj) != 14 {
		return false
	}



	base := cnpj[:12] // base cnpj
	dv := cnpj[12:]   // digito verificador

	if !regexFormat.MatchString(base) || !regexFormatDV.MatchString(dv) {
		return false
	}

	return true
}

func calcDV(base string) string {
	return "pareiaqui"
}

func calculaDigitoCNPJ(doc string) int {
	return 42
}

func allSame(s string) bool {
	for i := 1; i < len(s); i++ {
		if s[i] != s[0] {
			return false
		}
	}
	return len(s) > 0
}
//*              ________  ________   ________      ___
//*             |\   ____\|\   ___  \|\   __  \    |\  \
//*             \ \  \___|\ \  \\ \  \ \  \|\  \   \ \  \
//*              \ \  \    \ \  \\ \  \ \   ____\__ \ \  \
//*               \ \  \____\ \  \\ \  \ \  \___|\  \\_\  \
//*                \ \_______\ \__\\ \__\ \__\  \ \________\
//*                 \|_______|\|__| \|__|\|__|   \|________||
