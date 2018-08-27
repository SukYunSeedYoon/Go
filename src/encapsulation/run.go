package encapsulation

import "fmt"

// 클래스 기반 oop에서는 캡슐화가 private나 public변수 /메서드를 통해 이루어짐
// go에서는 패키지 수준에서 캡슐화가 이루어짐
// 캡슐화는 퍼블릭 메서드에 접근 제한을 둠으로써 데이터 구현체를 숨기는 매커니즘이다.

// 이 구조체는 패키지 밖으로 노출 될 수 있음
type Encapsulation struct{}

// Expose 메서드는 패키지 밖에 노출 될 수 있음
func (e *Encapsulation) Expose() {
	fmt.Println("Ha Im Expose")
}

// hide는 내부에서만 사용 가능
func (e *Encapsulation) hide() {
	fmt.Println("Sh this is super secret")
}

// Unhid는 노출되지 않은 hide 메서드를 사용
func (e *Encapsulation) UnHide() {
	e.hide()
	fmt.Println("Pack!")
}

func RunMain() {

	e := Encapsulation{}
	e.Expose()
	e.UnHide()
}
