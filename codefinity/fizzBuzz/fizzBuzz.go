package main

import (
	"fmt"
	"strconv"
)

func main() {
	res, err := fizzBuzz(0, ruleBasic)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}

func validateN(num int) error {
	if num < 1 {
		return fmt.Errorf("Invalid value of n")
	}
	return nil
}
func ruleBasic(num int) (string, error) {
	err := validateN(num)
	if err != nil {
		return "", err
	}
	var res string
	if num%3 == 0 && num%5 == 0 {
		res = "FizzBuzz"
	} else if num%3 == 0 {
		res = "Fizz"
	} else if num%5 == 0 {
		res = "Buzz"
	} else {
		res = strconv.Itoa(num)
	}
	return res, nil
}

func fizzBuzz(n int, rule func(int) (string, error)) ([]string, error) {
	var err error
	res := make([]string, n, n)
	for i := 0; i < n; i++ {
		res[i], err = rule(i + 1)
		if err != nil {
			return []string{}, err
		}

	}

	return res, nil
}
