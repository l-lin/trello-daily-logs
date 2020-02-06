package trello

// Card in trello
type Card struct {
	Name   string `json:"name"`
	Labels []Label
}

// Label in trello
type Label struct {
	Name string `json:"name"`
}

// ContainLabel checks if the card contains the given label
func (c Card) ContainLabel(label string) bool {
	for _, l := range c.Labels {
		if l.Name == label {
			return true
		}
	}
	return false
}
