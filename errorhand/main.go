package main

import "fmt"

func simplefun(a, b float64) (float64, string) {
	if b == 0 {
		return 0, "here i am returning string "
	}
	return a / b, "nil"
}

func simplefunc(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("error due to denominator, it is 0")
	}
	return a / b, nil
}

func main() {
	res, error := simplefun(20, 0)
	if error != "nil" {
		fmt.Println(error)
	} else {
		fmt.Println(res)
	}

	res, err := simplefunc(20, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}

	ans, _ := simplefunc(10, 4)
	fmt.Println("a divided by b:", ans)

}
