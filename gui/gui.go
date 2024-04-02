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

// Declaration of global variables
var ListFavorit []string
var ResultWindow *fyne.Container
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
			// Create a text widget to display group details in the window
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

			// Create a container to display group details
			groupDetails := container.NewVBox(
				img,
				artist,
				members,
				album,
				creationDate,
				backButton,
			)

			// Place group details in the center of the window
			content := container.NewBorder(nil, nil, nil, nil, groupDetails)
			W.SetContent(content)
			return
		}
	}
}

func ShowGroupDetails(groupID int, groupData []structdata.GroupData, W fyne.Window, SearchContainer *fyne.Container) {
	backButton := widget.NewButton("Retour", func() {
		// Return to search list
		W.SetContent(container.NewVScroll(ResultWindow))
	})

	for _, group := range groupData {
		if group.ID == groupID {
			// Create a text widget to display group details in the window
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

			// Create a container to display group details
			groupDetails := container.NewVBox(
				img,
				artist,
				members,
				album,
				creationDate,
				backButton,
			)

			// Place group details in the center of the window
			content := container.NewBorder(nil, nil, nil, nil, groupDetails)
			W.SetContent(content)
			return
		}
	}
}

func MakeMenu(a fyne.App) *fyne.MainMenu {
	//Create menu and add to it items
	menu := fyne.NewMainMenu(

		// Theme of the page
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
	//Range artists API to collect info about them
	for _, group := range groupData {
		var listmember string
		for _, memb := range group.Members {
			listmember += memb
			listmember += "  "

		}
		// Create the variable which are going to be display on each card
		name := canvas.NewText(group.Name, color.Black)
		members := canvas.NewText(listmember, color.Black)
		imageURL := group.Image

		heartOnImage, _ := fyne.LoadResourceFromPath("./heartOn.png")
		heartOffImage, _ := fyne.LoadResourceFromPath("./heartOff.png")

		// Create a boolean to trace the state of the favourit button
		var isPressed bool
		var heartButton *widget.Button
		nameGroup := group.Name
		// Create favourit button

		heartButton = widget.NewButton("", func() {
			nameG := nameGroup

			// Change state when the button is pressed
			isPressed = !isPressed
			// Update the icon of the button
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

		//Create a button view detail to see more about a group

		viewDetail := widget.NewButton("View Detail", func() {
			GrpID := groupID

			ShowGroupDetails2(GrpID, groupData, W, SearchContainer, Window)

		})
		//Get and resize the group image
		l, _ := fyne.LoadResourceFromURLString(imageURL)
		img := canvas.NewImageFromResource(l)
		img.FillMode = canvas.ImageFillContain
		img.SetMinSize(fyne.NewSize(200, 200))
		img.Resize(fyne.NewSize(200, 200))

		//Get background color
		r, g, b, a := colorAnalysis.CalculateAverageColor(img)

		//Create the card background
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

		//Gathered the informations of the group's card
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
		//Set up the card

		Card = container.New(layout.NewBorderLayout(nil, nil, nil, nil), background, info)

		Card.Resize(fyne.NewSize(300, 300))

		//Create border to the card
		border := canvas.NewRectangle(color.Transparent)
		border.SetMinSize(fyne.NewSize(300, 300))
		border.Resize(fyne.NewSize(296, 296))
		border.StrokeColor = color.Black
		border.StrokeWidth = 3
		border.CornerRadius = 20

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
		// Check if stringList is null
		if W.Content == nil {
			return
		}
		// New list for search results
		filteredList := []string{}

		// Browse the original list and add the corresponding items to the new list
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

	//Create suggestion container
	suggestion := container.NewVBox()
	suggestionScroll := container.NewVScroll(suggestion)
	//Suggestions are updtate every time a user writte somthing in the search bar
	search.OnChanged = func(query string) {
		searchText := strings.ToLower(query)
		//Analyze the query and create the butons of each suggestion
		if len(query) > 0 {
			//Refresh the suggestion every input
			suggestion.Objects = make([]fyne.CanvasObject, 0)

			for _, group := range groupData {
				//Suggestion of a group
				if strings.Contains(strings.ToLower(group.Name), searchText) {
					label := group.Name + "        - Groupe"
					h := widget.NewButton(label, func(groupID int) func() {
						return func() {
							ShowGroupDetails2(groupID, groupData, W, SearchContainer, Window) // Passer SearchContainer à la fonction
						}
					}(group.ID))
					h.Importance = widget.LowImportance
					suggestion.Add(h)

				}
				for _, groupMember := range group.Members {
					//Suggestion of a member
					if strings.Contains(strings.ToLower(groupMember), searchText) {
						label2 := groupMember + "         - Member"
						h := widget.NewButton(label2, func(groupID int) func() {
							return func() {
								ShowGroupDetails2(groupID, groupData, W, SearchContainer, Window) // Passer SearchContainer à la fonction
							}

						}(group.ID))
						h.Importance = widget.LowImportance

						suggestion.Add(h)

					}
				}

			}
			suggestion.Show()
			suggestionScroll.Show()
		} else {
			//Hide suggestion container if it is empty
			suggestion.Hide()
			suggestionScroll.Hide()
		}

	}
	//Create View Favourite button
	favourite := widget.NewButton("View favourites", func() {
		var favoriteCard *fyne.Container
		favoriteMenu := container.NewVBox()
		// Range the String list of the name's group wich are in favourite
		for _, favorisGroup := range ListFavorit {

			for _, group := range groupData {
				//Create card of favourite's group
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
					img.FillMode = canvas.ImageFillContain
					img.SetMinSize(fyne.NewSize(200, 200))
					img.Resize(fyne.NewSize(200, 200))

					r, g, b, a := colorAnalysis.CalculateAverageColor(img)

					background := canvas.NewRectangle(color.RGBA{uint8(r), uint8(g), uint8(b), uint8(a)})
					background.FillColor = color.RGBA{uint8(r), uint8(g), uint8(b), uint8(a)}

					background.SetMinSize(fyne.NewSize(300, 300))
					background.Resize(fyne.NewSize(296, 296))
					background.CornerRadius = 20

					background2 := canvas.NewRectangle(color.RGBA{255, 255, 255, 255})
					background2.FillColor = color.RGBA{255, 255, 255, 255}

					background2.SetMinSize(fyne.NewSize(100, 100))
					background2.Resize(fyne.NewSize(100, 100))
					background2.CornerRadius = 20

					iinfo := container.New(layout.NewVBoxLayout(),
						container.NewCenter(name),
						container.NewCenter(members))

					Infoback = container.New(layout.NewBorderLayout(nil, nil, nil, nil), background2, iinfo)

					info := container.New(layout.NewVBoxLayout(),

						img,
						container.NewCenter(Infoback),
					)

					favoriteCard = container.New(layout.NewBorderLayout(nil, nil, nil, nil), background, info)

					favoriteCard.Resize(fyne.NewSize(300, 300))

					border := canvas.NewRectangle(color.Transparent)
					border.SetMinSize(fyne.NewSize(300, 300))
					border.Resize(fyne.NewSize(296, 296))
					border.StrokeColor = color.Black
					border.StrokeWidth = 3
					border.CornerRadius = 20

					favoriteCard.Add(border)

					favoriteCard.Resize(fyne.NewSize(100, 300))

					favoriteMenu.Add(favoriteCard)

				}
			}

		}
		//Create button to go back to the homepage

		backButton := widget.NewButton("Retour", func() {

			W.SetContent(Window)
		})

		favoritCardScroll := container.NewScroll(favoriteMenu)
		favoritCardScroll.SetMinSize(fyne.NewSize(675, 675))

		favoritPage := container.NewVBox(
			backButton,
			favoritCardScroll)

		//Display the favorite page
		W.SetContent(favoritPage)

	})
	var clearButton *widget.Button
	// Clear button let the user go back to the homepage and reset the filter
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

		suggestionScroll = container.NewVScroll(suggestion)

		suggestionScroll.SetMinSize(fyne.NewSize(100, 100))
		suggestionScroll.Hide()
		spacer := layout.NewSpacer()
		sugg3 := container.NewHBox(suggestionScroll, spacer)
		spacer.Resize(fyne.NewSize(100, 200))

		researchbar := container.NewVBox(
			favourite,
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

	})

	searchButton.OnTapped = func() {
		// Disable search bar
		search.Disable()

		searchText := strings.ToLower(search.Text)
		suggestions := make([]fyne.CanvasObject, 0)

		verif := false
		for _, group := range groupData {
			imageURL := group.Image

			l, _ := fyne.LoadResourceFromURLString(imageURL)
			img := canvas.NewImageFromResource(l)
			img.FillMode = canvas.ImageFillContain
			img.SetMinSize(fyne.NewSize(120, 120))
			img.Resize(fyne.NewSize(120, 120))

			// Create a custom button with the group's image and name
			resultat := widget.NewButton("", func(groupID int) func() {
				return func() {
					// Pass SearchContainer to function
					ShowGroupDetails(groupID, groupData, W, SearchContainer)
				}
			}(group.ID))
			// Reduce importance so that it doesn't look like a standard button
			resultat.Importance = widget.LowImportance
			resultat.SetIcon(l)
			// Set image as button icon
			resultat.Resize(fyne.NewSize(200, 200))
			 // Define group name as button text
			resultat.SetText(group.Name)

			// Add button to suggestion list
			if strings.Contains(strings.ToLower(group.Name), searchText) {
				suggestions = append(suggestions, resultat)
				verif = true
			} else {
				for _, member := range group.Members {
					if strings.Contains(strings.ToLower(member), searchText) {
						suggestions = append(suggestions, resultat)
						verif = true
						break
					}
				}
			}

			if strings.Contains(fmt.Sprintf("%d", group.CreationDate), searchText) {
				suggestions = append(suggestions, resultat)
				verif = true
			}

			for _, date := range groupDataDates.Index {
				if strings.Contains(strings.ToLower(date.Dates[0]), searchText) {
					suggestions = append(suggestions, resultat)
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
			sugg3 := container.NewHBox(suggestionScroll, spacer)
			spacer.Resize(fyne.NewSize(100, 200))
			rsrch := container.NewVBox(search, sugg3, searchButton, clearButton)
			suggestionsContainer := container.NewVBox(suggestions...)
			ResultWindow = container.NewVBox(rsrch, suggestionsContainer)
			W.SetContent(container.NewVScroll(ResultWindow))
		} else {
			// Afficher un message si aucune resultat n'est trouvée
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
	// Slider effect when changes
	slider.OnChanged = func(value float64) {
		deft = true
		Listcard.RemoveAll()
		valueLabel.SetText(fmt.Sprintf("Creation Year : %d", int(slider.Value)))
		for _, group := range groupData {
			//Add to the list of card only the groups where their creation date is the same as the slider value
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
				img.FillMode = canvas.ImageFillContain
				img.SetMinSize(fyne.NewSize(200, 200))
				img.Resize(fyne.NewSize(200, 200))

				r, g, b, a := colorAnalysis.CalculateAverageColor(img)

				background := canvas.NewRectangle(color.RGBA{uint8(r), uint8(g), uint8(b), uint8(a)})
				background.FillColor = color.RGBA{uint8(r), uint8(g), uint8(b), uint8(a)}

				background.SetMinSize(fyne.NewSize(300, 300))
				background.Resize(fyne.NewSize(296, 296))
				background.CornerRadius = 20

				background2 := canvas.NewRectangle(color.RGBA{255, 255, 255, 255})
				background2.FillColor = color.RGBA{255, 255, 255, 255}

				background2.SetMinSize(fyne.NewSize(100, 100))
				background2.Resize(fyne.NewSize(100, 100))
				background2.CornerRadius = 20

				iinfo := container.New(layout.NewVBoxLayout(),
					container.NewCenter(name),
					container.NewCenter(members))

				Infoback = container.New(layout.NewBorderLayout(nil, nil, nil, nil), background2, iinfo)

				info := container.New(layout.NewVBoxLayout(),

					img,
					container.NewCenter(Infoback),
				)

				Card = container.New(layout.NewBorderLayout(nil, nil, nil, nil), background, info)

				Card.Resize(fyne.NewSize(300, 300))

				border := canvas.NewRectangle(color.Transparent)
				border.SetMinSize(fyne.NewSize(300, 300))
				border.Resize(fyne.NewSize(296, 296))
				border.StrokeColor = color.Black
				border.StrokeWidth = 3
				border.CornerRadius = 20

				Card.Add(border)

				Card.Resize(fyne.NewSize(100, 300))

				Listcard.Add(Card)
				deft = false

			}
		}
		// If their is no result matching the slider value, display all the group
		if deft {
			Listcard.RemoveAll()
			for _, o := range Def.Objects {
				Listcard.Add(o)

			}
		}

	}
	//Refine suggestion container

	suggestionScroll = container.NewVScroll(suggestion)

	suggestionScroll.SetMinSize(fyne.NewSize(100, 100))
	suggestionScroll.Hide()
	spacer := layout.NewSpacer()
	sugg3 := container.NewHBox(suggestionScroll, spacer)
	spacer.Resize(fyne.NewSize(100, 200))
	//Create a container whi contains all the features
	UpperUI := container.NewVBox(
		favourite,
		search,
		sugg3,
		searchButton,
		clearButton,
		valueLabel,
		slider,
	)

	return UpperUI

}
