package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/manifoldco/promptui"
)

type Combatant struct {
	name      string
	maxHp     int
	currentHp int
	maxDamage int
}

func (combatant *Combatant) Attack(opponent *Combatant) (int, bool) {
	rand.Seed(time.Now().Unix())

	damage := rand.Intn(combatant.maxDamage) + 1

	opponent.currentHp -= damage

	fmt.Printf("%s hit %s for %d damage\n", combatant.name, opponent.name, damage)

	return damage, opponent.currentHp <= 0
}

func (combatant *Combatant) Heal() (int, bool) {
	rand.Seed(time.Now().Unix())

	atFullHealth := false
	healthRestored := rand.Intn(combatant.maxDamage) + rand.Intn(combatant.maxDamage) + 2

	if healthRestored+combatant.currentHp > combatant.maxHp {
		healthRestored = combatant.maxHp - combatant.currentHp
		atFullHealth = true
	}

	combatant.currentHp += healthRestored

	fmt.Printf("%s healed for %d hp\n", combatant.name, healthRestored)

	return healthRestored, atFullHealth
}

func main() {
	namePrompt := promptui.Prompt{
		Label: "What is your name? ",
		Validate: func(input string) error {
			if input == "" {
				return errors.New("Must enter name")
			}
			return nil
		},
	}

	name, err := namePrompt.Run()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Welcome, %s\n", name)

	player := &Combatant{name, 50, 50, 8}
	opponent := &Combatant{"Opponent", 30, 30, 10}

	actionPrompt := promptui.Select{
		Label: "What will you do? ",
		Items: []string{"Attack", "Heal"},
	}

	for player.currentHp > 0 && opponent.currentHp > 0 {
		_, result, err := actionPrompt.Run()
		if err != nil {
			fmt.Println(err)
			return
		}

		if result == "Attack" {
			_, opponentIsDead := player.Attack(opponent)
			if opponentIsDead {
				fmt.Println("You win!!")
				return
			}
		} else if result == "Heal" {
			player.Heal()
		}

		_, playerIsDead := opponent.Attack(player)
		if playerIsDead {
			fmt.Println("You lose!!")
			return
		}
	}
}
