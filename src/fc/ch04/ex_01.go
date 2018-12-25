package main

import "fmt"

type player struct {
	name   string
	atBats float32
	hits   float32
	mark   bool
}

func (p *player) BatAvg() float32 {

	return (float32)(p.hits / p.atBats)

}

func (p *player) Mark() bool {
	p.mark = !p.mark
	return p.mark
}

func main() {

	players := []player{
		{"seed", 12, 10, false},
		{"hr", 18, 10, false},
		{"mj", 22, 10, false},
	}

	fmt.Printf("%p %p %p\n", &players[0], &players[1], &players[2])

	for _, p := range players {
		fmt.Printf("[%p] p, name: %s, avg [%f]\n", &p, p.name, p.BatAvg())
		p.Mark()
		fmt.Println(p)
	}

	for i := 0; i < len(players); i++ {
		fmt.Printf("[%p] name: %s, avg [%f]\n", &players[i], players[i].name, players[i].BatAvg())
		fmt.Println(players[i])

		players[i].Mark()
	}

	for _, p := range players {
		fmt.Println(p)
	}

}
