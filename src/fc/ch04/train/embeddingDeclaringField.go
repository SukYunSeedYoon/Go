package main

import "fmt"

type player struct {
	name  string
	email string
}

type admin struct {
	player

	level int
}

type manager struct {
	*player
	title string
}

type notifier interface {
	notify()
}

func (p *player) notify() {
	fmt.Printf("name : %s, email %s", p.name, p.email)
}

func (a *admin) notify() {
	fmt.Printf("admin!! name : %s, email %s\n", a.name, a.email)
}

func (m *manager) notify() {
	fmt.Printf("mgr!! name : %s, email %s\n", m.name, m.email)
}

func main() {

	ad := admin{
		player: player{
			name:  "seed",
			email: "seedyoon@gmail.com",
		},
		level: 199,
	}

	mn := manager{
		player: &player{
			name:  "haram",
			email: "haram@gmail.com",
		},
		title: "manager",
	}

	players := []notifier{&ad, &mn}

	for _, u := range players {
		u.notify()
	}

}
