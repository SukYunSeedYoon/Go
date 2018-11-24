package main

import (
	"fmt"
	"io/ioutil"
	"runtime"
	"unicode/utf8"
	"unsafe"
)

func main() {

	var (
		x, y      int = 10, 30
		age, name     = 10, "maria"

		number32 uint   = 0xf0000000
		number16 uint16 = 0xffff
		number8  int8   = 0x7f
	)

	fmt.Println(x, y, age, name)
	fmt.Println(number32, number16, number8)

	fmt.Println("Hello world")

	fmt.Println(unsafe.Sizeof(x), unsafe.Sizeof(name))

	fmt.Println(unsafe.Sizeof(number8), unsafe.Sizeof(number16))

	var c, d = unsafe.Sizeof("안녕"), utf8.RuneCountInString("안녕")
	fmt.Println(c, d)
	var cod1, cod2, cod3 = 'a', "ab", "se edyoongogogogo"
	fmt.Println(unsafe.Sizeof(cod1), unsafe.Sizeof(cod2), unsafe.Sizeof(cod3))
	fmt.Println(len(cod2), len(cod3))

	fmt.Println("CPU Count : ", runtime.NumCPU())

	// b, err := ioutil.ReadFile("./hello.txt")
	// if err == nil{
	// 	fmt.Printf("%s", b)
	// }

	if b, err := ioutil.ReadFile("/hello.txt"); err == nil {
		fmt.Printf("%s", b)
	}

	// 기본형
	for index := 0; index < 5; index++ {
		fmt.Println(index)
	}

	// 증가 체크
	i := 0
	for i < 5 {

		fmt.Println(i)
		i++
	}

	// 	무한루프
	//for{
	//}

	// 레이블을 다는 것과 다른 차이점이 무엇일까?

	for i := 0; i < 3; i++ {
	Loop:
		for j := 0; j < 3; j++ {

			if j == 1 {
				break Loop
			}
			fmt.Println(i, j)
		}
	}

	fmt.Println("Hello, world")


	k := 3

	switch k{
	case 4:
		fmt.Println(k)
		fallthrough
	default:
		fmt.Println("nothing..")
	}

	//	손쉬운 다항 처리
	switch k {
	case 2,3,4:
		fmt.Println(k)
	case 5,6,7:
		fmt.Println(k)

	}

	//	variable이 switch에 붙어 있지 않음 if else처럼 사용하네..
	switch {
	case k >= 5 && k < 10:
		fmt.Println(k)
	}

}
