package main

import (
	"fmt"
)

func main() {
	city := "Seoul"
	weather, err := getWeather(city)
	if err != nil {
		fmt.Println("service call faild")
	}

	fmt.Println("==========================================================================")
	fmt.Println(weather)
	fmt.Println("==========================================================================")

	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")

	fmt.Println("엔터 키를 눌러 종료하세요...")
    fmt.Scanln() // 사용자 입력을 기다립니다
}
