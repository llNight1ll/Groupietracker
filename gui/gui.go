package gui

import (
	"fmt"
	"groupie/colorAnalysis"
	"groupie/structdata"
	"image/color"
	"net/url"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

var ListFavorit []string
var Wind222 *fyne.Container
var Card *fyne.Container
var Infoback *fyne.Container
var Listcard *fyne.Container
var Def *fyne.Container
var Window *fyne.Container
var W fyne.Window
var SearchContainer *fyne.Container
var UpperUI *fyne.Container

func ShowGroupDetails2(groupID int, groupData []structdata.GroupData, W fyne.Window, SearchContainer *fyne.Container, Window *fyne.Container) {
	backButton := widget.NewButton("Retour", func() {

		W.SetContent(Window)
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
			W.SetContent(content)
			return
		}
	}
}

func ShowGroupDetails(groupID int, groupData []structdata.GroupData, W fyne.Window, SearchContainer *fyne.Container) {
	backButton := widget.NewButton("Retour", func() {
		W.SetContent(container.NewVScroll(Wind222)) // Revenir à la liste de recherche
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
			W.SetContent(content)
			return
		}
	}
}

func MakeMenu(a fyne.App) *fyne.MainMenu {
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

func MakeListCard(Card *fyne.Container, Infoback *fyne.Container, Listcard *fyne.Container, Def *fyne.Container, groupData []structdata.GroupData) {
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
		nameGroup := group.Name
		// Créer un bouton avec l'image initiale du cœur

		heartButton = widget.NewButton("", func() {
			nameG := nameGroup

			// Inverser l'état lors du clic sur le bouton
			isPressed = !isPressed
			// Mettre à jour l'image du bouton en fonction de l'état
			if isPressed {
				heartButton.SetIcon(heartOnImage)

				ListFavorit = append(ListFavorit, nameG)
			} else {
				heartButton.SetIcon(heartOffImage)
				for i, nam := range ListFavorit {
					if nam == nameG {
						ListFavorit = append(ListFavorit[:i], ListFavorit[i+1:]...)

					}
				}
			}
			fmt.Println(ListFavorit)

		})
		heartButton.SetIcon(heartOffImage)
		heartButton.Importance = widget.LowImportance

		groupID := group.ID

		viewDetail := widget.NewButton("View Detail", func() {
			GrpID := groupID

			ShowGroupDetails2(GrpID, groupData, W, SearchContainer, Window)

		})

		l, _ := fyne.LoadResourceFromURLString(imageURL)
		img := canvas.NewImageFromResource(l)
		img.FillMode = canvas.ImageFillContain // Gestion du fill image
		img.SetMinSize(fyne.NewSize(200, 200)) //Définir la taille minimum de l'image
		img.Resize(fyne.NewSize(200, 200))

		r, g, b, a := colorAnalysis.CalculateAverageColor(img)

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

		Infoback = container.New(layout.NewBorderLayout(nil, nil, nil, nil), background2, iinfo)

		info := container.New(layout.NewVBoxLayout(),

			img,
			container.NewCenter(Infoback),
		)

		Card = container.New(layout.NewBorderLayout(nil, nil, nil, nil), background, info)

		Card.Resize(fyne.NewSize(300, 300)) //Définir la taille minimum de la Card

		border := canvas.NewRectangle(color.Transparent) // Définir une couleur transparente pour le remplissage
		border.SetMinSize(fyne.NewSize(300, 300))        //Définir la taille minimum de la bordure
		border.Resize(fyne.NewSize(296, 296))            // Redimensionner pour inclure les coin
		border.StrokeColor = color.Black                 // Définir la couleur de la bordure
		border.StrokeWidth = 3                           // Définir l'épaisseur de la bordure
		border.CornerRadius = 20                         // Définir les coins

		Card.Add(border)

		Card.Resize(fyne.NewSize(100, 300))

		Listcard.Add(Card)
		Def.Add(Card)

	}

}

func MakeStringList(stringname []string, groupData []structdata.GroupData) *widget.List {

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

func MakeUpperUI(stringList *widget.List, stringname []string, groupData []structdata.GroupData, valueLabel *widget.Label, slider *widget.Slider, groupDataDates structdata.DatesData, stringdate [][]string) *fyne.Container {
	search := widget.NewEntry()
	searchButton := widget.NewButton("Rechercher", func() {
		// Vérifier si stringList est nul
		if W.Content == nil {
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
		stringList.Refresh()
	})

	stringList.OnSelected = func(id widget.ListItemID) {
		groupID := groupData[id].ID
		ShowGroupDetails(groupID, groupData, W, SearchContainer) // Passer la liste de recherche et la barre de recherche à la fonction
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
							ShowGroupDetails2(groupID, groupData, W, SearchContainer, Window) // Passer SearchContainer à la fonction
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
								ShowGroupDetails2(groupID, groupData, W, SearchContainer, Window) // Passer SearchContainer à la fonction
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
		for _, favorisGroup := range ListFavorit {

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

					r, g, b, a := colorAnalysis.CalculateAverageColor(img)

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

					Infoback = container.New(layout.NewBorderLayout(nil, nil, nil, nil), background2, iinfo)

					info := container.New(layout.NewVBoxLayout(),

						img,
						container.NewCenter(Infoback),
					)

					favoritCard = container.New(layout.NewBorderLayout(nil, nil, nil, nil), background, info)

					favoritCard.Resize(fyne.NewSize(300, 300)) //Définir la taille minimum de la Card

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

			W.SetContent(Window)
		})

		favoritCardScroll := container.NewScroll(favoritMenu)
		favoritCardScroll.SetMinSize(fyne.NewSize(675, 675))

		favoritPage := container.NewVBox(
			backButton,
			favoritCardScroll)

		W.SetContent(favoritPage)

	})
	var clearButton *widget.Button

	clearButton = widget.NewButtonWithIcon("", theme.DeleteIcon(), func() {
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

		slider.SetValue(0)

		cardscroll := container.NewScroll(Listcard)

		cardscroll.SetMinSize(fyne.NewSize(675, 675))

		sugg2 = container.NewVScroll(sugg)

		sugg2.SetMinSize(fyne.NewSize(100, 100))
		sugg2.Hide()
		spacer := layout.NewSpacer()
		sugg3 := container.NewHBox(sugg2, spacer)
		spacer.Resize(fyne.NewSize(100, 200))

		researchbar := container.NewVBox(
			favoris,
			search,
			sugg3,
			searchButton,
			clearButton,
			valueLabel,
			slider,
		)

		researchbar.Add(cardscroll)
		W.Resize(fyne.NewSize(800, 600))

		Window = container.NewVBox(researchbar)

		W.SetContent(Window)

		//W.SetContent(container.NewVBox(SearchContainer))

	})

	searchButton.OnTapped = func() {
		// Désactiver barre de recherche
		search.Disable()

		searchText := strings.ToLower(search.Text)
		suggestions := make([]fyne.CanvasObject, 0)

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
					ShowGroupDetails(groupID, groupData, W, SearchContainer) // Passer SearchContainer à la fonction
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
			W.SetContent(r)
			search.Enable()

			return
		}

		if len(suggestions) > 0 {
			spacer := layout.NewSpacer()
			sugg3 := container.NewHBox(sugg2, spacer)
			spacer.Resize(fyne.NewSize(100, 200))
			rsrch := container.NewVBox(search, sugg3, searchButton, clearButton)
			suggestionsContainer := container.NewVBox(suggestions...)
			Wind222 = container.NewVBox(rsrch, suggestionsContainer)
			W.SetContent(container.NewVScroll(Wind222))
		} else {
			// Afficher un message si aucune suggestion n'est trouvée
			r := container.NewVBox(widget.NewLabel("Aucun groupe trouvé avec ce nom."))
			r.Add(clearButton)

			W.SetContent(r)
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
		Listcard.RemoveAll()
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

				r, g, b, a := colorAnalysis.CalculateAverageColor(img)

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

				Infoback = container.New(layout.NewBorderLayout(nil, nil, nil, nil), background2, iinfo)

				info := container.New(layout.NewVBoxLayout(),

					img,
					container.NewCenter(Infoback),
				)

				Card = container.New(layout.NewBorderLayout(nil, nil, nil, nil), background, info)

				Card.Resize(fyne.NewSize(300, 300)) //Définir la taille minimum de la Card

				border := canvas.NewRectangle(color.Transparent) // Définir une couleur transparente pour le remplissage
				border.SetMinSize(fyne.NewSize(300, 300))        //Définir la taille minimum de la bordure
				border.Resize(fyne.NewSize(296, 296))            // Redimensionner pour inclure les coin
				border.StrokeColor = color.Black                 // Définir la couleur de la bordure
				border.StrokeWidth = 3                           // Définir l'épaisseur de la bordure
				border.CornerRadius = 20                         // Définir les coins

				Card.Add(border)

				Card.Resize(fyne.NewSize(100, 300))

				Listcard.Add(Card)
				deft = false

			}
		}
		if deft {
			Listcard.RemoveAll()
			for _, o := range Def.Objects {
				Listcard.Add(o)

			}
		}

	}

	sugg2 = container.NewVScroll(sugg)

	sugg2.SetMinSize(fyne.NewSize(100, 100))
	sugg2.Hide()
	spacer := layout.NewSpacer()
	sugg3 := container.NewHBox(sugg2, spacer)
	spacer.Resize(fyne.NewSize(100, 200))

	UpperUI := container.NewVBox(
		favoris,
		search,
		sugg3,
		searchButton,
		clearButton,
		valueLabel,
		slider,
	)
	return UpperUI

}
