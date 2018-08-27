package main

import (
	"github.com/SukYunSeedYoon/Go/src/abstraction"
	"github.com/SukYunSeedYoon/Go/src/composition"
	"github.com/SukYunSeedYoon/Go/src/encapsulation"
	"github.com/SukYunSeedYoon/Go/src/polymorphism"
	"github.com/SukYunSeedYoon/Go/src/type-and-method"
)

func main() {

	type_and_method.RunMain()

	e := encapsulation.Encapsulation{}
	e.Expose()
	e.UnHide()

	h := polymorphism.Hillary{}
	t := polymorphism.Trump{}

	h.Slogan()
	t.Slogan()

	// 인터페이스 전달해 주는 경우, 래퍼런스 타입으로 전달해주고 있음
	polymorphism.SaySlogan(&h)
	polymorphism.SaySlogan(&t)

	// 암묵적으로 만족하므로 이 역시 된다..
	polymorphism.SloganSayer.Slogan(&h)

	// composition
	amy := composition.Amy{
		Human: composition.Human{
			FirstName: "Amy",
			LastName:  "Chen",
			CanSwim:   true,
		},
	}

	alan := composition.Alan{
		Human: composition.Human{
			FirstName: "Alan",
			LastName:  "Chen",
			CanSwim:   false,
		},
	}

	// Human의 메서드 집합은 Amy 타입으로 전달
	amy.Name()
	amy.Swim()
	alan.Name()
	alan.Swim()

	// abstract
	hillary := abstraction.Hillary{}
	trump := abstraction.Trump{}

	// slogan의 힐러리, 트럼프 구현체는 추상화 되어 있다
	// 대신, Campaign은 단지 이게 SloganSayer이고 따라서 slogan을 호출할 수 있음을 알고 있음

	c1 := abstraction.Campagin{SloganSayer: &hillary}
	c2 := abstraction.Campagin{SloganSayer: &trump}

	c1.Slogan()
	c2.Slogan()

	// SloganSayer를 SaySlogan의 파라미터로 주입 할 수 있음
	abstraction.SaySlogan(&hillary)
	abstraction.SaySlogan(&trump)

	// campagin또한 h, t를 사용 가능
	abstraction.SaySlogan(c1)
	abstraction.SaySlogan(c2)

}
