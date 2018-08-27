package polymorphism

import "fmt"

// 다형성은 클래스들이 공통 인터페이스를 공유하면서 다른 기능을 갖는 객체 지향 프로그래밍의 패턴을 묘사
// 인터페이스가 암시적으로 충족. 인터페이스 또한 타입이다. - 많은 의미

// 인터페이스가 암시적으로 충족 : 인터페이스의 모든 메서드가 어떤 타입의 메서드 집합에 모두 포함되어 있을 경우, 해당 타입은 인터페이스를 만족한다
//							go는 implements 키워드가 없다

// 인터페이스는 타입이다 : 어떤 타입이 한 인터페이스를 만족하면, 그 타입은 또한 타입이 만족하는 인터페이스에 의해 제한되는 모든 타입을 만족한다

// 인터페이스 타입 정의
type SloganSayer interface {
	Slogan()
}

// SaySlogan은 SloganSayer타입을 파라메터로 받음
func SaySlogan(sayer SloganSayer) {
	sayer.Slogan()
}

// hillary는 slogan함수를 구현함으로써 암묵적으로 slogansayer 인터페이스를 만족!
// 따라서 hillary도 sloganSayer 타입이다
type Hillary struct {
}

func (h *Hillary) Slogan() {
	fmt.Println("Strong together")
}

// trump에 대해서도 동일하게 적용

type Trump struct {
}

func (t *Trump) Slogan() {
	fmt.Println("Make America great again")
}
