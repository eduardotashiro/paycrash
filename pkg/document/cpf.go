package document

import (
	"strconv"
)

func ValidateCPF(c string) bool {
	cpf := removeCharacters(c)

	if len(cpf) != 11 {
		return false
	}

	counter := 0
	for k := 0; k < len(cpf); k++ {
		if cpf[0] == cpf[k] {
			counter++
		}
		if counter == len(cpf) {
			return false
		}
	}

	//calculando primeiro dígito verificador
	sum1 := 0
	for i := range 9 {
		digit, err := strconv.Atoi(string(cpf[i]))
		if err != nil {
			return false
		}
		sum1 += digit * (10 - i)
	}

	rest1 := (sum1 * 10) % 11
	if rest1 == 10 {
		rest1 = 0
	}

	firstDigit, err := strconv.Atoi(string(cpf[9]))
	if err != nil {
		return false
	}

	if rest1 != firstDigit {
		return false
	}

	//calculando o segundo dígito verificador
	sum2 := 0
	for i := range 10 {
		digit, err := strconv.Atoi(string(cpf[i]))
		if err != nil {
			return false
		}
		sum2 += digit * (11 - i)
	}

	rest2 := (sum2 * 10) % 11
	if rest2 == 10 {
		rest2 = 0
	}

	secondDigit, err := strconv.Atoi(string(cpf[10]))
	if err != nil {
		return false
	}

	if rest2 != secondDigit {
		return false
	}

	// cpf válido
	return true
}


func calculaDigitoCPF(cpf string) int {


for i := 9; i < 11; i++ {
	soma:= 0
	for j := 0; j < i; j++ {
		// csm
	}
}


// 	sum := 0
//
// 	for i :=  len(cpf) - 1 {
// 		digit, err := strconv.Atoi(string(cpf[i]))
// 		if err != nil {
// 			return false
// 		}
// 		sum1 += digit * (10 - i)
// 	}
//
// 	rest1 := (sum1 * 10) % 11
// 	if rest1 == 10 {
// 		rest1 = 0
// 	}
//
// 	firstDigit, err := strconv.Atoi(string(cpf[9]))
// 	if err != nil {
// 		return false
// 	}
//
// 	if rest1 != firstDigit {
// 		return false
// 	}
//
// 	//calculando o segundo dígito verificador
// 	sum2 := 0
// 	for i := range 10 {
// 		digit, err := strconv.Atoi(string(cpf[i]))
// 		if err != nil {
// 			return false
// 		}
// 		sum2 += digit * (11 - i)
// 	}
//
// 	rest2 := (sum2 * 10) % 11
// 	if rest2 == 10 {
// 		rest2 = 0
// 	}
//
// 	secondDigit, err := strconv.Atoi(string(cpf[10]))
// 	if err != nil {
// 		return false
// 	}
//
// 	if rest2 != secondDigit {
// 		return false
// 	}
}
//*                ________  ________  ________
//*              |\   ____\|\   __  \|\  _____\
//*              \ \  \___|\ \  \|\  \ \  \__/
//*               \ \  \    \ \   ____\ \   __\
//*                \ \  \____\ \  \___|\ \  \_|
//*                 \ \_______\ \__\    \ \__\
//*                  \|_______|\|__|     \|__|
