package gorena

import (
	"fmt"
	"math/rand"
	"time"
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
