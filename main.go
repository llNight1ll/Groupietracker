package main

import (
	"fmt"

	"groupie/data"

	"groupie/gui"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"

	"fyne.io/fyne/v2/container"

	"fyne.io/fyne/v2/widget"
)

func main() {

	gui.SearchContainer = container.NewVBox()
	//Get data from the artist API
	apiURL := "https://groupietrackers.herokuapp.com/api/artists"
	groupData, err := data.FetchData(apiURL)
	if err != nil {
		fmt.Println("Erreur lors de la récupération des données de l'API:", err)
		return
	}

	//Get data from locations API
	apiURL2 := "https://groupietrackers.herokuapp.com/api/locations"
	groupDataLocations, err := data.FetchDataL(apiURL2)
	if err != nil {
		fmt.Println("Erreur lors de la récupération des données de l'API:", err)
		return
	}

	apiURL3 := "https://groupietrackers.herokuapp.com/api/dates"
	groupDataDates, err := data.FetchDataD(apiURL3)
	if err != nil {
		fmt.Println("Erreur lors de la récupération des données de l'API:", err)
		return
	}

	apiURL4 := "https://groupietrackers.herokuapp.com/api/relation"
	relations, err := data.FetchDataR(apiURL4)
	if err != nil {
		fmt.Println("Erreur lors de la récupération des données de l'API:", err)
		return
	}

	//Store only the locations
	var locations [][]string
	

	for _, lct := range groupDataLocations.Index {
		locations = append(locations, lct.Locations)
	}

	//Store only the locations
	var dates [][]string

	//stores only the geolocalisation

	for _, lct := range groupDataDates.Index {
		dates = append(dates, lct.Dates)
	}
	var stringname []string
	var stringlocation [][]string
	var stringdate [][]string
	for i, group := range groupData {

		fmt.Printf("SUUUUUUUUUUUUUUUUUU: %s\n", relations.Index[i])

		stringname = append(stringname, group.Name)
		stringlocation = append(stringlocation, locations[i])
		stringdate = append(stringdate, dates[i])

	}

	a := app.New()
	gui.W = a.NewWindow("jogoat + samgod + matthis")

	slider := widget.NewSlider(1900, 2024)

	// Étiquette pour afficher la valeur actuelle du slider
	valueLabel := widget.NewLabel(fmt.Sprintf("Year : %d", int(slider.Value)))

	gui.W.SetMainMenu(gui.MakeMenu(a))

	stringList := gui.MakeStringList(stringname, groupData)

	gui.Listcard = container.NewVBox()

	gui.Def = container.NewVBox()

	gui.MakeListCard(gui.Card, gui.Infoback, gui.Listcard, gui.Def, groupData)

	gui.UpperUI = gui.MakeUpperUI(stringList, stringname, groupData, valueLabel, slider, groupDataDates, stringdate)

	cardscroll := container.NewScroll(gui.Listcard)

	cardscroll.SetMinSize(fyne.NewSize(675, 675))
	gui.UpperUI.Add(cardscroll)

	gui.W.Resize(fyne.NewSize(800, 600))

	gui.Window = container.NewVBox(gui.UpperUI)

	gui.W.SetContent(gui.Window)

	gui.W.ShowAndRun()

}
