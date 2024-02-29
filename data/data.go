package data

type UserData struct {
	Artists   int    `json:"artists"`
	Locations string `json:"locations"`
	Dates     string `json:"dates"`
	Relations string `json:"relations"`
}

/* stringList := widget.NewList(
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
	}) */
