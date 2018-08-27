package composition

import "fmt"

// go는 상속 불가. 대신 임베딩을 사용해 조합 가능하고 재사용이 가능한 구조체를 만들 수 있다.
// go는 인터페이스나 구조체 안에 타입을 임베딩 할 수 있게 해준다.
// 임베딩을 통해 우리는 내부 타입에서 선언된 메서드를 외부 타입으로 전달 할 수 있다

// 타입을 임베드하면, 해당 타입의 메서드는 외부 타입의 메서드가 되지만,
// 메서드가 실행될 때의 메서드를 외부 타입으로 전달할 수 있다

type Human struct {
	FirstName string
	LastName  string
	CanSwim   bool
}

// Amy는 Human type으로 임베딩, 따라서 Human의 메서드 집합에 속하는 메서드를 실행할 수 있음
type Amy struct {
	Human
}

// 같음, 내가 무엇을 받았냐에 따라 사용 가능
type Alan struct {
	Human
}

func (h *Human) Name() {
	fmt.Printf("Hello my name is %s, %s", h.FirstName, h.LastName)
}

func (h *Human) Swim() {

	if h.CanSwim == true {
		fmt.Println("I can swim")
	} else {
		fmt.Println("I can not swim.")
	}
}
