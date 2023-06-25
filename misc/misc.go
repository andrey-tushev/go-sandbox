package main

import "fmt"

func main() {
	{
		str := "hello world"
		b := []byte(str) // копируется
		b[2] = '*'
		str2 := string(b)
		fmt.Println(str, str2)
		fmt.Println(&str, &str2) // разные адреса
	}

	{
		str := "мир"
		for pos, rune := range str {
			fmt.Println(pos, rune)
		}

		for i := 0; i < len(str); i++ {
			fmt.Println(i, str[i])
		}
	}

	{
		str := []byte("hello world")
		l := 0
		r := len(str) - 1
		for l < r {
			str[l], str[r] = str[r], str[l]
			l++
			r--
		}
		fmt.Println(string(str))
	}

}
