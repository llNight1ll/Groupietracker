package main

import (
	"encoding/json"
	"fmt"
	"image/color"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
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

func showGroupDetails(groupID int, groupData []GroupData, w fyne.Window, searchContainer *fyne.Container, stringList fyne.CanvasObject) {
	// backButton := widget.NewButton("Retour", func() {
	// 	w.SetContent(searchContainer) // Revenir à la liste de recherche
	// })

	for _, group := range groupData {
		if group.ID == groupID {
			// Créez un widget de texte pour afficher les détails du groupe dans la fenêtre
			artist := widget.NewLabel(group.Name)
			album := widget.NewLabel(group.FirstAlbum)

			// Créez un conteneur pour afficher les détails du groupe
			groupDetails := container.NewVBox(
				artist,
				album,
				// backButton,
			)

			// Placez les détails du groupe au centre de la fenêtre
			content := container.NewBorder(nil, nil, nil, nil, groupDetails)
			w.SetContent(content)
			return
		}
	}
}

func main() {
	searchContainer := container.NewVBox()
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

	menu := fyne.NewMainMenu(
		fyne.NewMenu("Quitter"),

		// Theme de le la page
		fyne.NewMenu("Thèmes",
			fyne.NewMenuItem("Thèmes sombre", func() {
				a.Settings().SetTheme(theme.DarkTheme())
			}),

			fyne.NewMenuItem("Thème clair", func() {
				a.Settings().SetTheme(theme.LightTheme())
			}),
		),

		fyne.NewMenu("En savoir plus",
			fyne.NewMenuItem("Spotify", func() {
				lien, _ := url.Parse("https://developer.spotify.com/documentation/embeds")
				_ = a.OpenURL(lien)
			}),
		))

	w.SetMainMenu(menu)

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
		imageURL := group.Image

		r, _ := fyne.LoadResourceFromURLString(imageURL)
		img := canvas.NewImageFromResource(r)
		img.FillMode = canvas.ImageFillContain // Gestion du fill image
		img.SetMinSize(fyne.NewSize(120, 120)) //Définir la taille minimum de l'image
		img.Resize(fyne.NewSize(120, 120))

		background := canvas.NewRectangle(color.RGBA{255, 0, 0, 255})
		background.FillColor = color.RGBA{0, 0, 255, 255}

		info := container.New(layout.NewVBoxLayout(),

			img,
			container.NewCenter(name),
			container.NewCenter(members),
			container.NewCenter(album),
			container.NewCenter(creationDate),
		)

		card.Add(
			container.New(layout.NewBorderLayout(nil, nil, nil, nil), background, info),
		)

	}

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

	stringList.OnSelected = func(id widget.ListItemID) {
		groupID := groupData[id].ID
		showGroupDetails(groupID, groupData, w, searchContainer, stringList) // Passer la liste de recherche et la barre de recherche à la fonction
	}

	searchButton.OnTapped = func() {
		searchText := strings.ToLower(search.Text)
		suggestions := make([]fyne.CanvasObject, 0)

		for _, group := range groupData {
			if strings.Contains(strings.ToLower(group.Name), searchText) {
				suggestion := widget.NewButton(group.Name, func(groupID int) func() {
					return func() {
						showGroupDetails(groupID, groupData, w, searchContainer, stringList) // Passer searchContainer à la fonction
					}
				}(group.ID))
				suggestions = append(suggestions, suggestion)
			}
		}

		if len(suggestions) > 0 {
			suggestionsContainer := container.NewVBox(suggestions...)
			content := container.NewBorder(nil, nil, nil, nil, search)
			content.Add(container.NewVScroll(suggestionsContainer))
			w.SetContent(content)
		} else {
			w.SetContent(widget.NewLabel("Aucun groupe trouvé avec ce nom."))
		}
	}

	cardscroll := container.NewScroll(card)
	cardscroll.SetMinSize(fyne.NewSize(675, 675))

	researchbar := container.NewVBox(
		search,
		searchButton,
		clearButton,
	)

	researchbar.Add(cardscroll)
	w.Resize(fyne.NewSize(800, 600))

	content := container.NewVSplit(
		researchbar,
		stringList,
	)

	w.SetContent(content)

	w.ShowAndRun()
	//Print the datas

}

func calculateAverageColor(img *canvas.Image) {
	// Obtention des dimensions de l'image
	width := int(img.MinSize().Width)
	height := int(img.MinSize().Height)

	// Initialisation des valeurs de couleur moyenne

	// Accès aux pixels de l'image

	// Parcours de tous les pixels de l'image
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			// Récupération de la couleur du pixel
			colorrr := img.Image.At(x, y)
			colorToRGB(colorrr)
		}
	}

	// Calcul des valeurs moyennes des composantes de couleur

	// Retour de la couleur moyenne
}

func colorToRGB(c color.Color) (r, g, b, a uint8) {
	switch c.(type) {
	case color.RGBA:
		rgba := c.(color.RGBA)
		r, g, b, a = rgba.R, rgba.G, rgba.B, rgba.A
	case color.RGBA64:
		rgba64 := c.(color.RGBA64)
		r, g, b, a = uint8(rgba64.R>>8), uint8(rgba64.G>>8), uint8(rgba64.B>>8), uint8(rgba64.A>>8)
	default:
		// Si le type de couleur n'est ni RGBA ni RGBA64, les composantes seront vides
	}
	return r, g, b, a
}
