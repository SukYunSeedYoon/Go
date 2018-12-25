package main

import (
	"fc/lib"
	"fmt"
)

func main() {

	bird := lib.GetDuck()
	bird.Quack()

	record := lib.Record{}

	record.Quack()

	// export 되지 않은 형태를 사용하려하면 역시 에러
	// 짧은 선언으로는 리턴해서받을 수 있음 와이?

	cartoon := lib.GetDonald()
	fmt.Println(cartoon)

}
