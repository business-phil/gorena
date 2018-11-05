package gorena

import (
	"errors"
	"fmt"

	"github.com/manifoldco/promptui"
)

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
