package main

import (
	"fmt"
	"math/rand"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func rollDice(sides int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(sides) + 1
}

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Dice Roller")

	// Widgets
	resultLabel := widget.NewLabel("Roll the Dice!")
	sidesEntry := widget.NewEntry()
	sidesEntry.SetPlaceHolder("Enter number of sides (e.g., 6)")

	rollButton := widget.NewButton("Roll", func() {
		sides := 6 // default

		if sidesEntry.Text != "" {
			_, err := fmt.Sscanf(sidesEntry.Text, "%d", &sides)
			if err != nil || sides < 2 {
				resultLabel.SetText("Invalid number of sides!")
				return
			}
		}

		result := rollDice(sides)
		resultLabel.SetText(fmt.Sprintf("🎲 Rolled: %d", result))
	})

	// Layout
	content := container.NewVBox(
		widget.NewLabel("Dice Roller"),
		sidesEntry,
		rollButton,
		resultLabel,
	)

	myWindow.SetContent(content)
	myWindow.Resize(fyne.NewSize(300, 200))
	myWindow.ShowAndRun()
}
