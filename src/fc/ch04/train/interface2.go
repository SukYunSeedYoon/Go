package main

import "fmt"

type printer interface {
	print()
}

type user struct {
	name string
}

func (u user) print() {
	fmt.Printf("User name %s\n", u.name)
}

func main() {

	// 타입의 user 값을 만든다
	u := user{"bill"}

	// 값과 포인터를 가지고 printer 인터페이스의 슬라이드를 만든다
	entities := []printer{
		// user 값의 복사를 인터페이스 값에 저장
		u,

		// user 값의 주소를 복사해서 인터페이스 값에 저장
		&u,
	}

	// name 필드를 바꾼다
	u.name = "Bill CHG"

	// entities 를 반복처리하여 각 요소에 대해 print 메서드를 호출한다
	for _, e := range entities {
		e.print()
	}

	// user 값의 복사를 인터페이스에 저장한 경우 원래 값의 변호를 받지 않는다
	// user 값의 주소를 복사해서 인터페이스에 저장한 경우는 원래 값의 변화를 그대로 반영

}
