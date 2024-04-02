package main

import (
	"encoding/json"
	"fmt"
	"image/color"
	"net/http"
	"net/url"

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

var listFavorit []string
var wind222 *fyne.Container
var card *fyne.Container
var infoback *fyne.Container
var listcard *fyne.Container
var def *fyne.Container
var window *fyne.Container
var w fyne.Window
var searchContainer *fyne.Container
var upperUI *fyne.Container

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

func fetchDataR(apiURL string) (RelationData, error) {
	var relations RelationData

	response, err := http.Get(apiURL)
	if err != nil {
		return relations, err
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&relations)
	if err != nil {
		return relations, err
	}
	return relations, nil

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

func showGroupDetails2(groupID int, groupData []GroupData, w fyne.Window, searchContainer *fyne.Container, window *fyne.Container) {
	backButton := widget.NewButton("Retour", func() {

		w.SetContent(window)
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

func showGroupDetails(groupID int, groupData []GroupData, w fyne.Window, searchContainer *fyne.Container) {
	backButton := widget.NewButton("Retour", func() {
		w.SetContent(container.NewVScroll(wind222)) // Revenir à la liste de recherche
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

func makeMenu(a fyne.App) *fyne.MainMenu {
	menu := fyne.NewMainMenu(

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
		),
	)
	return menu
}

func makeListCard(card *fyne.Container, infoback *fyne.Container, listcard *fyne.Container, def *fyne.Container, groupData []GroupData) {
	for _, group := range groupData {
		var listmember string
		for _, memb := range group.Members {
			listmember += memb
			listmember += "  "

		}
		name := canvas.NewText(group.Name, color.Black)
		members := canvas.NewText(listmember, color.Black)
		imageURL := group.Image
		heartOnImage, _ := fyne.LoadResourceFromPath("./heartOn.png")

		heartOffImage, _ := fyne.LoadResourceFromPath("./heartOff.png")

		// Créer un booléen pour suivre l'état du bouton
		var isPressed bool
		var heartButton *widget.Button
		// Créer un bouton avec l'image initiale du cœur

		heartButton = widget.NewButton("", func() {

			// Inverser l'état lors du clic sur le bouton
			isPressed = !isPressed
			// Mettre à jour l'image du bouton en fonction de l'état
			if isPressed {
				heartButton.SetIcon(heartOnImage)
				listFavorit = append(listFavorit, group.Name)
			} else {
				heartButton.SetIcon(heartOffImage)
				for i, name := range listFavorit {
					if name == group.Name {
						listFavorit = append(listFavorit[:i], listFavorit[i+1:]...)

					}
				}
			}

		})
		heartButton.SetIcon(heartOffImage)
		heartButton.Importance = widget.LowImportance

		viewDetail := widget.NewButton("View Detail", func() {

			showGroupDetails2(group.ID, groupData, w, searchContainer, window)

		})

		l, _ := fyne.LoadResourceFromURLString(imageURL)
		img := canvas.NewImageFromResource(l)
		img.FillMode = canvas.ImageFillContain // Gestion du fill image
		img.SetMinSize(fyne.NewSize(200, 200)) //Définir la taille minimum de l'image
		img.Resize(fyne.NewSize(200, 200))

		r, g, b, a := calculateAverageColor(img)

		background := canvas.NewRectangle(color.RGBA{uint8(r), uint8(g), uint8(b), uint8(a)})
		background.FillColor = color.RGBA{uint8(r), uint8(g), uint8(b), uint8(a)}

		background.SetMinSize(fyne.NewSize(300, 300)) // Définir la taille minimum du bakcground
		background.Resize(fyne.NewSize(296, 296))     // Redimensionner pour inclure les coin
		background.CornerRadius = 20                  // Définir les coins arrondis

		background2 := canvas.NewRectangle(color.RGBA{255, 255, 255, 255})
		background2.FillColor = color.RGBA{255, 255, 255, 255}

		background2.SetMinSize(fyne.NewSize(100, 100)) // Définir la taille minimum du bakcground
		background2.Resize(fyne.NewSize(100, 100))     // Redimensionner pour inclure les coin
		background2.CornerRadius = 20                  // Définir les coins arrondis

		iinfo := container.New(layout.NewVBoxLayout(),
			container.NewCenter(name),
			container.NewCenter(members),
			heartButton,
			viewDetail)

		infoback = container.New(layout.NewBorderLayout(nil, nil, nil, nil), background2, iinfo)

		info := container.New(layout.NewVBoxLayout(),

			img,
			container.NewCenter(infoback),
		)

		card = container.New(layout.NewBorderLayout(nil, nil, nil, nil), background, info)

		card.Resize(fyne.NewSize(300, 300)) //Définir la taille minimum de la card

		border := canvas.NewRectangle(color.Transparent) // Définir une couleur transparente pour le remplissage
		border.SetMinSize(fyne.NewSize(300, 300))        //Définir la taille minimum de la bordure
		border.Resize(fyne.NewSize(296, 296))            // Redimensionner pour inclure les coin
		border.StrokeColor = color.Black                 // Définir la couleur de la bordure
		border.StrokeWidth = 3                           // Définir l'épaisseur de la bordure
		border.CornerRadius = 20                         // Définir les coins

		card.Add(border)

		card.Resize(fyne.NewSize(100, 300))

		listcard.Add(card)
		def.Add(card)

	}

}

func makeStringList(stringname []string, groupData []GroupData) *widget.List {

	stringList := widget.NewList(
		func() int {
			return len(stringname)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("")
		},
		func(i widget.ListItemID, obj fyne.CanvasObject) {
			label := obj.(*widget.Label)
			label.SetText(fmt.Sprintf("%s %d", groupData[i].Name, groupData[i].ID))
			label.SetText(stringname[i])
			label.Resize(label.MinSize().Add(fyne.NewSize(100, 100))) // Largeur de 100 pixels
		},
	)
	return stringList
}

// Fonction pour mettre à jour les résultats en fonction des cases à cocher sélectionnées
func updateFilteredResults(w fyne.Window, groupData []GroupData, searchContainer *fyne.Container, checkboxFilters []*widget.Check) {
	selectedMembers := make(map[int]bool)
	for i, checkbox := range checkboxFilters {
		if checkbox.Checked {
			selectedMembers[i+2] = true
		}
	}

	// Filtrer les résultats en fonction des membres sélectionnés
	filteredResults := make([]GroupData, 0)
	for _, group := range groupData {
		if len(selectedMembers) == 0 {
			// Aucun filtre sélectionné, ajoutez simplement tous les résultats
			filteredResults = append(filteredResults, group)
		} else {
			// Vérifiez si le groupe a le nombre de membres sélectionné
			if selectedMembers[len(group.Members)] {
				filteredResults = append(filteredResults, group)
			}
		}
	}

	// Mettre à jour le contenu de la fenêtre avec les nouveaux résultats filtrés
	updateWindowContent(w, groupData, filteredResults, searchContainer)
}

// // Fonction pour mettre à jour le contenu de la fenêtre avec les résultats filtrés
func updateWindowContent(w fyne.Window, groupData []GroupData, results []GroupData, searchContainer *fyne.Container) {
	// Créer une nouvelle liste de boutons pour afficher les résultats
	resultButtons := make([]fyne.CanvasObject, len(results))

	// Parcourir les résultats et créer un bouton pour chaque élément
	for i, result := range results {
		// Créer un bouton personnalisé avec le nom du groupe
		button := widget.NewButton(result.Name, func(groupID int) func() {
			return func() {
				// Action à effectuer lorsque le bouton est cliqué
				showGroupDetails(groupID, groupData, w, searchContainer)
			}
		}(result.ID))
		button.Importance = widget.LowImportance // Réduire l'importance pour que cela ne ressemble pas à un bouton standard
		resultButtons[i] = button
	}

	// Créer un conteneur de type VBox pour afficher les boutons
	content := container.NewVBox(resultButtons...)

	// Mettre à jour le contenu de la fenêtre avec les résultats de la recherche
	w.SetContent(content)
}

func makeUpperUI(stringList *widget.List, stringname []string, groupData []GroupData, valueLabel *widget.Label, slider *widget.Slider, groupDataDates DatesData, stringdate [][]string) *fyne.Container {
	
	var checkboxFilters []*widget.Check

	// Créer les cases à cocher pour les filtres
	for i := 2; i <= 7; i++ {
		checkbox := widget.NewCheck(fmt.Sprintf("%d", i), func(members int) func(bool) {
			return func(checked bool) {
				// Ajouter votre logique de traitement ici en fonction de l'état de la case à cocher
				fmt.Printf("%d membres: %t\n", members, checked)
			}
		}(i))
		checkboxFilters = append(checkboxFilters, checkbox)
	}

	checkboxContainer := container.NewHBox()
	label := widget.NewLabel("Nombre de membres:")
	checkboxContainer.Add(label)
	for _, checkbox := range checkboxFilters {
		checkboxContainer.Add(checkbox)
	}

	// updateFilteredResults(w, groupData, searchContainer, checkboxFilters)
	
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

		for _, items := range stringdate {
			for _, item := range items {
				if strings.Contains(strings.ToLower(item), strings.ToLower(search.Text)) {
					filteredList = append(filteredList, item)
				}
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

		stringList.UpdateItem = func(index int, items fyne.CanvasObject) {
			items.(*widget.Label).SetText(filteredList[index])
		}
		// updateFilteredResults(w, groupData, searchContainer, checkboxFilters)
		stringList.Refresh()
	})

	stringList.OnSelected = func(id widget.ListItemID) {
		groupID := groupData[id].ID
		showGroupDetails(groupID, groupData, w, searchContainer) // Passer la liste de recherche et la barre de recherche à la fonction
	}

	sugg := container.NewVBox()
	sugg2 := container.NewVScroll(sugg)
	search.OnChanged = func(query string) {
		searchText := strings.ToLower(query)
		if len(query) > 0 {
			sugg.Objects = make([]fyne.CanvasObject, 0)

			for _, group := range groupData {
				if strings.Contains(strings.ToLower(group.Name), searchText) {
					label := group.Name + "        - Groupe"
					h := widget.NewButton(label, func(groupID int) func() {
						return func() {
							showGroupDetails2(groupID, groupData, w, searchContainer, window) // Passer searchContainer à la fonction
						}
					}(group.ID))
					h.Importance = widget.LowImportance
					sugg.Add(h)

				}
				for _, groupMember := range group.Members {
					if strings.Contains(strings.ToLower(groupMember), searchText) {
						label2 := groupMember + "         - Member"
						h := widget.NewButton(label2, func(groupID int) func() {
							return func() {
								showGroupDetails2(groupID, groupData, w, searchContainer, window) // Passer searchContainer à la fonction
							}

						}(group.ID))
						h.Importance = widget.LowImportance

						sugg.Add(h)

					}
				}

			}
			sugg.Show()
			sugg2.Show()
		} else {
			sugg.Hide()
			sugg2.Hide()
		}

	}
	favoris := widget.NewButton("Favoris", func() {
		var fg []string
		var favoritCard *fyne.Container
		favoritMenu := container.NewVBox()
		for _, favorisGroup := range listFavorit {

			for _, group := range groupData {
				if favorisGroup == group.Name {
					var listmember string
					for _, memb := range group.Members {
						listmember += memb
						listmember += "  "

					}
					name := canvas.NewText(group.Name, color.Black)
					members := canvas.NewText(listmember, color.Black)
					imageURL := group.Image

					l, _ := fyne.LoadResourceFromURLString(imageURL)
					img := canvas.NewImageFromResource(l)
					img.FillMode = canvas.ImageFillContain // Gestion du fill image
					img.SetMinSize(fyne.NewSize(200, 200)) //Définir la taille minimum de l'image
					img.Resize(fyne.NewSize(200, 200))

					r, g, b, a := calculateAverageColor(img)

					background := canvas.NewRectangle(color.RGBA{uint8(r), uint8(g), uint8(b), uint8(a)})
					background.FillColor = color.RGBA{uint8(r), uint8(g), uint8(b), uint8(a)}

					background.SetMinSize(fyne.NewSize(300, 300)) // Définir la taille minimum du bakcground
					background.Resize(fyne.NewSize(296, 296))     // Redimensionner pour inclure les coin
					background.CornerRadius = 20                  // Définir les coins arrondis

					background2 := canvas.NewRectangle(color.RGBA{255, 255, 255, 255})
					background2.FillColor = color.RGBA{255, 255, 255, 255}

					background2.SetMinSize(fyne.NewSize(100, 100)) // Définir la taille minimum du bakcground
					background2.Resize(fyne.NewSize(100, 100))     // Redimensionner pour inclure les coin
					background2.CornerRadius = 20                  // Définir les coins arrondis

					iinfo := container.New(layout.NewVBoxLayout(),
						container.NewCenter(name),
						container.NewCenter(members))

					infoback = container.New(layout.NewBorderLayout(nil, nil, nil, nil), background2, iinfo)

					info := container.New(layout.NewVBoxLayout(),

						img,
						container.NewCenter(infoback),
					)

					favoritCard = container.New(layout.NewBorderLayout(nil, nil, nil, nil), background, info)

					favoritCard.Resize(fyne.NewSize(300, 300)) //Définir la taille minimum de la card

					border := canvas.NewRectangle(color.Transparent) // Définir une couleur transparente pour le remplissage
					border.SetMinSize(fyne.NewSize(300, 300))        //Définir la taille minimum de la bordure
					border.Resize(fyne.NewSize(296, 296))            // Redimensionner pour inclure les coin
					border.StrokeColor = color.Black                 // Définir la couleur de la bordure
					border.StrokeWidth = 3                           // Définir l'épaisseur de la bordure
					border.CornerRadius = 20                         // Définir les coins

					favoritCard.Add(border)

					favoritCard.Resize(fyne.NewSize(100, 300))

					favoritMenu.Add(favoritCard)
					fg = append(fg, group.Name)

				}
			}

		}
		fmt.Println(fg)
		backButton := widget.NewButton("Retour", func() {

			w.SetContent(window)
		})

		favoritCardScroll := container.NewScroll(favoritMenu)
		favoritCardScroll.SetMinSize(fyne.NewSize(675, 675))

		favoritPage := container.NewVBox(
			backButton,
			favoritCardScroll)

		w.SetContent(favoritPage)

	})
	var clearButton *widget.Button

	clearButton = widget.NewButtonWithIcon("", theme.DeleteIcon(), func() {
		search.SetText("")

		// Réinitialiser les filtres sélectionnés
		for _, checkbox := range checkboxFilters {
			checkbox.Checked = false
		}

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

		slider.SetValue(0)

		cardscroll := container.NewScroll(listcard)

		cardscroll.SetMinSize(fyne.NewSize(675, 675))

		sugg2 = container.NewVScroll(sugg)

		sugg2.SetMinSize(fyne.NewSize(100, 100))
		sugg2.Hide()
		spacer := layout.NewSpacer()
		sugg3 := container.NewHBox(sugg2, spacer)
		spacer.Resize(fyne.NewSize(100, 200))

		researchbar := container.NewVBox(
			checkboxContainer,
			favoris,
			search,
			sugg3,
			searchButton,
			clearButton,
			valueLabel,
			slider,
		)

		researchbar.Add(cardscroll)
		w.Resize(fyne.NewSize(800, 600))

		window = container.NewVBox(researchbar)

		w.SetContent(window)

		//w.SetContent(container.NewVBox(searchContainer))

	})

	searchButton.OnTapped = func() {
		// Désactiver barre de recherche
		search.Disable()

		searchText := strings.ToLower(search.Text)
		suggestions := make([]fyne.CanvasObject, 0)

		
		// updateFilteredResults(w, groupData, searchContainer, checkboxFilters)

		verif := false
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
					showGroupDetails(groupID, groupData, w, searchContainer) // Passer searchContainer à la fonction
				}
			}(group.ID))
			suggestion.Importance = widget.LowImportance // Réduire l'importance pour que cela ne ressemble pas à un bouton standard
			suggestion.SetIcon(l)
			suggestion.Resize(fyne.NewSize(200, 200)) // Définir l'image comme icône du bouton
			suggestion.SetText(group.Name)            // Définir le nom du groupe comme texte du bouton

			// Ajouter le bouton à la liste des suggestions
			if strings.Contains(strings.ToLower(group.Name), searchText) {
				suggestions = append(suggestions, suggestion)
				verif = true
			} else {
				for _, member := range group.Members {
					if strings.Contains(strings.ToLower(member), searchText) {
						suggestions = append(suggestions, suggestion)
						verif = true
						break
					}
				}
			}

			if strings.Contains(fmt.Sprintf("%d", group.CreationDate), searchText) {
				suggestions = append(suggestions, suggestion)
				verif = true
			}

			// if strings.Contains(strings.ToLower(group.FirstAlbum), searchText) {
			// 	suggestions = append(suggestions, suggestion)
			// 	verif = true
			// }

			for _, date := range groupDataDates.Index {
				if strings.Contains(strings.ToLower(date.Dates[0]), searchText) {
					suggestions = append(suggestions, suggestion)
					verif = true
				}
			}
		}

		//Afficher un message si la date et l'annee ne correspond à aucun artiste
		if !verif {
			r := container.NewVBox(widget.NewLabel("Aucun groupe trouvé avec ce nom."))
			r.Add(clearButton)
			w.SetContent(r)
			search.Enable()

			return
		}

		if len(suggestions) > 0 {
			spacer := layout.NewSpacer()
			sugg3 := container.NewHBox(sugg2, spacer)
			spacer.Resize(fyne.NewSize(100, 200))
			rsrch := container.NewVBox(search, sugg3, searchButton, clearButton)
			suggestionsContainer := container.NewVBox(suggestions...)
			wind222 = container.NewVBox(rsrch, suggestionsContainer)
			w.SetContent(container.NewVScroll(wind222))
		} else {
			// Afficher un message si aucune suggestion n'est trouvée
			r := container.NewVBox(widget.NewLabel("Aucun groupe trouvé avec ce nom."))
			r.Add(clearButton)

			w.SetContent(r)
		}

		//reactiver la barre de recherche
		search.Enable()
	}

	search.OnSubmitted = func(text string) {
		// Lancer la recherche lorsque la touche "Entrer" est pressée
		searchButton.OnTapped()
	}
	var deft bool
	// Gérer le changement de valeur du slider
	slider.OnChanged = func(value float64) {
		deft = true
		listcard.RemoveAll()
		valueLabel.SetText(fmt.Sprintf("Year : %d", int(slider.Value)))
		for _, group := range groupData {
			if slider.Value == float64(group.CreationDate) {
				var listmember string
				for _, memb := range group.Members {
					listmember += memb
					listmember += "  "

				}
				name := canvas.NewText(group.Name, color.Black)
				members := canvas.NewText(listmember, color.Black)
				imageURL := group.Image

				l, _ := fyne.LoadResourceFromURLString(imageURL)
				img := canvas.NewImageFromResource(l)
				img.FillMode = canvas.ImageFillContain // Gestion du fill image
				img.SetMinSize(fyne.NewSize(200, 200)) //Définir la taille minimum de l'image
				img.Resize(fyne.NewSize(200, 200))

				r, g, b, a := calculateAverageColor(img)

				background := canvas.NewRectangle(color.RGBA{uint8(r), uint8(g), uint8(b), uint8(a)})
				background.FillColor = color.RGBA{uint8(r), uint8(g), uint8(b), uint8(a)}

				background.SetMinSize(fyne.NewSize(300, 300)) // Définir la taille minimum du bakcground
				background.Resize(fyne.NewSize(296, 296))     // Redimensionner pour inclure les coin
				background.CornerRadius = 20                  // Définir les coins arrondis

				background2 := canvas.NewRectangle(color.RGBA{255, 255, 255, 255})
				background2.FillColor = color.RGBA{255, 255, 255, 255}

				background2.SetMinSize(fyne.NewSize(100, 100)) // Définir la taille minimum du bakcground
				background2.Resize(fyne.NewSize(100, 100))     // Redimensionner pour inclure les coin
				background2.CornerRadius = 20                  // Définir les coins arrondis

				iinfo := container.New(layout.NewVBoxLayout(),
					container.NewCenter(name),
					container.NewCenter(members))

				infoback = container.New(layout.NewBorderLayout(nil, nil, nil, nil), background2, iinfo)

				info := container.New(layout.NewVBoxLayout(),

					img,
					container.NewCenter(infoback),
				)

				card = container.New(layout.NewBorderLayout(nil, nil, nil, nil), background, info)

				card.Resize(fyne.NewSize(300, 300)) //Définir la taille minimum de la card

				border := canvas.NewRectangle(color.Transparent) // Définir une couleur transparente pour le remplissage
				border.SetMinSize(fyne.NewSize(300, 300))        //Définir la taille minimum de la bordure
				border.Resize(fyne.NewSize(296, 296))            // Redimensionner pour inclure les coin
				border.StrokeColor = color.Black                 // Définir la couleur de la bordure
				border.StrokeWidth = 3                           // Définir l'épaisseur de la bordure
				border.CornerRadius = 20                         // Définir les coins

				card.Add(border)

				card.Resize(fyne.NewSize(100, 300))

				listcard.Add(card)
				deft = false

			}
		}
		if deft {
			listcard.RemoveAll()
			for _, o := range def.Objects {
				listcard.Add(o)

			}
		}

	}

	sugg2 = container.NewVScroll(sugg)

	sugg2.SetMinSize(fyne.NewSize(100, 100))
	sugg2.Hide()
	spacer := layout.NewSpacer()
	sugg3 := container.NewHBox(sugg2, spacer)
	spacer.Resize(fyne.NewSize(100, 200))

	upperUI := container.NewVBox(
		checkboxContainer,
		favoris,
		search,
		sugg3,
		searchButton,
		clearButton,
		valueLabel,
		slider,
	)
	return upperUI

}

func main() {

	searchContainer = container.NewVBox()
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

	apiURL4 := "https://groupietrackers.herokuapp.com/api/relation"
	relations, err := fetchDataR(apiURL4)
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
	var stringlocation [][]string
	var stringdate [][]string
	for i, group := range groupData {

		fmt.Printf("SUUUUUUUUUUUUUUUUUU: %s\n", relations.Index[i])

		stringname = append(stringname, group.Name)
		stringlocation = append(stringlocation, locations[i])
		stringdate = append(stringdate, dates[i])

	}

	a := app.New()
	w = a.NewWindow("jogoat + samgod")

	slider := widget.NewSlider(1900, 2024)

	// Étiquette pour afficher la valeur actuelle du slider
	valueLabel := widget.NewLabel(fmt.Sprintf("Year : %d", int(slider.Value)))

	w.SetMainMenu(makeMenu(a))

	stringList := makeStringList(stringname, groupData)

	listcard = container.NewVBox()

	def = container.NewVBox()

	makeListCard(card, infoback, listcard, def, groupData)

	upperUI = makeUpperUI(stringList, stringname, groupData, valueLabel, slider, groupDataDates, stringdate)

	cardscroll := container.NewScroll(listcard)

	cardscroll.SetMinSize(fyne.NewSize(675, 675))
	upperUI.Add(cardscroll)

	w.Resize(fyne.NewSize(800, 600))

	window = container.NewVBox(upperUI)

	w.SetContent(window)

	w.ShowAndRun()

}