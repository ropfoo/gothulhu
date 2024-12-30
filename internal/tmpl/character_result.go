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
            <div class="character-header flex-row">
                <h1>` + c.Name + `</h1>
            </div>
            <div class="character-info flex-row">
                <div class="character-info-left">
                    <div class="character-hp">
                        <p>HP ` + strconv.Itoa(c.HP) + `</p>
                    </div>
                        <p>Age: ` + strconv.Itoa(c.Age) + `</p>
                    <div class="character-damage-bonus">
                        <p>Damage Bonus: </p>
                        <p>Damage: ` + c.DamageBonus.Damage + `</p>
                        <p>Stature: ` + c.DamageBonus.Stature + `</p>
                    </div>
                </div>
                <div class="character-stats">
                    <div class="character-stats-row">
                        <p>STR</p>
                        <p class="character-stats-value">` + strconv.Itoa(c.Stats.STR) + `</p>
                    </div>
                    <div class="character-stats-row">
                        <p>DEX</p>
                        <p class="character-stats-value">` + strconv.Itoa(c.Stats.DEX) + `</p>
                    </div>
                    <div class="character-stats-row">
                        <p>CON</p>
                        <p class="character-stats-value">` + strconv.Itoa(c.Stats.CON) + `</p>
                    </div>
                    <div class="character-stats-row">
                        <p>INT</p>
                        <p class="character-stats-value">` + strconv.Itoa(c.Stats.INT) + `</p>
                    </div>
                    <div class="character-stats-row">
                        <p>WIS</p>
                        <p class="character-stats-value">` + strconv.Itoa(c.Stats.WIS) + `</p>
                    </div>
                    <div class="character-stats-row">
                        <p>CHA</p>
                        <p class="character-stats-value">` + strconv.Itoa(c.Stats.CHA) + `</p>
                    </div>
                </div>
            </div>
		</section>
	`
}
