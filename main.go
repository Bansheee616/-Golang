package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var roman = [...]string{"O", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X", "XI", "XII", "XIII", "XIV", "XV", "XVI", "XVII", "XVIII", "XIX", "XX",
	"XXI", "XXII", "XXIII", "XXIV", "XXV", "XXVI", "XXVII", "XXVIII", "XXIX", "XXX", "XXXI", "XXXII", "XXXIII", "XXXIV", "XXXV", "XXXVI", "XXXVII", "XXXVIII", "XXXIX", "XL",
	"XLI", "XLII", "XLIII", "XLIV", "XLV", "XLVI", "XLVII", "XLVIII", "XLIX", "L", "LI", "LII", "LIII", "LIV", "LV", "LVI", "LVII", "LVIII", "LIX", "LX",
	"LXI", "LXII", "LXIII", "LXIV", "LXV", "LXVI", "LXVII", "LXVIII", "LXIX", "LXX",
	"LXXI", "LXXII", "LXXIII", "LXXIV", "LXXV", "LXXVI", "LXXVII", "LXXVIII", "LXXIX", "LXXX",
	"LXXXI", "LXXXII", "LXXXIII", "LXXXIV", "LXXXV", "LXXXVI", "LXXXVII", "LXXXVIII", "LXXXIX", "XC",
	"XCI", "XCII", "XCIII", "XCIV", "XCV", "XCVI", "XCVII", "XCVIII", "XCIX", "C"}

func main() {
	reader := bufio.NewReader(os.Stdin)
	exps, _ := reader.ReadString('\n')
	exps = strings.TrimSpace(exps)
	exps2 := strings.Split(exps, " ")

	if len(exps2) != 3 {
		fmt.Println("Ошибка, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *). ")
		return
	}

	numOneType := "arabic"
	numTwoType := "arabic"
	numOne, errO := strconv.Atoi(exps2[0])
	if errO != nil {

	}
	numTwo, errT := strconv.Atoi(exps2[2])
	if errT != nil {

	}

	for i := 0; i < 11; i++ {
		if exps2[0] == roman[i] {
			numOneType = "roman"
			numOne = i
		}
		if exps2[2] == roman[i] {
			numTwoType = "roman"
			numTwo = i
		}

	}

	if numOneType != numTwoType {
		fmt.Println("Ошибка,совершать операции можно только с цифрами из одной системы ")
		return
	}

	if numTwo < 1 || numOne < 1 || numTwo > 10 || numOne > 10 {
		fmt.Println("Ошибка,только целые числа от одного до десяти ")
		return
	}
	var result int

	switch exps2[1] {
	case "+":
		result = numOne + numTwo
	case "-":
		result = numOne - numTwo
	case "*":
		result = numOne * numTwo
	case "/":
		result = numOne / numTwo
	default:
		fmt.Println("Ошибка ввода")
	}

	if numOneType == "roman" && numTwoType == "roman" {
		if result < 0 {
			fmt.Println("Ошибка, так как в римской системе нет отрицательных чисел.")
		} else if result == 0 {
			fmt.Println("Римскими цифрами можно записать только натуральные числа и ноль к ним не относится.")
		} else {
			fmt.Println(roman[result])
		}

	} else {
		fmt.Println(result)
	}

}
