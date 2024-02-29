package main

import (
	"encoding/json"
	"fmt"
	"image/color"
	"log"
	"net/http"
	"strconv"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

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
		DatesLocations map[string]interface{} `json:"datesLocations"`
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
	/* 	dateLocationMap := make(map[string][]string)
	   	dateLocationMap["2024-02-04"] = append(dateLocationMap["2024-02-04"], "Location1") */

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

	/* 	apiURL4 := "https://groupietrackers.herokuapp.com/api/relation"
	   	groupDataRelations, err := fetchDataR(apiURL4)
	   	if err != nil {
	   		fmt.Println("Erreur lors de la récupération des données de l'API:", err)
	   		return
	   	} */

	relation, er := http.Get("https://groupietrackers.herokuapp.com/api/relation")
	if er != nil {
		log.Fatalf("Erreur lors de la lecture du JSON 1: %s", er)
	}
	defer relation.Body.Close()

	var relations RelationData
	if er := json.NewDecoder(relation.Body).Decode(&relations); er != nil {
		log.Fatalf("Erreur lors de la lecture du JSON 2: %s", er)
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
	var stringlocation [][]string
	var stringdate [][]string
	for i, group := range groupData {
		fmt.Printf("ID: %d\n", group.ID)
		fmt.Printf("Image: %s\n", group.Image)
		fmt.Printf("Name: %s\n", group.Name)
		fmt.Printf("Members: %v\n", group.Members)
		fmt.Printf("Creation Date: %d\n", group.CreationDate)
		fmt.Printf("First Album: %s\n", group.FirstAlbum)
		/* 	fmt.Printf("Locations: %s\n", locations[i])
		fmt.Printf("Concert Dates: %s\n", dates[i])
		fmt.Printf("Relations: %s\n", group.Relations) */

		fmt.Printf("SUUUUUUUUUUUUUUUUUU: %s\n", relations.Index[i])

		stringname = append(stringname, group.Name)
		stringlocation = append(stringlocation, locations[i])
		stringdate = append(stringdate, dates[i])

	}

	a := app.New()
	w := a.NewWindow("Hello")

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

	search := widget.NewEntry()
	searchButton := widget.NewButton("Rechercher", func() {
		// Vérifier si stringList est nul
		if w.Content == nil {
			return
		}
		// Nouvelle liste pour les résultats de la recherche
		filteredList := []string{}

		// Parcourir la liste d'origine et ajouter les éléments correspondants à la nouvelle liste
		for _, item := range stringname {
			if strings.Contains(strings.ToLower(item), strings.ToLower(search.Text)) {
				filteredList = append(filteredList, item)
			}
		}

		// Mettre à jour la liste avec les résultats de la recherche
		stringList.Length = func() int {
			return len(filteredList)
		}
		stringList.CreateItem = func() fyne.CanvasObject {
			return widget.NewLabel("")
		}
		stringList.UpdateItem = func(index int, item fyne.CanvasObject) {
			item.(*widget.Label).SetText(filteredList[index])
		}
		stringList.Refresh()
	})
	clearButton := widget.NewButtonWithIcon("", theme.DeleteIcon(), func() {
		search.SetText("")

		// Mettre à jour la liste avec les résultats de la recherche
		stringList.Length = func() int {
			return len(stringname)
		}
		stringList.CreateItem = func() fyne.CanvasObject {
			return widget.NewLabel("")
		}
		stringList.UpdateItem = func(index int, item fyne.CanvasObject) {
			item.(*widget.Label).SetText(stringname[index])
		}
		stringList.Refresh()
	})

	card := container.NewVBox()

	for _, group := range groupData {
		var listmember string
		for _, memb := range group.Members {
			listmember += memb
			listmember += "  "

		}
		name := canvas.NewText(group.Name, color.Black)
		album := canvas.NewText(group.FirstAlbum, color.Black)
		creationDate := canvas.NewText(strconv.Itoa(group.CreationDate), color.Black)
		members := canvas.NewText(listmember, color.Black)

		card.Add(
			container.NewVBox(
				name,
				members,
				album,
				creationDate,
			),
		)
	}

	researchbar := container.NewVBox(
		search,
		searchButton,
		clearButton,
	)

	cardscroll := container.NewScroll(card)

	cardscroll.SetMinSize(fyne.NewSize(675, 675))
	researchbar.Add(cardscroll)
	w.Resize(fyne.NewSize(800, 600))

	w.SetContent(researchbar)

	w.ShowAndRun()
	//Print the datas

}
