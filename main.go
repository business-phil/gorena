package main

import (
	"fmt"
	"math/rand"
	"time"
)

type combatant struct {
	name      string
	hp        int
	maxDamage int
}

func (combatant *combatant) attack(opponent *combatant) (int, bool) {
	damage := rand.Intn(combatant.maxDamage) + 1

	opponent.hp = opponent.hp - damage

	return damage, opponent.hp <= 0
}

func main() {
	var name string

	fmt.Println("What is your name?")

	// TODO: Accept name with spaces
	if _, err := fmt.Scanf("%s", &name); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Welcome, %s\n", name)

	player := &combatant{name, 50, 8}
	opponent := &combatant{"Opponent", 30, 10}

	rand.Seed(time.Now().Unix())

	for player.hp > 0 && opponent.hp > 0 {
		playerDamage, opponentIsDead := player.attack(opponent)
		fmt.Printf("%s hit %s for %d damage\n", player.name, opponent.name, playerDamage)
		if opponentIsDead {
			fmt.Println("You win!!")
			return
		}

		opponentDamage, playerSsDead := opponent.attack(player)
		fmt.Printf("%s hit %s for %d damage\n", opponent.name, player.name, opponentDamage)
		if playerSsDead {
			fmt.Println("You lose!!")
			return
		}
	}
}
