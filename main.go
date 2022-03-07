package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func popup(app fyne.App, window fyne.Window, message string) {
	popup := app.NewWindow("Error")

	popup.SetMaster()
	popup.SetContent(container.NewVBox(
		widget.NewLabel(message),
		widget.NewButton("Ok", func() {
			popup.Hide()
			window.SetMaster()
			window.Show()
		}),
	))
	popup.Show()
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
				popup(myApp, continentWindow, "Cannot open json file")
			} else {
				continentList := continents.Sort()

				for _, i := range continentList {
					output += fmt.Sprintf("Name: %v\n", i.Continent)
					output += fmt.Sprintf("Countries: %v\n", i.Countries)
					output += fmt.Sprintf("Average HDI: %v\n\n", i.HdiAverage)
				}
			}

			content := container.NewVBox(
				widget.NewLabel(output),
				widget.NewButton("Back", func() {
					continentWindow.Hide()
					homeWindow.Show()
					homeWindow.SetMaster()
				}),
			)

			continentWindow.SetContent(content)
			continentWindow.Resize(fyne.NewSize(300, 500))
			continentWindow.Show()
		}),

		widget.NewButton("Countries", func() {
			homeWindow.Hide()
			countriesWindow.SetMaster()

			content := container.NewVBox(
				widget.NewButton("Search for a country", func() {
					searchWindow := myApp.NewWindow("Search for a country")
					input := widget.NewEntry()
					output := widget.NewLabel("")

					content := container.NewVBox(
						input,
						widget.NewButton("Find", func() {
							countries, err := GetCountries()

							if len(input.Text) == 0 {
								popup(myApp, searchWindow, "Please enter something")

								return
							}

							if err != nil {
								popup(myApp, searchWindow, "Cannot get countries")

								return
							}

							lst := countries.SortByCountry()
							country, index, errr := Search(lst, input.Text)

							if errr != nil {
								popup(myApp, searchWindow, fmt.Sprintf("Cannot find a country named '%v'", input.Text))
							}

							fmt.Println(country)

							output.SetText(country.Country +
								"\nContinent: " + country.Continent +
								fmt.Sprintf("\nHuman Development Index: %v\n", country.Hdi) +
								fmt.Sprintf("Position: %v", index))
						}),

						widget.NewButton("Back", func() {
							searchWindow.Hide()
							countriesWindow.Show()
							countriesWindow.SetMaster()
						}),

						output,
					)

					searchWindow.SetContent(content)
					searchWindow.Resize(fyne.NewSize(300, 600))
					searchWindow.Show()
					searchWindow.SetMaster()
					countriesWindow.Hide()
				}),

				widget.NewButton("Show statistics", func() {
					statsWindow := myApp.NewWindow("Statistics")
					statsWindow.SetMaster()
					countriesWindow.Hide()

					countries, err := GetCountries()

					if err != nil {
						popup(myApp, statsWindow, "Cannot get countries")

						return
					}

					lst := countries.SortByHdi()

					quartiles := GetQuartiles(lst)
					first, second, third := GetRealQuartiles(lst)

					firstCountry := lst[first]
					secondCountry := lst[second]
					thirdCountry := lst[third]

					real := "Real Quartiles:\n" +
						fmt.Sprintf("First quartile: %v (HDI: %v)\n", firstCountry.Country, firstCountry.Hdi) +
						fmt.Sprintf("Second quartile: %v (HDI: %v)\n", secondCountry.Country, secondCountry.Hdi) +
						fmt.Sprintf("Third quartile: %v (HDI: %v)\n", thirdCountry.Country, thirdCountry.Hdi) +
						fmt.Sprintf("IQR = %v", thirdCountry.Hdi-firstCountry.Hdi)

					statsWindow.SetContent(container.NewVBox(
						widget.NewLabel(fmt.Sprintf(
							"Quartiles:\nFirst quartile: %v\nSecond quartile: %v\nThird quartile: %v\nIQR = %v",
							round(quartiles.first, 3), round(quartiles.second, 3),
							round(quartiles.third, 3), round(quartiles.third-quartiles.first, 3),
						)),

						widget.NewLabel(real),
						widget.NewButton("Back", func() {
							statsWindow.Hide()
							countriesWindow.Show()
							countriesWindow.SetMaster()
						}),
					))
					statsWindow.Show()
				}),

				widget.NewButton("Back", func() {
					countriesWindow.Hide()
					homeWindow.Show()
					homeWindow.SetMaster()
				}),
			)

			countriesWindow.SetContent(content)
			countriesWindow.Show()
		}),

		widget.NewButton("Quit", func() {
			homeWindow.Close()
		}),
	)

	homeWindow.Resize(fyne.NewSize(400, 80))
	homeWindow.SetContent(content)
	homeWindow.ShowAndRun()
}
