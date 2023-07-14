package main

import (
	"errors"
	"strings"
	"os"
	"bufio"
	"fmt"
	"strconv"
)

var text, rome_result string
var split_text []string
var index, result, operand_counter, num1, num2, res_int, result_calculate int
var is_rome1, is_rome2 bool

func main() {
	reader := bufio.NewReader(os.Stdin)

	rim_num := map[string](int){
		"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5, "VI": 6, "VII": 7, "VIII": 8,
		"IX": 9, "X": 10, "XI": 11, "XII": 12, "XIII": 13, "XIV": 14, "XV": 15,
		"XVI": 16, "XVII": 17, "XVIII": 18, "XIX": 19, "XX": 20, "XXI": 21, "XXII": 22,
		"XXIII": 23, "XXIV": 24, "XXV": 25, "XXVI": 26, "XXVII": 27, "XXVIII": 28, "XXIX": 29,
		"XXX": 30, "XXXI": 31, "XXXII": 32, "XXXIII": 33, "XXXIV": 34, "XXXV": 35, "XXXVI": 36,
		"XXXVII": 37, "XXXVIII": 38, "XXXIX": 39, "XL": 40, "XLI": 41, "XLII": 42, "XLIII": 43,
		"XLIV": 44, "XLV": 45, "XLVI": 46, "XLVII": 47, "XLVIII": 48, "XLIX": 49, "L": 50,
		"LI": 51, "LII": 52, "LIII": 53, "LIV": 54, "LV": 55, "LVI": 56, "LVII": 57, "LVIII": 58,
		"LIX": 59, "LX": 60, "LXI": 61, "LXII": 62, "LXIII": 63, "LXIV": 64, "LXV": 65, "LXVI": 66,
		"LXVII": 67, "LXVIII": 68, "LXIX": 69, "LXX": 70, "LXXI": 71, "LXXII": 72, "LXXIII": 73,
		"LXXIV": 74, "LXXV": 75, "LXXVI": 76, "LXXVII": 77, "LXXVIII": 78, "LXXIX": 79,
		"LXXX": 80, "LXXXI": 81, "LXXXII": 82, "LXXXIII": 83, "LXXXIV": 84, "LXXXV": 85,
		"LXXXVI": 86, "LXXXVII": 87, "LXXXVIII": 88, "LXXXIX": 89, "XC": 90, "XCI": 91,
		"XCII": 92, "XCIII": 93, "XCIV": 94, "XCV": 95, "XCVI": 96, "XCVII": 97, "XCVIII": 98,
		"XCIX": 99, "C": 100,
}

	fmt.Println("Введите выражение:\n")
	text, _ := reader.ReadString(byte('\n'))
	text = strings.TrimSpace(text)

	split_result := split(text)


	num1, err1 := strconv.Atoi(split_result[0])
	if err1 != nil {
		num1 = rim_num[split_result[0]]
		is_rome1 = true
	}

	num2, err2 := strconv.Atoi(split_result[2])
	if err2 != nil {
		num2 = rim_num[split_result[2]]
		is_rome2 = true
	}


	if num1 == 0 || num2 == 0 {
		err := errors.New("Вводите только римские и арабские цифры от 1 до 10!")
		fmt.Println(err)
		os.Exit(1)
	}
	if (num1 < 1 || num1 > 10) || (num2 < 1 || num2 > 10) {
		err := errors.New("Операнды должны быть в диапозоне от 1 до 10!")
		fmt.Println(err)
		os.Exit(1)
	}


	result_calculate = calculate(num1, num2, split_result[1])

	if is_rome1 && is_rome2 {
		if result_calculate > 1 {
			for key, value := range(rim_num){
				if value == result_calculate {
					rome_result = key
					break
				}
			}
			fmt.Println(rome_result)
		} else {
			err := errors.New("Римские числа не могут быть меньше 1!")
			fmt.Println(err)
			os.Exit(1)
		}

	} else if is_rome1 || is_rome2 {
		err := errors.New("Вводите только римские и арабские цифры от 1 до 10!")
		fmt.Println(err)
		os.Exit(1)
	} else {
		fmt.Println(result_calculate)
	}
	
}

func split(text string) []string{

	split_text = strings.Split(text, " ")
	if len(split_text) != 3 {
		err := errors.New("Неверное выражение! Может быть только два операнда и один оператор!")
		fmt.Println(err)
		os.Exit(1)
	}

	return split_text
			
}

func calculate(x int, y int, operator string) int{

	if operator == "+"{
		result = x + y
	}

	if operator == "-"{
		result = x - y
	}

	if operator == "*"{
		result = x * y
	}

	if operator == "/"{
		result = x / y
	}

	return result
}
