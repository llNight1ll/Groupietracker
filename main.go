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

	//Get data from the dates API
	apiURL3 := "https://groupietrackers.herokuapp.com/api/dates"
	groupDataDates, err := data.FetchDataD(apiURL3)
	if err != nil {
		fmt.Println("Erreur lors de la récupération des données de l'API:", err)
		return
	}

	//Get data from the relation API
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

	//Store only the dates
	var dates [][]string

	//stores only the geolocalisation

	for _, lct := range groupDataDates.Index {
		dates = append(dates, lct.Dates)
	}
	//Variables wich contain all the group's name, concert locations and dates
	var stringname []string
	var stringlocation [][]string
	var stringdate [][]string
	for i, group := range groupData {

		fmt.Println(relations.Index[i])

		stringname = append(stringname, group.Name)
		stringlocation = append(stringlocation, locations[i])
		stringdate = append(stringdate, dates[i])

	}
	//Create a new app
	a := app.New()

	gui.W = a.NewWindow("MusicData")

	gui.SearchContainer = container.NewVBox()

	//Create a slider from 1900 to 2024
	slider := widget.NewSlider(1900, 2024)

	// Label for the current value of the slider
	valueLabel := widget.NewLabel(fmt.Sprintf("Creation Year : %d", int(slider.Value)))

	//Set the menu
	gui.W.SetMainMenu(gui.MakeMenu(a))

	// List of the result after clicking on the search bar
	stringList := gui.MakeStringList(stringname, groupData)

	// Container which is going to contains the card ( filtered )
	gui.Listcard = container.NewVBox()

	// Container which is going to contains all the card
	gui.Def = container.NewVBox()

	//Make the card list of the home page
	gui.MakeListCard(gui.Card, gui.Infoback, gui.Listcard, gui.Def, groupData)
	//Make the Ui on top of the card list
	gui.UpperUI = gui.MakeUpperUI(stringList, stringname, groupData, valueLabel, slider, groupDataDates, stringdate)

	cardscroll := container.NewScroll(gui.Listcard)

	cardscroll.SetMinSize(fyne.NewSize(675, 675))
	gui.UpperUI.Add(cardscroll)

	gui.W.Resize(fyne.NewSize(800, 600))

	gui.Window = container.NewVBox(gui.UpperUI)

	// Set the home page
	gui.W.SetContent(gui.Window)

	//Run the app
	gui.W.ShowAndRun()

}
