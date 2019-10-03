package main

import "fmt"

// FizzBuzzを実行する中で、"Fizz", "Buzz", "FizzBuzz", それ以外の数値 がそれぞれ何回出力されたか
// カウントし、出力してください、ただし、それ以外の数値は "others" としてカウントします
// 例: 1以上16未満のFizzBuzzの場合
//
//   Fizz 4
//   Buzz 2
//   FizzBuzz 1
//   others 8
//

func main() {
	m := make(map[string]int, 3)
	for i := 1; i <= 100; i++ {
		switch {
		case i % 15 == 0:
			fmt.Println("FizzBuzz")
			m["FizzBuzz"]++
		case i % 5 == 0:
			fmt.Println("Buzz")
			m["Buzz"]++
		case i % 3 == 0:
			fmt.Println("Fizz")
			m["Fizz"]++
		default: 
			fmt.Println(i)
			m["default"]++
		}
	}
	fmt.Println(m);
}