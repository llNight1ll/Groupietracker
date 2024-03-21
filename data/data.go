package data

type UserData struct {
	Artists   int    `json:"artists"`
	Locations string `json:"locations"`
	Dates     string `json:"dates"`
	Relations string `json:"relations"`
}



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







	func createCardGeneralInfo(artist Artist) fyne.CanvasObject {
		image := canvas.NewImageFromFile(artist.Image) // Redimensionner l'image
		image.FillMode = canvas.ImageFillContain       // Gestion du fill image
		image.SetMinSize(fyne.NewSize(120, 120))       //Définir la taille minimum de l'image
		image.Resize(fyne.NewSize(120, 120))           //Définir la nouvelle image de l'image
	
		averageColor := getAverageColor(artist.Image) // Obtenir la couleur moyenne de l'image
	
		background := canvas.NewRectangle(averageColor) // Création d'un rectangle coloré pour l'arrière-plan de la card
		background.SetMinSize(fyne.NewSize(300, 300))   // Définir la taille minimum du bakcground
		background.Resize(fyne.NewSize(296, 296))       // Redimensionner pour inclure les coin
		background.CornerRadius = 20                    // Définir les coins arrondis
	
		button := widget.NewButton("Search", func() {
			fmt.Println(artist.Name)
		})
	nameLabel := widget.NewLabelWithStyle(artist.Name, fyne.TextAlignLeading, fyne.TextStyle{Bold: true}) // Nom de l'artiste en gras et plus gros
	
		yearLabel := widget.NewLabel(fmt.Sprintf("%d", artist.YearStarted)) // Date de début de l'artiste en plus petit
	
		labelsContainer := container.NewHBox( // Créer un conteneur HBox pour afficher les labels avec un espace entre eux
			nameLabel,
			yearLabel,
		)
	
		var membersText string        // Gestion des membres du groupe
		if len(artist.Members) == 1 { // Contion du nombre de membres
			membersText = "Solo Artist"
		} else if len(artist.Members) > 0 { // Contion du nombre de membres
			membersText = "Members:\n " + strings.Join(artist.Members, ", ")
		}
		membersLabel := widget.NewLabel(membersText) // EXPLIQUE ICI
		membersLabel.Wrapping = fyne.TextWrapWord    // Activer le wrapping du texte
	
		infoContainer := container.New(layout.NewVBoxLayout(), // Créer le conteneur pour les informations sur l'artiste
			layout.NewSpacer(), // Ajout d'un espace vertical
			image,              // Ajout de l'image
			labelsContainer,    // Ajout titre et date
			membersLabel,       // Afficher les membres du groupe
			layout.NewSpacer(), // Ajout d'un petit espace vertical
			button,
		)
	
		infoContainer.Resize(fyne.NewSize(300, 180)) // Définir la taille fixe pour le conteneur d'informations
	
		cardContent := container.New(layout.NewBorderLayout(nil, nil, nil, nil), background, infoContainer) //
	Créer le conteneur pour la card de l'artiste
		cardContent.Resize(fyne.NewSize(300, 300))                                                          //Définir la taille minimum de la card                                               //
	
		border := canvas.NewRectangle(color.Transparent) // Définir une couleur transparente pour le remplissage
		border.SetMinSize(fyne.NewSize(300, 300))        //Définir la taille minimum de la bordure
		border.Resize(fyne.NewSize(296, 296))            // Redimensionner pour inclure les coin
		border.StrokeColor = color.Black                 // Définir la couleur de la bordure
		border.StrokeWidth = 3                           // Définir l'épaisseur de la bordure
		border.CornerRadius = 20                         // Définir les coins
	
		cardContent.Add(border) // Ajouter le rectangle de contour à la carte
	
		return cardContent
	}
	