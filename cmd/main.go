package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	strIn, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		fmt.Println("Data reading error")
		return
	}
	strIn = strings.Trim(strIn, "\n")
	slice := strings.Split(strIn," ")
	if len(slice) != 3 {
		fmt.Println("Incorrect input data")
		return
	}

	symbol := slice[1]
	symbols := map[string]bool{"+":true, "-":true, "*":true, "/":true}
	_, ok := symbols[symbol]
	if !ok {
		fmt.Println("Incorrect input data")
		return
	}

	num1, err1 := strconv.Atoi(slice[0])
	num2, err2 := strconv.Atoi(slice[2])
	if err1 == nil && err2 == nil && num1 >= 0 && num1 <= 10 && num2 >= 0 && num2 <= 10 {
		res, err := calculation(num1, num2, symbol)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(res)
		return
	}
	num1, err1 = RomToInt(slice[0])
	num2, err2 = RomToInt(slice[2])
	if err1 == nil && err2 == nil {
		res, err := calculation(num1, num2, symbol)
		if err != nil {
			fmt.Println(err)
			return
		}
		if res < 1 {
			fmt.Println("The result can only be positive Roman numbers")
			return
		}
		resRom := IntToRom(res)
		fmt.Println(resRom)
		return
	}
	fmt.Println("Incorrect input data")

}

func calculation(num1, num2 int, symbol string) (res int, err error) {
	switch symbol{
	case "+":
		res = num1 + num2
	case "-":
		res = num1 - num2
	case "*":
		res = num1 * num2
	case "/":
		if num2 == 0{
			return 0, errors.New("Zero division error")
		}
		res = num1 / num2
	}
	return res, nil
}

func RomToInt(num string) (int, error) {
	mapRom := map[string]int{"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5, "VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10}
	if mapRom[num] !=0 {
		return mapRom[num], nil
	}
	return 0, errors.New("Incorrect Roman number")
}

func IntToRom(num int) (res string) {
	sliceRom := []string{"C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	sliceInt := []int{100, 90, 50, 40, 10, 9, 5, 4, 1}
	for i := 0; num > 0; i++ {
		for sliceInt[i] <= num {
			res += sliceRom[i]
			num -= sliceInt[i]
		}
	}
	return res
}