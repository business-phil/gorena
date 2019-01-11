package combatant

import (
	"fmt"
	"math/rand"
	"time"
)

type Combatant struct {
	Name      string
	MaxHp     int
	CurrentHp int
	MaxDamage int
}

func (combatant *Combatant) Attack(opponent *Combatant) (int, bool) {
	rand.Seed(time.Now().Unix())

	damage := rand.Intn(combatant.MaxDamage) + 1

	opponent.CurrentHp -= damage

	fmt.Printf("%s hit %s for %d damage\n", combatant.Name, opponent.Name, damage)

	return damage, opponent.CurrentHp <= 0
}

func (combatant *Combatant) Heal() (int, bool) {
	rand.Seed(time.Now().Unix())

	atFullHealth := false
	healthRestored := rand.Intn(combatant.MaxDamage) + rand.Intn(combatant.MaxDamage) + 2

	if healthRestored+combatant.CurrentHp > combatant.MaxHp {
		healthRestored = combatant.MaxHp - combatant.CurrentHp
		atFullHealth = true
	}

	combatant.CurrentHp += healthRestored

	fmt.Printf("%s healed for %d hp\n", combatant.Name, healthRestored)

	return healthRestored, atFullHealth
}
