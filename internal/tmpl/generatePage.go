package tmpl

import "github.com/ropfoo/gothulhu/internal/model"

func GeneratePage(c model.Character) string {
	return `
		<!DOCTYPE html>
		<html>
			<head>
				<title>Generate Character</title>
				<link rel="stylesheet" href="/styling/main.css">
			</head>
			<script>
				function resetStatValues() {
					const stats = document.getElementById("stats");
					stats.querySelectorAll("input").forEach(input => {
						input.value = "";
					});
				}
			</script>
			<body>
				<div class="title-container flex-column">
					<h1>Call of Cthulhu</h1>
					<h2>Character Generator</h2>
				</div>
				<div class="character-container">
					` + characterForm(c) + `
					` + characterResult(c) + `
				</div>
			</body>
		</html>
	`
}
