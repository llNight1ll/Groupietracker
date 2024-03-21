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
	backButton := widget.NewButton("Retour", func() {
		w.SetContent(searchContainer) // Revenir à la liste de recherche
	})

	for _, group := range groupData {
		if group.ID == groupID {
			// Créez un widget de texte pour afficher les détails du groupe dans la fenêtre
			artist := widget.NewLabel(group.Name)
			members := widget.NewLabel(strings.Join(group.Members, ", ")) // Convertir le slice en chaîne de caractères
			album := widget.NewLabel(group.FirstAlbum)
			creationDate := widget.NewLabel(fmt.Sprintf("%d", group.CreationDate))
			imageURL := group.Image

			r, _ := fyne.LoadResourceFromURLString(imageURL)
			img := canvas.NewImageFromResource(r)
			img.FillMode = canvas.ImageFillContain // Gestion du fill image
			img.SetMinSize(fyne.NewSize(120, 120)) //Définir la taille minimum de l'image
			img.Resize(fyne.NewSize(120, 120))

			// Créez un conteneur pour afficher les détails du groupe
			groupDetails := container.NewVBox(
				img,
				artist,
				members,
				album,
				creationDate,
				backButton,
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
		// fyne.NewMenu("Quitter"),

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
			return len(groupData)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("")
		},
		func(i widget.ListItemID, obj fyne.CanvasObject) {
			label := obj.(*widget.Label)
			label.SetText(fmt.Sprintf("%s %d", groupData[i].Name, groupData[i].ID))
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

		l, _ := fyne.LoadResourceFromURLString(imageURL)
		img := canvas.NewImageFromResource(l)
		img.FillMode = canvas.ImageFillContain // Gestion du fill image
		img.SetMinSize(fyne.NewSize(120, 120)) //Définir la taille minimum de l'image
		img.Resize(fyne.NewSize(120, 120))

		r, g, b, a := calculateAverageColor(img)

		background := canvas.NewRectangle(color.RGBA{uint8(r), uint8(g), uint8(b), uint8(a)})
		background.FillColor = color.RGBA{uint8(r), uint8(g), uint8(b), uint8(a)}

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
		// Désactiver barre de recherche
		search.Disable()
	
		searchText := strings.ToLower(search.Text)
		suggestions := make([]fyne.CanvasObject, 0)
	
		for _, group := range groupData {
			imageURL := group.Image
	
			l, _ := fyne.LoadResourceFromURLString(imageURL)
			img := canvas.NewImageFromResource(l)
			img.FillMode = canvas.ImageFillContain // Gestion du fill image
			img.SetMinSize(fyne.NewSize(120, 120)) // Définir la taille minimum de l'image
			img.Resize(fyne.NewSize(120, 120))
	
			// Créer un bouton personnalisé avec l'image et le nom du groupe
			suggestion := widget.NewButton("", func(groupID int) func() {
				return func() {
					showGroupDetails(groupID, groupData, w, searchContainer, stringList) // Passer searchContainer à la fonction
				}
			}(group.ID))
			suggestion.Importance = widget.LowImportance // Réduire l'importance pour que cela ne ressemble pas à un bouton standard
			suggestion.SetIcon(l)
			suggestion.Resize(fyne.NewSize(200, 200)) // Définir l'image comme icône du bouton
			suggestion.SetText(group.Name)            // Définir le nom du groupe comme texte du bouton
	
			// Ajouter le bouton à la liste des suggestions
			if strings.Contains(strings.ToLower(group.Name), searchText) {
				suggestions = append(suggestions, suggestion)
			}
		}
	
		if len(suggestions) > 0 {
			a := container.NewVBox(search, searchButton, clearButton)
			suggestionsContainer := container.NewVBox(suggestions...)
			b := container.NewVBox(a, suggestionsContainer)
			w.SetContent(container.NewVScroll(b))
		} else {
			// Afficher un message si aucune suggestion n'est trouvée
			w.SetContent(widget.NewLabel("Aucun groupe trouvé avec ce nom."))
		}

		//reactiver la barre de recherche
		search.Enable()
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

	w.SetContent(container.NewVBox(searchContainer))
	w.SetContent(researchbar)

	w.ShowAndRun()
}

func calculateAverageColor(img *canvas.Image) (r uint32, g uint32, b uint32, a uint32) {

	width := int(img.MinSize().Width)
	height := int(img.MinSize().Height)
	var rMoyenne uint32
	var gMoyenne uint32
	var bMoyenne uint32
	var aMoyenne uint32
	var j uint32

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {

			colorrr := img.Image.At(x, y)
			r, g, b, a := colorToRGB(colorrr)
			rMoyenne += r
			gMoyenne += g
			bMoyenne += b
			aMoyenne += a
			j++

		}
	}
	rMoyenne = rMoyenne / j
	gMoyenne = gMoyenne / j
	bMoyenne = bMoyenne / j
	aMoyenne = aMoyenne / j

	return rMoyenne, gMoyenne, bMoyenne, aMoyenne

}

func colorToRGB(c color.Color) (r, g, b, a uint32) {
	r, g, b, a = c.RGBA()
	r = r / 257
	g = g / 257
	b = b / 257
	a = a / 257
	return r, g, b, a
}