package tmpl

import (
	"strconv"

	"github.com/ropfoo/gothulhu/internal/model"
)

// return html template for character
func characterForm(c model.Character) string {
	return `
	<form action="/generate#result" method="post" accept-charset="utf-8" class="character-form">
		<section class="character-section general-section">
			<div class="character-section-header">
				<h3>General</h3>
			</div>
			<div class="input-row">
				<label>Name</label>
				<input id="name" type="text" name="name" placeholder="Name" value="` + c.Name + `">
			</div>
			<div class="input-row">
				<label>Age</label>
				<input id="age" type="number" name="age" placeholder="Age" value="` + strconv.Itoa(c.Age) + `">
			</div>
		</section>
		<section id="stats" class="character-section stats-section">
			<div class="character-section-header">
				<h3>Stats</h3>
				<button type="button" onclick="resetStatValues()">Reset</button>
			</div>
			<div class="input-row">
				<label>Strength</label>
				<input type="number" name="str" value="` + strconv.Itoa(c.Stats.STR) + `">
			</div>
			<div class="input-row">
				<label>Dexterity</label>
				<input type="number" name="dex" value="` + strconv.Itoa(c.Stats.DEX) + `">
			</div>
			<div class="input-row">
				<label>Constitution</label>
				<input type="number" name="con" value="` + strconv.Itoa(c.Stats.CON) + `">
			</div>
			<div class="input-row">
				<label>Size</label>
				<input type="number" name="siz" value="` + strconv.Itoa(c.Stats.SIZ) + `">
			</div>
			<div class="input-row">
				<label>Intelligence</label>
				<input type="number" name="intl" value="` + strconv.Itoa(c.Stats.INT) + `">
			</div>
			<div class="input-row">
				<label>Wisdom</label>
				<input type="number" name="wis" value="` + strconv.Itoa(c.Stats.WIS) + `">
			</div>
			<div class="input-row">
				<label>Charisma</label>
				<input type="number" name="cha" value="` + strconv.Itoa(c.Stats.CHA) + `">
			</div>
			<button type="submit">Generate</button>
		</section>
	</form>
	`
}
