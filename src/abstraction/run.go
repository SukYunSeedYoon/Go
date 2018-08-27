package abstraction

import "fmt"

// 추상화란 내부적으로 어떻게 동작하는지 모르는 채 사용법만 알고 작업한다는 것을 의미

// 구조체 내에 구주체를 임베딩하는 것과 유사하게, 구조체 내 인터페이스를 임베딩 할 수 있다
// 한 인터페이스를 만족하는 모든 타입들은 또한 해당 인터페이스를 갖는다는 것

type SloganSayer interface {
	Slogan()
}

// campaign 은 자기 자신의 인스턴스에 slogansayer를 받을 수 있음
// campaign은 또한 sloganslayer 인터페이스를 구현하고 있으므로 slogansayer이기도 함
// 이는 체이닝 시 유용

type Campagin struct {
	SloganSayer
}

// saySlogan은 파라미터로 Campaign을 받을 수 있게 됨, 같으니
func SaySlogan(s SloganSayer) {
	s.Slogan()
}

// 힐러리는 slogan sayer 인터페이스 구현
// 그녀는 slogansayer임
type Hillary struct {
}

func (h *Hillary) Slogan() {
	fmt.Println("Strong together")
}

// 	트럼프도 마찬가지
type Trump struct {
}

func (t *Trump) Slogan() {
	fmt.Println("Make American great again")
}
