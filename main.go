package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/container"
	"fyne.io/fyne/widget"
	/* 	"fyne.io/fyne/app"
	   	"fyne.io/fyne/widget" */)

type GroupData struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}

type LocationData struct {
	Index []struct {
		ID        int      `json:"id"`
		Locations []string `json:"locations"`
		Dates     string   `json:"dates"`
	} `json:"index"`
}

type DatesData struct {
	Index []struct {
		ID    int      `json:"id"`
		Dates []string `json:"dates"`
	} `json:"index"`
}

type RelationData struct {
	Index []struct {
		ID    int      `json:"id"`
		Dates []string `json:"datesLocations"`
	} `json:"index"`
}

func fetchData(apiURL string) ([]GroupData, error) {
	response, err := http.Get(apiURL)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var data []GroupData
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func fetchDataL(apiURL string) (LocationData, error) {
	var data LocationData

	response, err := http.Get(apiURL)
	if err != nil {
		return data, err
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return data, err
	}

	return data, nil
}

func fetchDataD(apiURL string) (DatesData, error) {
	var data DatesData

	response, err := http.Get(apiURL)
	if err != nil {
		return data, err
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return data, err
	}

	return data, nil
}

func main() {
	dateLocationMap := make(map[string][]string)
	dateLocationMap["2024-02-04"] = append(dateLocationMap["2024-02-04"], "Location1")

	//Get data from the artist API
	apiURL := "https://groupietrackers.herokuapp.com/api/artists"
	groupData, err := fetchData(apiURL)
	if err != nil {
		fmt.Println("Erreur lors de la récupération des données de l'API:", err)
		return
	}

	//Get data from locations API
	apiURL2 := "https://groupietrackers.herokuapp.com/api/locations"
	groupDataLocations, err := fetchDataL(apiURL2)
	if err != nil {
		fmt.Println("Erreur lors de la récupération des données de l'API:", err)
		return
	}

	apiURL3 := "https://groupietrackers.herokuapp.com/api/dates"
	groupDataDates, err := fetchDataD(apiURL3)
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

	for _, lct := range groupDataDates.Index {
		dates = append(dates, lct.Dates)
	}
	var stringname []string
	for i, group := range groupData {
		fmt.Printf("ID: %d\n", group.ID)
		fmt.Printf("Image: %s\n", group.Image)
		fmt.Printf("Name: %s\n", group.Name)
		fmt.Printf("Members: %v\n", group.Members)
		fmt.Printf("Creation Date: %d\n", group.CreationDate)
		fmt.Printf("First Album: %s\n", group.FirstAlbum)
		fmt.Printf("Locations: %s\n", locations[i])
		fmt.Printf("Concert Dates: %s\n", dates[i])
		fmt.Printf("Relations: %s\n", group.Relations)
		stringname = append(stringname, group.Name)

	}
	a := app.New()
	w := a.NewWindow("Hello")

	hello := widget.NewLabel("Hello Fyne!")
	w.SetContent(container.NewVBox(
		hello,
		widget.NewButton("Hi!", func() {
			stringList := widget.NewList(
				func() int {
					return len(stringname)
				},
				func() fyne.CanvasObject {
					return widget.NewLabel("")
				},
				func(i widget.ListItemID, obj fyne.CanvasObject) {
					label := obj.(*widget.Label)
					label.SetText(stringname[i])
					label.Resize(label.MinSize().Add(fyne.NewSize(5000, 5000))) // Largeur de 100 pixels
				},
			)

			w.SetContent(container.NewVBox(
				widget.NewLabel("Tableau de Strings"),
				stringList,
			))

		}),
	))

	w.ShowAndRun()
	//Print the datas

}
