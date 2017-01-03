package view

import "html/template"

// This class will deal with creating our HTML pages by adding the necessary
// assets (like CSS and Javascript) and facilitating the body customization.
type ViewModel struct {
    // This will describe the CSS style of the page
    Style template.CSS

    // This mapping will relate the data produced by the model to the view.
    Body map[string]string

    // TODO Add Javacript assets
}

// Creates a new view model.
func NewViewModel() *ViewModel {
    vm := ViewModel {
        Style: template.CSS(loadCss()),
        Body: make(map[string]string),
    }
    return &vm
}

// TODO Create add body function
