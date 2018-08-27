package type_and_method

import "fmt"

type SomeStruct struct {
	Field string
}

// foo는 SomeStruct의 메서드 집합에 속함
// (s* SomeStruct) 는 SomeStruct 포인터에 대한 리시버이다
func (s *SomeStruct) foo(field string) {
	s.Field = field
}

func RunMain() {

	s := SomeStruct{}

	s.foo("a")
	fmt.Printf(s.Field) // a

	s.foo("b") // b

	fmt.Printf(s.Field)
}
