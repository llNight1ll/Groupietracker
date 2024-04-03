package gui

import (
	"context"
	"fmt"
	"groupie/colorAnalysis"
	"groupie/structdata"
	"image/color"
	"net/url"
	"os"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/zmb3/spotify"
	"golang.org/x/oauth2/clientcredentials"
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

func ShowGroupDetails2(groupID int, groupData []structdata.GroupData, W fyne.Window, SearchContainer *fyne.Container, Window *fyne.Container, a fyne.App) {
	backButton := widget.NewButton("Retour", func() {
		W.SetContent(Window)
	})

	for _, group := range groupData {
		if group.ID == groupID {

			// Create a page title
			title := widget.NewLabelWithStyle("Info Artist :", fyne.TextAlignLeading, fyne.TextStyle{Bold: true})
			title.Move(fyne.NewPos(520, 0))

			// Create a text widget to display group details in the window
			artist := widget.NewLabelWithStyle(group.Name, fyne.TextAlignLeading, fyne.TextStyle{Bold: true})
			artistCenter := fyne.NewContainerWithLayout(layout.NewCenterLayout(), artist)

			// Convert slice to string
			membersLabel := "- "
			for i, member := range group.Members {
				if i > 0 {
					membersLabel += ", \n -"
				}
				membersLabel += member
			}
			members := widget.NewLabelWithStyle("Membres: \n"+membersLabel, fyne.TextAlignLeading, fyne.TextStyle{Bold: true, Italic: true})

			album := widget.NewLabel(group.FirstAlbum)
			albumText := album.Text
			// Create a label widget to display ""First album released" and the band name with the date of the first album released
			firstAlbum := widget.NewLabelWithStyle("Premier album sorti de "+group.Name+" sorti le: "+albumText, fyne.TextAlignLeading, fyne.TextStyle{Bold: true})

			creationDate := widget.NewLabel(fmt.Sprintf("%d", group.CreationDate))
			// Get the text value of creationDate
			creationDateText := creationDate.Text
			// Create a label widget to display "Career start" and the group name with creation date
			carriere := widget.NewLabelWithStyle("Debut de carrière de "+group.Name+" en "+creationDateText, fyne.TextAlignLeading, fyne.TextStyle{Bold: true})
			imageURL := group.Image

			r, _ := fyne.LoadResourceFromURLString(imageURL)
			img := canvas.NewImageFromResource(r)
			img.FillMode = canvas.ImageFillContain // Image fill management
			img.SetMinSize(fyne.NewSize(350, 350)) // Set minimum image size
			img.Resize(fyne.NewSize(350, 350))
			trackinfo := GetSpotifyData(group.Name, a)
			info := container.NewHBox(container.NewVBox(
				artistCenter,
				members,
				firstAlbum,
				carriere,
			), container.NewVScroll(trackinfo))

			// Create a container to display group details
			groupDetails := container.NewVBox(
				title,
				img,
				info,
				backButton,
			)

			// Place group details in the center of the window
			content := container.NewBorder(nil, nil, nil, nil, groupDetails)
			W.SetContent(content)
			return
		}
	}

}

func ShowGroupDetails(groupID int, groupData []structdata.GroupData, W fyne.Window, SearchContainer *fyne.Container, a fyne.App) {
	backButton := widget.NewButton("Retour", func() {
		// Retour à la liste de recherche
		W.SetContent(container.NewVScroll(ResultWindow))
	})

	for _, group := range groupData {
		if group.ID == groupID {

			// Create a page title
			title := widget.NewLabelWithStyle("Info Artist :", fyne.TextAlignLeading, fyne.TextStyle{Bold: true})
			title.Move(fyne.NewPos(520, 0))

			// Create a text widget to display group details in the window
			artist := widget.NewLabelWithStyle(group.Name, fyne.TextAlignLeading, fyne.TextStyle{Bold: true})
			artistCenter := fyne.NewContainerWithLayout(layout.NewCenterLayout(), artist)

			// Convert slice to string
			membersLabel := "- "
			for i, member := range group.Members {
				if i > 0 {
					membersLabel += ", \n -"
				}
				membersLabel += member
			}
			members := widget.NewLabelWithStyle("Membres: \n"+membersLabel, fyne.TextAlignLeading, fyne.TextStyle{Bold: true, Italic: true})

			album := widget.NewLabel(group.FirstAlbum)
			albumText := album.Text
			// Create a label widget to display ""First album released" and the band name with the date of the first album released
			firstAlbum := widget.NewLabelWithStyle("Premier album sorti de "+group.Name+" sorti le: "+albumText, fyne.TextAlignLeading, fyne.TextStyle{Bold: true})

			creationDate := widget.NewLabel(fmt.Sprintf("%d", group.CreationDate))
			// Get the text value of creationDate
			creationDateText := creationDate.Text
			// Create a label widget to display "Career start" and the group name with creation date
			carriere := widget.NewLabelWithStyle("Debut de carrière de "+group.Name+" en "+creationDateText, fyne.TextAlignLeading, fyne.TextStyle{Bold: true})
			imageURL := group.Image

			r, _ := fyne.LoadResourceFromURLString(imageURL)
			img := canvas.NewImageFromResource(r)
			img.FillMode = canvas.ImageFillContain // Image fill management
			img.SetMinSize(fyne.NewSize(350, 350)) // Set minimum image size
			img.Resize(fyne.NewSize(350, 350))

			trackinfo := GetSpotifyData(group.Name, a)
			info := container.NewHBox(container.NewVBox(
				artistCenter,
				members,
				firstAlbum,
				carriere,
			), container.NewVScroll(trackinfo))

			// Create a container to display group details
			groupDetails := container.NewVBox(
				title,
				img,
				info,
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
				lien, _ := url.Parse("https://open.spotify.com")
				_ = a.OpenURL(lien)
			}),
		),
	)
	return menu
}

func MakeListCard(Card *fyne.Container, Infoback *fyne.Container, Listcard *fyne.Container, Def *fyne.Container, groupData []structdata.GroupData, a fyne.App) {
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

		})
		heartButton.SetIcon(heartOffImage)
		heartButton.Importance = widget.LowImportance

		groupID := group.ID

		//Create a button view detail to see more about a group

		viewDetail := widget.NewButton("View Detail", func() {
			GrpID := groupID

			ShowGroupDetails2(GrpID, groupData, W, SearchContainer, Window, a)

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

func MakeUpperUI(stringList *widget.List, stringname []string, groupData []structdata.GroupData, valueLabel *widget.Label, slider *widget.Slider, groupDataDates structdata.DatesData, stringdate [][]string, a fyne.App) *fyne.Container {

	var checkboxFilters []*widget.Check

	// Create checkboxes for filters
	for i := 2; i <= 7; i++ {
		checkbox := widget.NewCheck(fmt.Sprintf("%d", i), func(members int) func(bool) {
			return func(checked bool) {
			}
		}(i))
		checkboxFilters = append(checkboxFilters, checkbox)
	}

	// Create a container to organize checkboxes
	checkboxContainer := container.NewHBox()
	label := widget.NewLabel("Nombre de membres:")
	checkboxContainer.Add(label)

	// Ajoute toutes les cases à cocher au conteneur
	for _, checkbox := range checkboxFilters {
		checkboxContainer.Add(checkbox)
	}

	// Initialize a map to store hints for selected checkboxes
	selectedMembers := make(map[int]bool)
	// Scroll through each checkbox to determine whether it is checked
	for i, checkbox := range checkboxFilters {
		if checkbox.Checked {
			selectedMembers[i+2] = true
		}
	}

	// Filter results by selected members
	filteredResults := make([]structdata.GroupData, 0)
	for _, group := range groupData {
		if len(selectedMembers) == 0 {
			// No filter selected, simply add all results
			filteredResults = append(filteredResults, group)
		} else {
			// Check if the group has the selected number of members
			if selectedMembers[len(group.Members)] {
				filteredResults = append(filteredResults, group)
			}
		}
	}

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

		for _, annee := range stringdate {
			for _, item := range annee {
				if strings.Contains(strings.ToLower(item), strings.ToLower(search.Text)) {
					filteredList = append(filteredList, item)
				}
			}
		}

		// Update list with search results
		stringList.Length = func() int {
			return len(filteredList)
		}
		stringList.CreateItem = func() fyne.CanvasObject {
			return widget.NewLabel("")
		}
		stringList.UpdateItem = func(index int, item fyne.CanvasObject) {
			item.(*widget.Label).SetText(filteredList[index])
		}

		stringList.UpdateItem = func(index int, annee fyne.CanvasObject) {
			annee.(*widget.Label).SetText(filteredList[index])
		}
		// Initialize a map to store hints for selected checkboxes
		selectedMembers := make(map[int]bool)
		// Scroll through each checkbox to determine whether it is checked
		for i, checkbox := range checkboxFilters {
			if checkbox.Checked {
				selectedMembers[i+2] = true
			}
		}

		// Filter results by selected members
		filteredResults := make([]structdata.GroupData, 0)
		for _, group := range groupData {
			if len(selectedMembers) == 0 {
				// No filter selected, simply add all results
				filteredResults = append(filteredResults, group)
			} else {
				// Check if the group has the selected number of members
				if selectedMembers[len(group.Members)] {
					filteredResults = append(filteredResults, group)
				}
			}
		}
		stringList.Refresh()
	})

	stringList.OnSelected = func(id widget.ListItemID) {
		groupID := groupData[id].ID
		// Switch the search list and the search bar to the function
		ShowGroupDetails(groupID, groupData, W, SearchContainer, a)
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
							ShowGroupDetails2(groupID, groupData, W, SearchContainer, Window, a) // Passer SearchContainer à la fonction
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
								ShowGroupDetails2(groupID, groupData, W, SearchContainer, Window, a) // Passer SearchContainer à la fonction
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

		// Reset selected filters
		for _, checkbox := range checkboxFilters {
			checkbox.Checked = false
		}

		// Update list with search results
		stringList.Length = func() int {
			return len(stringname)
		}
		stringList.CreateItem = func() fyne.CanvasObject {
			return widget.NewLabel("")
		}
		stringList.UpdateItem = func(index int, item fyne.CanvasObject) {
			item.(*widget.Label).SetText(stringname[index])
		}

		// Initialize a map to store hints for selected checkboxes
		selectedMembers := make(map[int]bool)
		// Scroll through each checkbox to determine whether it is checked
		for i, checkbox := range checkboxFilters {
			if checkbox.Checked {
				selectedMembers[i+2] = true
			}
		}

		// Filter results by selected members
		filteredResults := make([]structdata.GroupData, 0)
		for _, group := range groupData {
			if len(selectedMembers) == 0 {
				// No filter selected, simply add all results
				filteredResults = append(filteredResults, group)
			} else {
				// Check if the group has the selected number of members
				if selectedMembers[len(group.Members)] {
					filteredResults = append(filteredResults, group)
				}
			}
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
			checkboxContainer,
			search,
			sugg3,
			searchButton,
			clearButton,
			valueLabel,
			slider,
		)

		researchbar.Add(cardscroll)
		W.Resize(fyne.NewSize(1250, 800))

		Window = container.NewVBox(researchbar)

		W.SetContent(Window)

	})

	searchButton.OnTapped = func() {
		// Disable search bar
		search.Disable()
	
		suggestions := make([]fyne.CanvasObject, 0)
		verif := false
	
		// Initialize a map to store hints for selected checkboxes
		selectedMembers := make(map[int]bool)
		// Scroll through each checkbox to determine whether it is checked
		for i, checkbox := range checkboxFilters {
			if checkbox.Checked {
				selectedMembers[i+2] = true
			}
		}

		// Filter results by selected members
		filteredResults := make([]structdata.GroupData, 0)
		for _, group := range groupData {
			if len(selectedMembers) == 0 {
				// No filter selected, simply add all results
				filteredResults = append(filteredResults, group)
			} else {
				// Check if the group has the selected number of members
				if selectedMembers[len(group.Members)] {
					filteredResults = append(filteredResults, group)
				}
			}
		}
	
		// Browse filtered search results
		for _, group := range filteredResults {
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
					ShowGroupDetails(groupID, groupData, W, SearchContainer, a)
				}
			}(group.ID))
			// Reduce its size so that it doesn't look like a standard button
			resultat.Importance = widget.LowImportance
			resultat.SetIcon(l)
			// Set image as button icon
			resultat.Resize(fyne.NewSize(200, 200))
			// Define group name as button text
			resultat.SetText(group.Name)
	
			// Add button to suggestion list
			suggestions = append(suggestions, resultat)
			verif = true
		}
	
		// Display suggestions or a message if no results are found
		if verif {
			spacer := layout.NewSpacer()
			sugg3 := container.NewHBox(suggestionScroll, spacer)
			spacer.Resize(fyne.NewSize(100, 200))
			rsrch := container.NewVBox(search, sugg3, searchButton, clearButton)
			suggestionsContainer := container.NewVBox(suggestions...)
			ResultWindow = container.NewVBox(rsrch, suggestionsContainer)
			W.SetContent(container.NewVScroll(ResultWindow))
		} else {
			// Display a message if no results are found
			r := container.NewVBox(widget.NewLabel("Aucun groupe trouvé avec ce nom."))
			r.Add(clearButton)
			W.SetContent(r)
		}
	
		// Reactivate search bar
		search.Enable()
	}
	

	search.OnSubmitted = func(text string) {
		// Start search when "Enter" key is pressed
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
		checkboxContainer,
		search,
		sugg3,
		searchButton,
		clearButton,
		valueLabel,
		slider,
	)

	return UpperUI
}

func GetSpotifyData(name string, a fyne.App) *fyne.Container {

	// Config a client
	config := &clientcredentials.Config{
		ClientID:     "670da40f1af74c558d0644f13b2b4898",
		ClientSecret: "2faaac80d70e48a7a9df5fbfee5f127d",
		TokenURL:     spotify.TokenURL,
	}
	client := config.Client(context.Background())

	// Use the client to use spotify
	spotifyClient := spotify.NewClient(client)

	artistName := name

	// Research the page of an artist
	result, err := spotifyClient.Search(artistName, spotify.SearchTypeArtist)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to search artist: %v\n", err)
	}

	// Get back the ID of the first artist found
	if len(result.Artists.Artists) > 0 {
		artistID := result.Artists.Artists[0].ID

		//Get the top 10 tracks of an artist

		tracks, err := spotifyClient.GetArtistsTopTracks(artistID, "US")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to retrieve artist's top tracks: %v\n", err)
		}

		//Create the container of the spotify informations
		imgtrck := container.NewVBox()
		topTracksContainer := container.NewVBox()
		topTracksContainer.Add(widget.NewLabel("Top tracks :"))

		for _, track := range tracks {
			//Contains the track's name
			trackName := widget.NewLabel(track.Name)

			//Get track informations
			infotrack, err := spotifyClient.GetTrack(track.ID)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Unable to retrieve artist's top tracks: %v\n", err)
			}
			//Contain the image of the cover of the track
			imglink := infotrack.Album.Images[0].URL
			r, _ := fyne.LoadResourceFromURLString(imglink)
			img2 := canvas.NewImageFromResource(r)
			img2.FillMode = canvas.ImageFillContain // Image fill management
			img2.SetMinSize(fyne.NewSize(50, 50))   // Set minimum image size
			img2.Resize(fyne.NewSize(50, 50))
			imgtrck.Add(img2)

			trackinfo := container.NewHBox(trackName, img2)
			spotifyLink := fmt.Sprintf("https://open.spotify.com/track/%s", track.ID)

			tracklink := widget.NewButton("Listen on Spotify", func() {
				lien, _ := url.Parse(spotifyLink)
				_ = a.OpenURL(lien)

			})
			tranckInfoandLink := container.NewVBox(trackinfo, tracklink)

			topTracksContainer.Add(tranckInfoandLink)

		}

		return topTracksContainer

	} else {
		return nil
	}
}
