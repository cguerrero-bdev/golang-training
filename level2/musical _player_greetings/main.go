package main

import "fmt"

type Trumpeter struct {
	Name string
}

func (trumpeter *Trumpeter) Greetings() {
	fmt.Printf("Hi, my name is %s and I'm a trumpeter\n", trumpeter.Name)
}

type Violinist struct {
	Name string
}

func (violinist *Violinist) Greetings() {
	fmt.Printf("Hi, my name is %s and I'm a violinist\n", violinist.Name)
}

type MusicalPlayer interface {
	Greetings()
}

func main() {

	musicalPlayers := make([]MusicalPlayer, 0)

	musicalPlayers = append(musicalPlayers, &Trumpeter{"Mike"})
	musicalPlayers = append(musicalPlayers, &Violinist{"John"})
	musicalPlayers = append(musicalPlayers, &Violinist{"Lois"})
	musicalPlayers = append(musicalPlayers, &Trumpeter{"Peter"})

	for i := 0; i < len(musicalPlayers); i++ {
		musicalPlayers[i].Greetings()
	}
}
