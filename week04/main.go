package main

import (
	"fmt"
	"math"
	"reflect"
	"strings"
)

func main() {
	// var a int //선언
	// a = 7     // 값 지정

	// var a int = 7 //선언 및 값 지정
	// var a = 7     //선언 및 값 지정

	a := 7
	fmt.Println(a, reflect.TypeOf(a))

	// b :=8.34 //float64
	var b float32 = 8.34
	fmt.Println(a * int(b))
	fmt.Println(float32(a) > b)

	fmt.Println(b, reflect.TypeOf(b))

	var c7 string //변수명은 숫자로 시작할 수 없음
	var d int
	var e bool
	var f float64
	var G = 99
	// koreanzombie := "정찬성" //문법적으로 문제없지만 camel케이스를 사용하는 관례를 따르자
	koreanzombie := "정찬성"
	fmt.Println(koreanzombie)

	fmt.Println(c7, d, e, f, G)

	fmt.Println('Z', '2', '\n', '김', '인', '하') //rune literals(int32) 90,50, 10, 44608, 51064, 54616
	fmt.Println(reflect.TypeOf('Z'), reflect.TypeOf('2'), reflect.TypeOf("Hi"), reflect.TypeOf(4.99), reflect.TypeOf(false))
	// fmt.Println(math.Floor("삼,오"), math.Ceil("이백십칠쩜칠"), math.Sqrt("sixteen"))
	// fmt.Println(strings.Title(("3.141592")))0
	fmt.Println(math.Floor(2.17), math.Ceil(2.17), math.Sqrt(16))
	fmt.Println(strings.Title(("open sourse\tp\tograming\n\"go\"!"))) //문자열 "" 사용
	//strings.Title(("오픈소스 프로그래밍"))
	//fmt.Println("OpenSource Programming~","GO")
}
