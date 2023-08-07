package schema

type slideLink struct {
	Type        string `json:"type"`
	Description string `json:"description"`
	Link        string `json:"link"`
}
type Publication struct {
	ID       int         `json:"id"`
	Title    string      `json:"title"`
	Cite     string      `json:"cite"`
	Link     string      `json:"link,omitempty"`
	Slides   []slideLink `json:"slides,omitempty"`
	Abstract string      `json:"abstract"`
}

type ReadingList struct {
	ID          int               `json:"id"`
	Description string            `json:"description"`
	Items       map[string]string `json:"items,omitempty"`
}
