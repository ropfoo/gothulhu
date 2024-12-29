package tmpl

import (
	"strconv"

	"github.com/ropfoo/gothulhu/internal/model"
)

func characterResult(c model.Character) string {
	if c.Name == "" {
		return ""
	}
	return `
		<section id="result" class="character-result">
            <div>
                <h1>` + c.Name + `</h1>
                <p>Age: ` + strconv.Itoa(c.Age) + `</p>
            </div>
            <div class="character-hp">
                <p>HP: ` + strconv.Itoa(c.HP) + `</p>
            </div>
            <div class="character-damage-bonus">
                <p>Damage Bonus: </p>
                <p>Damage: ` + c.DamageBonus.Damage + `</p>
                <p>Stature: ` + c.DamageBonus.Stature + `</p>
            </div>
            <div class="character-stats">
                <div class="character-stats-row">
                    <p>STR: </p>
                    <p>` + strconv.Itoa(c.Stats.STR) + `</p>
                </div>
                <div class="character-stats-row">
                    <p>DEX: </p>
                    <p>` + strconv.Itoa(c.Stats.DEX) + `</p>
                </div>
                <div class="character-stats-row">
                    <p>CON: </p>
                    <p>` + strconv.Itoa(c.Stats.CON) + `</p>
                </div>
                <div class="character-stats-row">
                    <p>INT: </p>
                    <p>` + strconv.Itoa(c.Stats.INT) + `</p>
                </div>
                <div class="character-stats-row">
                    <p>WIS: </p>
                    <p>` + strconv.Itoa(c.Stats.WIS) + `</p>
                </div>
                <div class="character-stats-row">
                    <p>CHA: </p>
                    <p>` + strconv.Itoa(c.Stats.CHA) + `</p>
                </div>
            </div>
		</section>
	`
}
