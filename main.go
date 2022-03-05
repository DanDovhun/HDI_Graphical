package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func popup(app fyne.App, message string) {
	window := app.NewWindow("Error")

	window.SetContent(widget.NewLabel(message))
	window.Show()
}

func main() {
	myApp := app.New()
	homeWindow := myApp.NewWindow("Human Development Index")
	homeWindow.SetMaster()
	countriesWindow := myApp.NewWindow("Countries")

	content := container.NewVBox(
		widget.NewButton("Continents", func() {
			homeWindow.Hide()
			continentWindow := myApp.NewWindow("Continents")
			continentWindow.SetMaster()

			output := "Continents:\n"

			continents, err := GetContinents()

			if err != nil {
				popup(myApp, "Cannot open json file")
			} else {
				continentList := continents.Sort()

				for _, i := range continentList {
					output += fmt.Sprintf("Name: %v\n", i.Continent)
					output += fmt.Sprintf("Countries: %v\n", i.Countries)
					output += fmt.Sprintf("Average HDI: %v\n\n", i.HdiAverage)
				}
			}

			continentWindow.SetContent(widget.NewLabel(output))
			continentWindow.Resize(fyne.NewSize(300, 500))
			continentWindow.Show()
		}),

		widget.NewButton("Countries", func() {
			homeWindow.Hide()
			countriesWindow.SetMaster()

			countriesWindow.Show()
		}),
	)

	homeWindow.Resize(fyne.NewSize(400, 80))
	homeWindow.SetContent(content)
	homeWindow.ShowAndRun()
}
